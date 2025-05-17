package api

import (
	"context"
	"net/http"
	"time"

	"eagle-backend/internal/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApplyJobRequest struct {
	JobID   string `json:"jobId" binding:"required"`
	TutorID string `json:"tutorId" binding:"required"`
	Message string `json:"message"`
}

type ApplyJobResponse struct {
	ID string `json:"id"`
}

func ApplyToJob(c *gin.Context) {
	var req ApplyJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jobObjID, err := primitive.ObjectIDFromHex(req.JobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jobId"})
		return
	}
	tutorObjID, err := primitive.ObjectIDFromHex(req.TutorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutorId"})
		return
	}
	app := db.Application{
		ID:          primitive.NewObjectID(),
		JobID:       jobObjID,
		TutorID:     tutorObjID,
		Message:     req.Message,
		Status:      db.AppPending,
		SubmittedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	coll := db.GetApplicationCollection()
	_, err = coll.InsertOne(context.Background(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to apply to job"})
		return
	}
	c.JSON(http.StatusCreated, ApplyJobResponse{ID: app.ID.Hex()})
}

// List applications for a job or tutor
func ListApplications(c *gin.Context) {
	jobId := c.Query("jobId")
	tutorId := c.Query("tutorId")
	filter := map[string]interface{}{}
	if jobId != "" {
		jobObjID, err := primitive.ObjectIDFromHex(jobId)
		if err == nil {
			filter["jobId"] = jobObjID
		}
	}
	if tutorId != "" {
		tutorObjID, err := primitive.ObjectIDFromHex(tutorId)
		if err == nil {
			filter["tutorId"] = tutorObjID
		}
	}
	coll := db.GetApplicationCollection()
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list applications"})
		return
	}
	defer cursor.Close(context.Background())
	apps := []db.Application{}
	for cursor.Next(context.Background()) {
		var app db.Application
		if err := cursor.Decode(&app); err == nil {
			apps = append(apps, app)
		}
	}
	c.JSON(http.StatusOK, apps)
}
