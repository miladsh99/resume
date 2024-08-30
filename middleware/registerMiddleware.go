package middleware

import (
	"github.com/gofiber/fiber/v3"
	"project1/dto"
	"project1/utills"
)

func ValidateRegister(c fiber.Ctx) error {

	user, rErr := utills.ReadRequest(utills.Register, c)
	if rErr != nil {
		utills.ErrorManagement(c, rErr)
		return nil
	}

	if len(user.Name) < 1 {
		utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.EmptyName})
		return nil
	}

	Email := utills.CheckRegexEmail(user.Email)
	if Email != user.Email {
		utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.InvalidEmail})
		return nil
	}

	if len(user.GetPassword()) < 8 {
		utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.ShortPass})
		return nil
	}

	return c.Next()
}
