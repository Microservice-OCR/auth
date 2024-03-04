package models

type Output struct {
    ID          string    `json:"id"`
    Token       string    `json:"token"`
    ConnectedAt int64 `json:"connectedAt"`
}
