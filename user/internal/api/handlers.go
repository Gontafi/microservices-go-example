package api

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"user/internal/models"
	"user/pkg/utils"
)

func (rest *REST) RegisterUser(c *fiber.Ctx) error {
	var userBody struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	err := c.BodyParser(&userBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = rest.service.RegisterUser(
		c.Context(),
		models.User{
			Username: userBody.Username,
			Email:    userBody.Email,
			Password: userBody.Password,
		},
		models.UserProfile{
			FirstName: userBody.FirstName,
			LastName:  userBody.LastName,
		})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user successfully registered"})
}

func (rest *REST) LoginUser(c *fiber.Ctx) error {
	var userBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.BodyParser(&userBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := rest.service.LoginUser(c.Context(), userBody.Username, userBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"access_token": token})
}

func (rest *REST) ResetPassword(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Code     string `json:"code"`
	}

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = rest.service.ResetPassword(c.Context(), body.Username, body.Email, body.Password, body.Code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user password changed"})
}

func (rest *REST) ChangeUserProfile(c *fiber.Ctx) error {
	var body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	userIDString, ok := c.Locals("sub").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can not parse sub from fiber context"})
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = rest.service.ChangeUserProfile(c.Context(), models.UserProfile{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}, userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user profile changed"})
}

func (rest *REST) GetUserInfo(c *fiber.Ctx) error {
	var response struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	userIDString, ok := c.Locals("sub").(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can not parse sub from fiber context"})
	}

	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	profile, err := rest.service.GetUserProfile(c.Context(), userID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	response.FirstName = profile.FirstName
	response.LastName = profile.LastName

	return c.Status(fiber.StatusOK).JSON(response)
}
