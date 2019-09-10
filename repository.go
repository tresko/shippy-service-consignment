package main

import (
	"context"
	pb "github.com/tresko/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"errors"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type MongoRepository struct {
	collection *mongo.Collection
}

// Create - create consignment
func (repo *MongoRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	// insert
	res, err := repo.collection.InsertOne(context.Background(), consignment)
	id := res.InsertedID

	// retrieve
	var consignmentsResult *pb.Consignment
	err = repo.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&consignmentsResult)
	if err != nil {
		return nil, err
	}
	return consignmentsResult, nil
}

// GetAll - get all consignments
func (repo *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repo.collection.Find(ctx, nil, nil)
	if err != nil {
		return nil, errors.New("Could not find data")
	}
	defer cur.Close(ctx)
	var consignments []*pb.Consignment
	for cur.Next(ctx) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
