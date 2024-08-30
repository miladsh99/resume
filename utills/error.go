package utills

import (
	"github.com/gofiber/fiber/v3"
	"project1/dto"
)

func ErrorManagement(c fiber.Ctx, err *dto.ErrorHandle) {

	switch err.Type {
	case InvalidEmail:
		err.Message = "your input is not email"
	case ExistEmail:
		err.Message = "This email has already been used"
	case NotExistEmail:
		err.Message = "user not found"
	case DB:
		err.Message = "something went wrong , Database not found"
	case InvalidMethod:
		err.Message = "invalid request method"
	case Body:
		err.Message = "The body cannot read"
	case Unmarshal:
		err.Message = "The body cannot Unmarshal"
	case Pass:
		err.Message = "Wrong Password"
	case Other:
		err.Message = "something went wrong"
	case TokenErr:
		err.Message = "The token was not parsed"
	case NotReceiveToken:
		err.Message = "Token information not received"
	case InvalidToken:
		err.Message = "The token is invalid"
	case FailedGetDataFromDB:
		err.Message = "something went wrong , cannot get data from Database"
	case ShortPass:
		err.Message = " Password is Short"
	case EmptyName:
		err.Message = "Name is Empty"
	default:
		err.Message = "something went wrong"
	}
	c.SendString(err.Message)
}
