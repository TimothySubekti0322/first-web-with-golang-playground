package controllers

import (
	"errors"

	"github.com/TimothySubekti0322/first-web-with-golang-playground/database"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/dto"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/models/entity"
	"github.com/TimothySubekti0322/first-web-with-golang-playground/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	body := dto.AuthRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidateStruct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user := entity.User{}

	err := database.DB.First(&user, "email = ?", body.Email).Error

	// Check Email
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check Password
	if !utils.CheckPassword(user.Password, body.Password) {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	// Create Token
	token, err := utils.GenerateJWT(user.Name, user.Email, user.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.AuthResponse{
		Token: token,
		Name:  user.Name,
		Email: user.Email,
	}

	// Check Password
	return c.Status(200).JSON(fiber.Map{
		"message" : "success",
		"data" : response,
		"status": 200,
	})
}

func Register(c *fiber.Ctx) error {
	body := dto.AddUserRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := utils.ValidateStruct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(body.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := entity.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPassword,
	}

	err = database.DB.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message" : "success",
		"data" : user,
		"status": 200,
	})
}