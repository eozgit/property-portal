package main

import (
	"context"
	"log"
	"time"

	pb "github.com/eozgit/property-portal/propertyportal"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPropertyPortalClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cc := ClientContext{ctx, client}

	cc.findProperties()

	cc.getPropertyDetails()

	cc.getPropertyImages()
}
