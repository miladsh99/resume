package utills

import (
	"fmt"
	"net/http"
	"resume/dto"
)

func ErrorManagement(w http.ResponseWriter, err *dto.ErrorHandle) {

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
	default:
		err.Message = "something went wrong"
	}

	fmt.Fprintf(w, "%v", err.Message)
}
