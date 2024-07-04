package services

import (
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/dto"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/entity"
)

func ConvertUserToUserResponse(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		// Password: user.Password,
	}
}