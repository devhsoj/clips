package commands

import (
	"clips/internal/config"
	"clips/internal/db"
	"clips/pkg/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(args []string) error {

	var username string
	var password string

	fmt.Print("Enter username: ")
	if _,err := fmt.Scanln(&username); err != nil {
		return err
	}

	fmt.Print("Enter password: ")
	if _,err := fmt.Scanln(&password); err != nil {
		return err
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),config.BcryptCost)

	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	 if _,err = db.Database.Model(&user).Insert(); err != nil {
		 return err
	 }

	return nil
}
