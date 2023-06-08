package userservice

import models "github.com/Dazzler/My-RestServer/pkg/models"

// UserService defines the service contract for user-related operations
type UserService interface {
	CreateUser(*models.User) error      // used to create a user
	Login(*models.User) (string, error) // used to authenticate a user and generate a JWT token
}
