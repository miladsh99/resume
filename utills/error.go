package utills

import (
	"fmt"
	"net/http"
)

type ErrorType int
type ErrorResponse string

func ErrorManagement(w http.ResponseWriter, errorType ErrorType) {

	var err ErrorResponse
	switch errorType {
	case Invalid:
		err = "your input is not email"
	case Exist:
		err = "This email has already been used"
	case NotExist:
		err = "user not found"
	case DB:
		err = "something went wrong"
	case InvalidMethod:
		err = "invalid request method"
	case Body:
		err = "The body cannot read"
	case Unmarshal:
		err = "The body cannot Unmarshal"
	case Pass:
		err = "Wrong Password"
	case Other:
		err = "something went wrong"
	default:
		err = "something went wrong"
	}

	fmt.Fprintf(w, "%v", err)
}
