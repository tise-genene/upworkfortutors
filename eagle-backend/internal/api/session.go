package api

import (
	"context"
	"net/http"
	"time"

	"eagle-backend/internal/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateSessionRequest struct {
	JobID       string `json:"jobId" binding:"required"`
	TutorID     string `json:"tutorId" binding:"required"`
	ParentID    string `json:"parentId" binding:"required"`
	SessionDate string `json:"sessionDate" binding:"required"` // ISO8601
	Duration    int    `json:"duration" binding:"required"`
	ProofOfSession string `json:"proofOfSession"`
	Feedback    string `json:"feedback"`
}

type CreateSessionResponse struct {
	ID string `json:"id"`
}

func CreateSession(c *gin.Context) {
	var req CreateSessionRequest
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
	parentObjID, err := primitive.ObjectIDFromHex(req.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parentId"})
		return
	}
	sessionTime, err := time.Parse(time.RFC3339, req.SessionDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sessionDate"})
		return
	}
	session := db.Session{
		ID:            primitive.NewObjectID(),
		JobID:         jobObjID,
		TutorID:       tutorObjID,
		ParentID:      parentObjID,
		SessionDate:   primitive.NewDateTimeFromTime(sessionTime),
		Duration:      req.Duration,
		ProofOfSession: req.ProofOfSession,
		Feedback:      req.Feedback,
	}
	coll := db.GetSessionCollection()
	_, err = coll.InsertOne(context.Background(), session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}
	c.JSON(http.StatusCreated, CreateSessionResponse{ID: session.ID.Hex()})
}

func ListSessions(c *gin.Context) {
	jobId := c.Query("jobId")
	parentId := c.Query("parentId")
	tutorId := c.Query("tutorId")
	filter := map[string]interface{}{}
	if jobId != "" {
		jobObjID, err := primitive.ObjectIDFromHex(jobId)
		if err == nil {
			filter["jobId"] = jobObjID
		}
	}
	if parentId != "" {
		parentObjID, err := primitive.ObjectIDFromHex(parentId)
		if err == nil {
			filter["parentId"] = parentObjID
		}
	}
	if tutorId != "" {
		tutorObjID, err := primitive.ObjectIDFromHex(tutorId)
		if err == nil {
			filter["tutorId"] = tutorObjID
		}
	}
	coll := db.GetSessionCollection()
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list sessions"})
		return
	}
	defer cursor.Close(context.Background())
	sessions := []db.Session{}
	for cursor.Next(context.Background()) {
		var session db.Session
		if err := cursor.Decode(&session); err == nil {
			sessions = append(sessions, session)
		}
	}
	c.JSON(http.StatusOK, sessions)
}
