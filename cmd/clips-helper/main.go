package main

import (
	"clips/internal/app/commands"
	"clips/internal/config"
	"clips/internal/db"
	"clips/pkg/models"
	"fmt"
	"os"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	db.Setup()
	models.Setup(db.Database)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: clips-helper [command:new-user]")
		os.Exit(1)
	}

	command := args[1]
	commandArgs := args[1:]

	switch command {
	case "new-user":
		if err := commands.NewUser(commandArgs); err != nil {
			fmt.Printf("Failed to run %s: %s\n", command, err.Error())
			os.Exit(1)
		}
	}
}
