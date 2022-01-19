package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/eozgit/property-portal/propertyportal"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type propertyPortalServer struct {
	pb.UnimplementedPropertyPortalServer
}

func (s *propertyPortalServer) FindProperties(ctx context.Context, filters *pb.Filters) (*pb.Properties, error) {
	log.Printf("FindProperties")
	props := pb.Properties{}
	props.Properties = append(props.Properties, &pb.Property{Id: 1, Title: "3 bedroom house"})
	return &props, nil
}

func (s *propertyPortalServer) GetPropertyDetails(ctx context.Context, filters *pb.Property) (*pb.PropertyDetails, error) {
	return &pb.PropertyDetails{}, nil
}

func (s *propertyPortalServer) ListFeatures(rect *pb.Property, stream pb.PropertyPortal_GetPropertyImagesServer) error {
	if err := stream.Send(&pb.Image{}); err != nil {
		return err
	}
	return nil
}

func main() {
	initDb()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPropertyPortalServer(grpcServer, &propertyPortalServer{})
	log.Printf("grpcServer starting")
	grpcServer.Serve(lis)
}

type Location struct {
	gorm.Model
	Rank       uint32
	District   string
	Population uint64
	Type       string
	County     string
	Region     string
}

func initDb() {
	db_file := "/tmp/property.db"
	os.Remove(db_file)
	db, err := gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	if err != nil {
		panic("Database connection failure")
	}

	db.AutoMigrate(&Location{})

	seedLocations(db)
}

func seedLocations(db *gorm.DB) {
	r := csv.NewReader(strings.NewReader(data))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var locations []Location

	for _, record := range records {
		rank, errRank := strconv.ParseUint(strings.ReplaceAll(record[0], ",", ""), 10, 32)
		if errRank != nil {
			log.Fatalf("failed to convert rank '%s'", record[0])
		}
		population, errPopulation := strconv.ParseUint(strings.ReplaceAll(record[2], ",", ""), 10, 64)
		if errPopulation != nil {
			log.Fatalf("failed to convert population '%s'", record[2])
		}
		location := Location{
			Rank:       uint32(rank),
			District:   record[1],
			Population: population,
			Type:       record[3],
			County:     record[4],
			Region:     record[5],
		}
		locations = append(locations, location)
	}

	db.Create(&locations)
}
