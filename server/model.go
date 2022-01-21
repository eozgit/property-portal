package main

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	Title           string
	Description     string
	Location        string
	Price           float64
	Beds            uint32
	Bathrooms       uint32
	PropertyType    uint32
	Garden          bool
	Parking         bool
	NewHome         bool
	CurrentRating   uint32
	PotentialRating uint32
}
