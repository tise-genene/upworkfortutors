package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"eagle-backend/internal/api"
	"eagle-backend/internal/middleware"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/api/register", api.RegisterUser)
	r.POST("/api/login", api.LoginUser)

	// Public endpoints
	r.GET("/api/jobs", api.ListJobs)
	r.GET("/api/applications", api.ListApplications)

	// Session endpoints (protected)
	r.POST("/api/sessions", middleware.AuthRequired(), api.CreateSession)
	r.GET("/api/sessions", middleware.AuthRequired(), api.ListSessions)

	// Protected endpoints
	r.POST("/api/jobs", middleware.AuthRequired(), api.CreateJob)
	r.POST("/api/applications", middleware.AuthRequired(), api.ApplyToJob)
	r.POST("/api/jobs/:jobId/hire", middleware.AuthRequired(), api.HireTutor)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on :%s...", port)
	r.Run(":" + port)
}
