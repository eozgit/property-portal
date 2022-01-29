package main

import (
	"context"
	"fmt"
	"io"
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
	log.Printf("Received FindProperties, filters: %+v\n\n", filters)
	props := dal.findProperties(filters)
	log.Printf("Sending properties: %+v\n\n\n", props)
	return props, nil
}

func (s *propertyPortalServer) GetPropertyDetails(ctx context.Context, property *pb.Property) (*pb.PropertyDetails, error) {
	log.Printf("Received GetPropertyDetails, property: %+v\n\n", property)
	prop := dal.getPropertyDetails(property)
	log.Printf("Sending property: %+v\n\n\n", prop)
	return prop, nil
}

var propertyImageIteration = map[uint]uint{}

func (s *propertyPortalServer) GetPropertyImages(stream pb.PropertyPortal_GetPropertyImagesServer) error {
	for {
		prop, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received GetPropertyImages, property: %+v\n\n", prop)

		propId := uint(prop.Id)
		iteration := uint(0)
		ok := false
		if iteration, ok = propertyImageIteration[propId]; !ok {
			propertyImageIteration[propId] = 1
		} else {
			propertyImageIteration[propId] += 1
		}

		images := dal.getPropertyImages(propId)

		index := iteration % uint(len(images))

		image := &images[index]

		log.Printf("Sending image: %+v\n\n\n", image)
		if err := stream.Send(image); err != nil {
			return err
		}
	}
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
	log.Printf("grpcServer started\n\n")
	grpcServer.Serve(lis)
}
