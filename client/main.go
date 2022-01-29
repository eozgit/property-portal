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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	log.Printf("grpcClient started\n\n")

	cc := ClientContext{ctx, client}

	cc.findProperties()

	time.Sleep(time.Second * 3)

	cc.getPropertyDetails()

	time.Sleep(time.Second * 3)

	cc.getPropertyImages()
}
