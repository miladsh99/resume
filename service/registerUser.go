package service

import (
	"fmt"
	"net/http"
	"resume/repository"
	"resume/utills"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utills.ErrorManagement(w, utills.InvalidMethod)
		return
	}

	user := ReadRequest(r, w)

	if user.Name == "" {
		return
	}

	invalidEmail := CheckRegexEmail(user.Email, w)
	if invalidEmail != user.Email {
		return
	}

	cErr := CheckEmail(user.Email, repository.ConnectDatabase(), w)
	if cErr != utills.NotExist {
		utills.ErrorManagement(w, cErr)
		return
	}

	registeredUser, rErr := InsertUserDataInDB(user)
	if rErr != nil {
		utills.ErrorManagement(w, utills.DB)
	}

	res := Token(registeredUser)
	fmt.Fprintf(w, `{ "response" : %+v }`, res)
}
