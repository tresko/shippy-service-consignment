package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"github.com/micro/go-micro"
	// Import the generated protobuf code
	pb "github.com/tresko/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/tresko/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.consignment.service"),
	)

	// Init will parse the command line flags.
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())
	consignmentCollection := client.Database("shippy").Collection("consignment")

	repository := &MongoRepository{consignmentCollection}

	// Vesel service
	vesselClient := vesselProto.NewVesselService("shippy.vessel.service", srv.Client())

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &ShippingService{repository, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
