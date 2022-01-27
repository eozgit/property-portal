package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	pb "github.com/eozgit/property-portal/propertyportal"
	"google.golang.org/grpc"
)

var dal DataAccessLayer

type propertyPortalServer struct {
	pb.UnimplementedPropertyPortalServer
}

func (s *propertyPortalServer) FindProperties(ctx context.Context, filters *pb.Filters) (*pb.Properties, error) {
	props := dal.findProperties(filters)
	log.Printf("FindProperties: %+v", props)
	return props, nil
}

func (s *propertyPortalServer) GetPropertyDetails(ctx context.Context, property *pb.Property) (*pb.PropertyDetails, error) {
	prop := dal.getPropertyDetails(property)
	log.Printf("GetPropertyDetails: %+v", prop)
	return prop, nil
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

func (s *propertyPortalServer) GetPropertyImages(property *pb.Property, stream pb.PropertyPortal_GetPropertyImagesServer) error {
	sample := getSample(&rooms, 3, 5)
	url := gofakeit.URL()
	for _, room := range sample {
		if err := stream.Send(&pb.Image{
			Image: &pb.Image_Url{
				Url: fmt.Sprintf("%s/%s.png", url, room),
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dal = DataAccessLayer{}
	dal.initDb()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPropertyPortalServer(grpcServer, &propertyPortalServer{})
	log.Printf("grpcServer started")
	grpcServer.Serve(lis)
}
