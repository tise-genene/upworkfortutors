package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobStatus string

const (
	JobOpen   JobStatus = "open"
	JobClosed JobStatus = "closed"
	JobHired  JobStatus = "hired"
)

type Job struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	ParentID    primitive.ObjectID   `bson:"parentId" json:"parentId"`
	Title       string               `bson:"title" json:"title"`
	Subject     string               `bson:"subject" json:"subject"`
	GradeLevel  string               `bson:"gradeLevel" json:"gradeLevel"`
	Schedule    string               `bson:"schedule" json:"schedule"`
	Location    string               `bson:"location" json:"location"`
	Description string               `bson:"description" json:"description"`
	Status      JobStatus            `bson:"status" json:"status"`
	Applicants  []primitive.ObjectID `bson:"applicants" json:"applicants"`
	HiredTutor  *primitive.ObjectID  `bson:"hiredTutor,omitempty" json:"hiredTutor,omitempty"`
	CreatedAt   primitive.DateTime   `bson:"createdAt" json:"createdAt"`
}

func GetJobCollection() *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		panic(err)
	}
	return client.Database(dbName).Collection("jobs")
}
