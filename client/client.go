package main

import (
	"context"
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
