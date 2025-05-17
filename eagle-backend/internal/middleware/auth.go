package middleware

import (
	"context"
	"net/http"

	"eagle-backend/internal/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Attach user to context if X-User-Id header is present and valid
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("X-User-Id")
		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing X-User-Id header"})
			c.Abort()
			return
		}
		objId, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
			c.Abort()
			return
		}
		coll := db.GetUserCollection()
		var user db.User
		err = coll.FindOne(context.Background(), map[string]interface{}{ "_id": objId }).Decode(&user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
