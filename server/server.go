package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/eozgit/property-portal/propertyportal"
	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPropertyPortalServer(grpcServer, &propertyPortalServer{})
	log.Printf("grpcServer starting")
	grpcServer.Serve(lis)
}
