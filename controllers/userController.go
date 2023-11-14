package controllers

import (
	"khvnkay/golang-jwt-project/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollaction *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() {

}

func veriryPassword() {

}

func signUp() {

}
func Login() {

}

func GetUsers() {

}
func GetUser() {

}
