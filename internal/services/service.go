package services

import (
	"clips/pkg/models"
)

type Response struct {
	Status  int
	Success bool
	Message string
	User    models.User
}
