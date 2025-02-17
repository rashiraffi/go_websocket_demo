package handler

import "github.com/gofiber/fiber/v2"

func (h *handler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	if email == "" || pass == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email and password are required",
		})
	}

	token, err := h.service.Login(email, pass)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
