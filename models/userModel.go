package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `bson:"_id"`
	Email         string    `json:"email" bson:"email"`
	Password      string    `json:"password" bson:"password"`
	Token         *string   `json:"token"`
	Refresh_token *string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
