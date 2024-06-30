package service

import (
	"fmt"
	"net/http"
	"resume/dto"
	"resume/repository"
	"resume/utills"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidMethod})
		return
	}

	//------------------------------------------

	user, rErr := ReadRegisterRequest(r)
	if rErr != nil {
		utills.ErrorManagement(w, rErr)
		return
	}

	Email := CheckRegexEmail(user.Email)
	if Email != user.Email {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidEmail})
		return
	}

	_, cErr := repository.CheckUserByEmail(user.Email)
	if cErr.Type != utills.NotExistEmail {
		utills.ErrorManagement(w, cErr)
		return
	}

	//------------------------------------------

	registeredUser, iErr := repository.InsertUserDataInDB(user)
	if iErr != nil {
		utills.ErrorManagement(w, iErr)
		return
	}

	//------------------------------------------

	res := CreateToken(registeredUser)
	fmt.Fprintf(w, `{ "response" : %+v }`, res)

}
