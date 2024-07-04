package controllers

import (
	"fmt"

	"github.com/TimothySubekti0322/first-web-with-golang-playground/database"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/dto"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/entity"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := []entity.User{}

	database.DB.Find(&users)

	fmt.Println(users)

	userDTOs := make([]dto.UserResponse, len(users))

	for i, user := range users {
		userDTOs[i] = services.ConvertUserToUserResponse(user)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    userDTOs,
	})
}