package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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

	dal.db.AutoMigrate(&Property{})

	dal.seedProperties()
}

func (dal *DataAccessLayer) seedProperties() {
	r := csv.NewReader(strings.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var properties = []Property{}

	for _, record := range records {
		population, errPopulation := strconv.ParseUint(strings.ReplaceAll(record[2], ",", ""), 10, 64)
		if errPopulation != nil {
			log.Fatalf("failed to convert population '%s'", record[2])
		}
		district := record[1]
		county := record[4]
		region := record[5]

		propertyCount := population/100000 + 1

		for i := 0; i < int(propertyCount); i++ {
			beds := rand.Intn(6)
			var bathrooms uint32
			if beds == 0 {
				bathrooms = 1
			} else {
				bathrooms = uint32(rand.Intn(beds))
				if bathrooms == 0 {
					bathrooms = 1
				}
			}
			price := (rand.Intn(20)*5 + (beds+1)*100) * 1000
			propertyType := rand.Intn(4) + 1
			currentRating := rand.Intn(100)
			property := Property{
				Title:           getTitle(beds, propertyType),
				Description:     getTitle(beds, propertyType),
				Location:        fmt.Sprintf("%s, %s, %s", district, county, region),
				Price:           float64(price),
				Beds:            uint32(beds),
				Bathrooms:       uint32(bathrooms),
				PropertyType:    uint32(propertyType),
				Garden:          rand.Intn(2) == 0,
				Parking:         rand.Intn(2) == 0,
				NewHome:         rand.Intn(2) == 0,
				CurrentRating:   uint32(currentRating),
				PotentialRating: uint32(rand.Intn(100-currentRating) + currentRating),
			}
			properties = append(properties, property)
		}
	}

	dal.db.CreateInBatches(&properties, 100)
}

var propertyTypeMap = map[uint32]string{
	0: "detached house",
	1: "semi-detached house",
	2: "terraced house",
	3: "flat",
}

func getTitle(beds int, propertyType int) string {
	if beds > 0 {
		return fmt.Sprintf("%d bedroom %s", beds, propertyTypeMap[uint32(propertyType)])
	} else {
		var titleType string
		if propertyType == 3 {
			titleType = "flat"
		} else {
			titleType = "house"
		}
		return fmt.Sprintf("Studio %s", titleType)
	}
}

func (dal *DataAccessLayer) findProperties(filters *pb.Filters) *pb.Properties {
	var results []Property
	query := dal.db
	if len(filters.Location) > 3 {
		query = query.Where("location like ?", fmt.Sprintf("%%%s%%", filters.Location))
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
		if filters.MustHaves.Garden > -1 {
			query = query.Where("garden = ?", filters.MustHaves.Garden > 0)
		}
		if filters.MustHaves.Parking > -1 {
			query = query.Where("parking = ?", filters.MustHaves.Parking > 0)
		}
		if filters.MustHaves.NewHome > -1 {
			query = query.Where("newHome = ?", filters.MustHaves.NewHome > 0)
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
