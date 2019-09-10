package main

import (
	"context"
	"log"
	// Import the generated protobuf code
	pb "github.com/tresko/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/tresko/shippy-service-vessel/proto/vessel"
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type ShippingService struct {
	repo repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *ShippingService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *ShippingService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}
