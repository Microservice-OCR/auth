package models

import "time"

type Output struct {
	Token       string `bson:"token"`
	ConnectedAt time.Time
}
