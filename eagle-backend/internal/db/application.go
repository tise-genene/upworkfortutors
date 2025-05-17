package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationStatus string

const (
	AppPending  ApplicationStatus = "pending"
	AppAccepted ApplicationStatus = "accepted"
	AppRejected ApplicationStatus = "rejected"
)

type Application struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	JobID      primitive.ObjectID `bson:"jobId" json:"jobId"`
	TutorID    primitive.ObjectID `bson:"tutorId" json:"tutorId"`
	Message    string             `bson:"message" json:"message"`
	Status     ApplicationStatus  `bson:"status" json:"status"`
	SubmittedAt primitive.DateTime `bson:"submittedAt" json:"submittedAt"`
}

func GetApplicationCollection() *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}
	return client.Database(dbName).Collection("applications")
}
