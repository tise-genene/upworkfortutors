package api

import (
	"context"
	"net/http"
	"time"

	"eagle-backend/internal/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterUserRequest struct {
	Name    string         `json:"name" binding:"required"`
	Phone   string         `json:"phone" binding:"required"`
	Role    db.UserRole    `json:"role" binding:"required,oneof=parent tutor"`
	Profile db.Profile     `json:"profile"`
}

type RegisterUserResponse struct {
	ID string `json:"id"`
}

func RegisterUser(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := db.User{
		ID:         primitive.NewObjectID(),
		Name:       req.Name,
		Phone:      req.Phone,
		Role:       req.Role,
		Profile:    req.Profile,
		IsVerified: false,
		CreatedAt:  primitive.NewDateTimeFromTime(time.Now()),
	}
	coll := db.GetUserCollection()
	_, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, RegisterUserResponse{ID: user.ID.Hex()})
}

// Simple phone-based login (no OTP)
type LoginRequest struct {
	Phone string `json:"phone" binding:"required"`
}

type LoginResponse struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Role    db.UserRole `json:"role"`
	Profile db.Profile  `json:"profile"`
}

func LoginUser(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coll := db.GetUserCollection()
	var user db.User
	err := coll.FindOne(context.Background(), map[string]interface{}{ "phone": req.Phone }).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, LoginResponse{
		ID:      user.ID.Hex(),
		Name:    user.Name,
		Role:    user.Role,
		Profile: user.Profile,
	})
}
