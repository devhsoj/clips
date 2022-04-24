package services

import (
	"clips/pkg/models"
)

// Response is a structure used in all services representing the success & information returned from these services.
type Response struct {
	Status  int
	Success bool
	Message string
	User    models.User
}
