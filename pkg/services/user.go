package services

import (
	"clips/pkg/config"
	"clips/pkg/data"
	"clips/pkg/db"
	"clips/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Signup(login data.Login) (Response,error) {

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(login.Password),config.BcryptCost)

	if err != nil {
		log.Printf("Failed to generate bcrypt hash: %s",err)
		return Response{Success: false, Status: 500, Message: "Unexpected Error"},err
	}

	user := models.User{}

	exists,err := db.Database.Model(&user).
		Where("username = ?",login.Username).
		Exists()

	if err != nil {
		log.Printf("Failed to find if user exists: %s",err)
		return Response{Success: false, Status: 500, Message: "Unexpected Error"},err
	}

	if exists {
		return Response{Success: false, Status: 409, Message: "User already exists!"},err
	}

	user.Username = login.Username
	user.Password = string(hashedPassword)

	_,err = db.Database.Model(&user).Insert()

	if err != nil {
		log.Printf("Failed to insert new user: %s",err)
		return Response{Success: false, Status: 500, Message: "Unexpected Error"},err
	}

	return Response{Success: true, Status: 200, Message: "Success", User: user},nil
}