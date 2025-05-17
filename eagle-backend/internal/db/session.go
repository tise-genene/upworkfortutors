package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Session struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	JobID         primitive.ObjectID `bson:"jobId" json:"jobId"`
	TutorID       primitive.ObjectID `bson:"tutorId" json:"tutorId"`
	ParentID      primitive.ObjectID `bson:"parentId" json:"parentId"`
	SessionDate   primitive.DateTime `bson:"sessionDate" json:"sessionDate"`
	Duration      int                `bson:"duration" json:"duration"` // in minutes
	Feedback      string             `bson:"feedback,omitempty" json:"feedback,omitempty"`
	ProofOfSession string            `bson:"proofOfSession,omitempty" json:"proofOfSession,omitempty"`
}

func GetSessionCollection() *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}
	return client.Database(dbName).Collection("sessions")
}
