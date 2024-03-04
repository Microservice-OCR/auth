package models

import (
	"time"
)

type Output struct {
	Token       string             `json:"token"`
	ConnectedAt time.Time          `json:"connectedAt"`
}
