package main

import (
	"context"
	"io"
	"log"

	pb "github.com/eozgit/property-portal/propertyportal"
)

type ClientContext struct {
	ctx    context.Context
	client pb.PropertyPortalClient
}

func (cc *ClientContext) findProperties() {
	properties, err := cc.client.FindProperties(cc.ctx, &pb.Filters{
		Location:     "Cambridge",
		MinPrice:     200000,
		MaxPrice:     350000,
		MinBeds:      1,
		MaxBeds:      3,
		PropertyType: pb.PropertyType_Any,
		MustHaves: &pb.Features{
			Garden: 1,
		},
	})
	if err != nil {
		log.Fatalf("%v.FindProperties(_) = _, %v: ", cc.client, err)
	}
	log.Printf("Search results: %+v\n", properties)
}

func (cc *ClientContext) getPropertyDetails() {
	propertyDetails, err := cc.client.GetPropertyDetails(cc.ctx, &pb.Property{
		Id: 1,
	})
	if err != nil {
		log.Fatalf("%v.GetPropertyDetails(_) = _, %v: ", cc.client, err)
	}
	log.Printf("Property details: %+v\n", propertyDetails)
}

func (cc *ClientContext) getPropertyImages() {
	stream, err := cc.client.GetPropertyImages(cc.ctx)
	if err != nil {
		log.Fatalf("%v.GetPropertyImages(_) = _, %v", cc.client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			image, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Image data: %v", image)
		}
	}()
	property := pb.Property{
		Id: 1,
	}
	for i := 0; i < 3; i++ {
		if err := stream.Send(&property); err != nil {
			log.Fatalf("Failed to send property info: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
