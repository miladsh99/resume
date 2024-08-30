package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"project1/dto"
	"project1/repository"
	"project1/utills"
)

func RegisterUser(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {

		user, rErr := utills.ReadRequest(utills.Register, c)
		if rErr != nil {
			utills.ErrorManagement(c, rErr)
			return nil
		}

		//------------------------------------------

		_, cErr := repository.FindUserByEmail(db, user.Email)
		if cErr.Type != utills.NotExistEmail {
			utills.ErrorManagement(c, cErr)
			return nil
		}

		//------------------------------------------

		pass, pErr := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
		if pErr != nil {
			utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.Other})
			return nil
		}
		user.SetPassword(string(pass))

		//------------------------------------------

		registeredUser, iErr := repository.InsertUserDataInDB(db, user)
		if iErr != nil {
			utills.ErrorManagement(c, iErr)
			return nil
		}

		//------------------------------------------

		res := utills.CreateToken(registeredUser)
		c.JSON(res)

		return nil

	}
}
