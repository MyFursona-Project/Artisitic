package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Art struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      uint               `json:"userid,omitempty" bson:"userid,omitempty" binding:"required"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty" binding:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty" binding:"required"`
	Art         string             `json:"art,omitempty" bson:"art,omitempty" binding:"required"`
	Likes       uint               `json:"likes,omitempty" bson:"likes,omitempty" binding:"required"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty" binding:"required"`
}
