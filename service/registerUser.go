package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"project1/dto"
	"project1/repository"
	"project1/utills"
)

func RegisterUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidMethod})
			return
		}

		//------------------------------------------

		user, rErr := utills.ReadRequest(utills.Register, r)
		if rErr != nil {
			utills.ErrorManagement(w, rErr)
			return
		}

		Email := utills.CheckRegexEmail(user.Email)
		if Email != user.Email {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidEmail})
			return
		}

		_, cErr := repository.FindUserByEmail(db, user.Email)
		if cErr.Type != utills.NotExistEmail {
			utills.ErrorManagement(w, cErr)
			return
		}

		//------------------------------------------

		registeredUser, iErr := repository.InsertUserDataInDB(db, user)
		if iErr != nil {
			utills.ErrorManagement(w, iErr)
			return
		}

		//------------------------------------------

		res := utills.CreateToken(registeredUser)
		fmt.Fprintf(w, `{ "response" : %+v }`, res)

	}

}
