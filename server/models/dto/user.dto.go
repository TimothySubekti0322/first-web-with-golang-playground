package dto

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	// Password  string         `json:"password"`
}

type AddUserRequest struct {
	Name      string         `json:"name" validate:"required"`
	Email     string         `json:"email" validate:"required,email"`
	Password  string         `json:"password" validate:"required,min=8,max=32"`
}

