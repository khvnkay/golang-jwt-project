package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate: "required,min=2, max=100"`
	Last_name     *string            `json:"last_name" validate: "required,min=2, max=100"`
	Password      *string            `json:"password" validate: "required,min=6"`
	Email         *string            `json:"password" validate: "email ,required"`
	Phone         *string            `json:"phone" validate: "required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate: "required, eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Create_at     time.Time          `json:"create_at"`
	Update_at     time.Time          `json:"update_at"`
	User_id       string             `json:"user_id"`
}

func models() {

}
