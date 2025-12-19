package handler

import (
	v10 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validationErrorResponse(c *fiber.Ctx, err error) error {
	errors := make(map[string]string)

	for _, e := range err.(v10.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		errors[field] = validationMessage(tag)
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errors": errors,
	})
}

func validationMessage(tag string) string {
	switch tag {
	case "required":
		return "is required"
	case "min":
		return "is too short"
	case "max":
		return "is too long"
	case "datetime":
		return "must be YYYY-MM-DD"
	default:
		return "is invalid"
	}
}
