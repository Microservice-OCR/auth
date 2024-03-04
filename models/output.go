package models

import "time"

type Output struct {
    ID          string    `json:"id"`
    Token       string    `json:"token"`
    ConnectedAt time.Time `json:"connectedAt"`
}
