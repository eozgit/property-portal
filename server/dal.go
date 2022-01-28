package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	pb "github.com/eozgit/property-portal/propertyportal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataAccessLayer struct {
	db *gorm.DB
}

func (dal *DataAccessLayer) initDb() {
	db_file := "/tmp/property.db"
	os.Remove(db_file)
	db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	if err != nil {
		panic("Database connection failure")
	}

	dal.db = db

	dal.db.AutoMigrate(&Property{}, &PropertyImage{})

	dal.seedProperties()
	dal.seedPropertyImages()
}

func (dal *DataAccessLayer) seedProperties() {
	seedProperties(dal)
}

func (dal *DataAccessLayer) seedPropertyImages() {
	seedPropertyImages(dal)
}

var propertyTypeMap = map[uint32]string{
	0: "dwelling",
	1: "detached house",
	2: "semi-detached house",
	3: "terraced house",
	4: "flat",
}

var studioPropertyTypeMap = map[uint32]string{
	0: "dwelling",
	1: "house",
	2: "house",
	3: "house",
	4: "flat",
}

func getTitle(beds int, propertyType int) string {
	if beds > 0 {
		return fmt.Sprintf("%d bedroom %s", beds, propertyTypeMap[uint32(propertyType)])
	} else {
		return fmt.Sprintf("Studio %s", studioPropertyTypeMap[uint32(propertyType)])
	}
}

func (dal *DataAccessLayer) findProperties(filters *pb.Filters) *pb.Properties {
	var results []Property
	query := dal.db
	if len(filters.Location) > 3 {
		query = query.Where("lower(location) like ?", fmt.Sprintf("%%%s%%", strings.ToLower(filters.Location)))
	}
	if filters.MinPrice > 0 {
		query = query.Where("price >= ?", filters.MinPrice)
	}
	if filters.MaxPrice > 0 {
		query = query.Where("price <= ?", filters.MinPrice)
	}
	if filters.MinBeds > 0 {
		query = query.Where("beds >= ?", filters.MinBeds)
	}
	if filters.MaxBeds > 0 {
		query = query.Where("beds <= ?", filters.MaxBeds)
	}
	if filters.PropertyType > 0 {
		query = query.Where("propertyType = ?", filters.PropertyType)
	}
	if filters.MustHaves != nil {
		if filters.MustHaves.Garden != 0 {
			query = query.Where("garden = ?", filters.MustHaves.Garden == 1)
		}
		if filters.MustHaves.Parking != 0 {
			query = query.Where("parking = ?", filters.MustHaves.Parking == 1)
		}
		if filters.MustHaves.NewHome != 0 {
			query = query.Where("newHome = ?", filters.MustHaves.NewHome == 1)
		}
	}
	query.Find(&results)

	properties := []*pb.Property{}
	for _, record := range results {
		property := pb.Property{
			Id:    uint64(record.ID),
			Title: record.Title,
		}
		properties = append(properties, &property)
	}
	response := pb.Properties{Properties: properties}
	return &response
}

var features = []string{
	"Very well presented throughout",
	"Modern kitchen",
	"Sought after location",
	"Rarely Available Cul-De-Sac Location",
	"Utility Room",
	"Driveway Parking For Multiple Vehicles",
	"Short Distance to Bus Stops",
	"Close To Local Amenities",
	"Conservatory",
	"No upward chain",
	"Gas central heating",
}

func getSample(population *[]string, min int, max int) []string {
	sample := append([]string(nil), *population...)
	rand.Shuffle(len(*population), func(i, j int) { sample[i], sample[j] = sample[j], sample[i] })
	k := rand.Intn(max-min) + min
	return sample[:k]
}

func getDescription() string {
	sample := getSample(&features, 1, 3)
	return strings.Join(sample[:], ", ")
}

func (dal *DataAccessLayer) getPropertyDetails(property *pb.Property) *pb.PropertyDetails {
	var dto *Property
	dal.db.First(&dto, property.Id)
	details := pb.PropertyDetails{
		Property: &pb.Property{
			Id:    uint64(dto.ID),
			Title: dto.Title,
		},
		Description:  dto.Description,
		Location:     dto.Location,
		Price:        dto.Price,
		Beds:         dto.Beds,
		Bathrooms:    dto.Bathrooms,
		PropertyType: pb.PropertyType(dto.PropertyType),
		Features: &pb.Features{
			Garden:  -1,
			Parking: -1,
			NewHome: -1,
		},
		EnergyEfficiencyRating: &pb.EnergyEfficiencyRating{
			Current:   dto.CurrentRating,
			Potential: dto.PotentialRating,
		},
	}

	features := details.Features
	if dto.Garden {
		features.Garden = 1
	}
	if dto.Parking {
		features.Parking = 1
	}
	if dto.NewHome {
		features.NewHome = 1
	}

	return &details
}

func (dal *DataAccessLayer) getPropertyImages(propertyId uint) []pb.Image {
	var dto []PropertyImage
	dal.db.Where("property_id = ?", propertyId).Find(&dto)

	images := []pb.Image{}
	for _, record := range dto {
		image := pb.Image{
			Image: &pb.Image_Url{
				Url: record.Url,
			},
		}
		images = append(images, image)
	}

	return images
}
