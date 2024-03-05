package models

type Invalid struct {
	Token     string `bson:"token"`
	Invalid   bool  `bson:"invalid"`
}