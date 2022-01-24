package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/eozgit/property-portal/propertyportal"
	"google.golang.org/grpc"
)

var dal DataAccessLayer

type propertyPortalServer struct {
	pb.UnimplementedPropertyPortalServer
}

func (s *propertyPortalServer) FindProperties(ctx context.Context, filters *pb.Filters) (*pb.Properties, error) {
	props := dal.findProperties(filters)
	log.Printf("FindProperties count: %d", len(props.Properties))
	return props, nil
}

func (s *propertyPortalServer) GetPropertyDetails(ctx context.Context, filters *pb.Property) (*pb.PropertyDetails, error) {
	return &pb.PropertyDetails{}, nil
}

func (s *propertyPortalServer) GetPropertyImages(rect *pb.Property, stream pb.PropertyPortal_GetPropertyImagesServer) error {
	if err := stream.Send(&pb.Image{}); err != nil {
		return err
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
