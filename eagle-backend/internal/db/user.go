package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	RoleParent UserRole = "parent"
	RoleTutor  UserRole = "tutor"
)

type Profile struct {
	Bio        string   `bson:"bio,omitempty" json:"bio,omitempty"`
	Photo      string   `bson:"photo,omitempty" json:"photo,omitempty"`
	Subjects   []string `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Experience string   `bson:"experience,omitempty" json:"experience,omitempty"`
	GradeLevels []string `bson:"gradeLevels,omitempty" json:"gradeLevels,omitempty"`
	Location   string   `bson:"location,omitempty" json:"location,omitempty"`
	Rate       float64  `bson:"rate,omitempty" json:"rate,omitempty"`
	Rating     float64  `bson:"rating,omitempty" json:"rating,omitempty"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Phone     string             `bson:"phone" json:"phone"`
	Role      UserRole           `bson:"role" json:"role"`
	Profile   Profile            `bson:"profile" json:"profile"`
	IsVerified bool              `bson:"isVerified" json:"isVerified"`
	CreatedAt primitive.DateTime `bson:"createdAt" json:"createdAt"`
}
