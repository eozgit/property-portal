package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func seedProperties(dal *DataAccessLayer) {
	absPath, _ := filepath.Abs("./server/locations.csv")
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Unable to read locations file", err)
	}
	defer f.Close()

	r := csv.NewReader(f)

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
		if county == "Cambridgeshire" { // Add extra data for demo
			propertyCount *= 100
		}

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
				Description:     getDescription(),
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

var rooms = []string{
	"living",
	"dining",
	"kitchen",
	"bedroom",
	"bathroom",
	"hall",
	"conservatory",
	"garage",
	"balcony",
}

func seedPropertyImages(dal *DataAccessLayer) {
	var properties []Property
	dal.db.Find(&properties)

	var images = []PropertyImage{}
	for _, property := range properties {
		sample := getSample(&rooms, 3, 5)
		url := gofakeit.URL()
		for _, room := range sample {
			image := PropertyImage{PropertyId: property.ID, Url: fmt.Sprintf("%s/%s.png", url, room)}
			images = append(images, image)
		}
	}

	dal.db.CreateInBatches(&images, 100)
}
