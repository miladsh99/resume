package service

import (
	"fmt"
	"net/http"
	"resume/repository"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid request method")
		return
	}

	user := ReadRequest(r)

	exist := CheckEmail(user.Email, repository.ConnectDatabase())

	switch exist {
	case 1:
		fmt.Fprintf(w, "something went wrong")
		return
	case 2:
		fmt.Fprintf(w, "This email has already been used")
		return
	case 3:
		fmt.Fprintf(w, "your input is not email")
		return
	}

	registeredUser, rErr := InsertUserDataInDB(user)
	if rErr != nil {
		fmt.Fprintf(w, rErr.Error())
	}

	res := Token(registeredUser)
	fmt.Fprintf(w, `{ "response" : %+v }`, res)
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
