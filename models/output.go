package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Output struct {
	Token       string             `json:"token"`
	ConnectedAt time.Time          `json:"connectedAt"`
	UserID      primitive.ObjectID `json:"userId,omitempty"`
}
