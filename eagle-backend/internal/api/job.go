package api

import (
	"context"
	"net/http"
	"time"

	"eagle-backend/internal/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateJobRequest struct {
	ParentID    string `json:"parentId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Subject     string `json:"subject" binding:"required"`
	GradeLevel  string `json:"gradeLevel" binding:"required"`
	Schedule    string `json:"schedule"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type JobResponse struct {
	ID          string    `json:"id"`
	ParentID    string    `json:"parentId"`
	Title       string    `json:"title"`
	Subject     string    `json:"subject"`
	GradeLevel  string    `json:"gradeLevel"`
	Schedule    string    `json:"schedule"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Applicants  []string  `json:"applicants"`
	HiredTutor  *string   `json:"hiredTutor,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
}

type CreateJobResponse struct {
	ID string `json:"id"`
}

func CreateJob(c *gin.Context) {
	var req CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parentObjID, err := primitive.ObjectIDFromHex(req.ParentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parentId"})
		return
	}
	job := db.Job{
		ID:          primitive.NewObjectID(),
		ParentID:    parentObjID,
		Title:       req.Title,
		Subject:     req.Subject,
		GradeLevel:  req.GradeLevel,
		Schedule:    req.Schedule,
		Location:    req.Location,
		Description: req.Description,
		Status:      db.JobOpen,
		Applicants:  []primitive.ObjectID{},
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	coll := db.GetJobCollection()
	_, err = coll.InsertOne(context.Background(), job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create job"})
		return
	}
	c.JSON(http.StatusCreated, CreateJobResponse{ID: job.ID.Hex()})
}

func ListJobs(c *gin.Context) {
	coll := db.GetJobCollection()
	cursor, err := coll.Find(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseJobs []JobResponse
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var job db.Job
			if err := cursor.Decode(&job); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Convert MongoDB types to plain objects
			responseJob := JobResponse{
				ID:          job.ID.Hex(),
				ParentID:    job.ParentID.Hex(),
				Title:       job.Title,
				Subject:     job.Subject,
				GradeLevel:  job.GradeLevel,
				Schedule:    job.Schedule,
				Location:    job.Location,
				Description: job.Description,
				Status:      string(job.Status),
				Applicants:  convertObjectIDsToStrings(job.Applicants),
				HiredTutor:  convertObjectIDToString(job.HiredTutor),
				CreatedAt:   job.CreatedAt.Time(),
			}
			responseJobs = append(responseJobs, responseJob)
		}
	}()
	<-done

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, responseJobs)
}

func convertObjectIDsToStrings(objectIDs []primitive.ObjectID) []string {
	var result []string
	for _, objectID := range objectIDs {
		result = append(result, objectID.Hex())
	}
	return result
}

func convertObjectIDToString(objectID *primitive.ObjectID) *string {
	if objectID == nil {
		return nil
	}
	hex := objectID.Hex()
	return &hex
}

// Hire a tutor for a job
// POST /api/jobs/:jobId/hire { "tutorId": "..." }
type HireTutorRequest struct {
	TutorID string `json:"tutorId" binding:"required"`
}

type HireTutorResponse struct {
	Success bool `json:"success"`
}

func HireTutor(c *gin.Context) {
	jobId := c.Param("jobId")
	tutorId := ""
	var req HireTutorRequest
	if err := c.ShouldBindJSON(&req); err == nil {
		tutorId = req.TutorID
	}
	jobObjID, err := primitive.ObjectIDFromHex(jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jobId"})
		return
	}
	tutorObjID, err := primitive.ObjectIDFromHex(tutorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tutorId"})
		return
	}
	coll := db.GetJobCollection()
	update := map[string]interface{}{
		"status":     db.JobHired,
		"hiredTutor": tutorObjID,
	}
	_, err = coll.UpdateOne(context.Background(), map[string]interface{}{"_id": jobObjID}, map[string]interface{}{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hire tutor"})
		return
	}
	c.JSON(http.StatusOK, HireTutorResponse{Success: true})
}
