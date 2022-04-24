package commands

import (
	"clips/internal/config"
	"clips/internal/db"
	"clips/pkg/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// NewUser takes either two sys arguments or inputs (username and password), and creates a user in the database
func NewUser(args []string) error {

	var username string
	var password string

	if len(args) != 2 {
		fmt.Print("Enter username: ")
		if _,err := fmt.Scanln(&username); err != nil {
			return err
		}

		fmt.Print("Enter password: ")
		if _,err := fmt.Scanln(&password); err != nil {
			return err
		}
	} else {
		username = args[0]
		password = args[1]
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
