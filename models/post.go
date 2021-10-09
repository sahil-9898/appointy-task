package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID  `json:"_id"`
	User      primitive.ObjectID  `json:"user"`
	Caption   string              `json:"caption"`
	Image     string              `json:"image"`
	Timestamp primitive.Timestamp `json:"timestamp"`
}
