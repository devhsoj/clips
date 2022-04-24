package services

import (
	"clips/internal/data"
	"clips/internal/db"
	"clips/pkg/models"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// APIKeyAuth returns a Response generated from the results of trying to authenticate a user with the given apiKey.
func APIKeyAuth(apiKey string) Response {
	user := models.User{}

	err := db.Database.Model(&user).
		Where("api_key = ?",apiKey).
		Select()

	if err != nil {

		if err == pg.ErrNoRows {
			return Response{Success: false, Status: 401, Message: "Invalid Credentials"}
		}

		log.Printf("Failed to authenticate using API Key: %s",err)

		return Response{Success: false, Status: 500, Message: "Unexpected Error"}
	}

	return Response{Success: true, Status: 200, Message: "Success", User: user}
}

// UserAuth returns a Response generated from the results of trying to authenticate with the given data.Login.
func UserAuth(login data.Login) Response {
	user := models.User{}

	err := db.Database.Model(&user).
		Where("username = ?",login.Username).
		Select()

	if err != nil {

		if err == pg.ErrNoRows {
			return Response{Success: false, Status: 401, Message: "Invalid Credentials"}
		}

		log.Printf("Failed to query user: %s",err)

		return Response{Success: false, Status: 500, Message: "Unexpected Error"}
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(login.Password)); err != nil {
		return Response{Success: false, Status: 401, Message: "Invalid Credentials"}
	}

	return Response{Success: true, Status: 200, Message: "Success", User: user}
}