package models

import (
	"time"
)

type User struct {
	// ID           primitive.ObjectID `bson:"_id"`
	FirstName    string    `json:"first_name" bson:"FirstName" validate:"required, min=2, max=10"`
	LastName     string    `json:"last_name" validate:"required"`
	Password     string    `json:"password" validate:"required, min=6"` // You can add more validations for passwords
	Email        string    `json:"email" validate:"email, required"`    // It will check email kind of validation
	Phone        string    `json:"phone" validate:"required"`
	Token        string    `json:"token" validate:"required"`
	UserType     string    `json:"user_type" validate:"required, eq=ADMIN|eq=USER"` // this validation is just like enums. Any one option will be chosen
	RefreshToken string    `json:"refresh_token"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	User_id      string    `json:"user_id"`
}
