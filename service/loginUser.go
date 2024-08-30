package service

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"project1/dto"
	"project1/repository"
	"project1/utills"
)

func LoginUser(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {

		user, rErr := utills.ReadRequest(utills.Login, c)
		if rErr != nil {
			utills.ErrorManagement(c, rErr)
			return nil
		}

		//------------------------------------------

		invalidEmail := utills.CheckRegexEmail(user.Email)
		if invalidEmail != user.Email {
			utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.InvalidEmail})
			return nil
		}

		//------------------------------------------

		exUser, cErr := repository.FindUserByEmail(db, user.Email)
		if cErr.Type != utills.ExistEmail {
			utills.ErrorManagement(c, cErr)
			return nil
		}

		//------------------------------------------

		pErr := bcrypt.CompareHashAndPassword([]byte(exUser.GetPassword()), []byte(user.GetPassword()))
		if pErr != nil {
			utills.ErrorManagement(c, &dto.ErrorHandle{Type: utills.Pass})
			fmt.Println(pErr)
			return nil
		}

		//------------------------------------------

		res := utills.CreateToken(exUser)
		c.JSON(res)

		return nil
	}

}
