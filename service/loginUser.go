package service

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"project1/dto"
	"project1/repository"
	"project1/utills"
)

func LoginUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidMethod})
			return
		}

		//------------------------------------------

		user, rErr := utills.ReadRequest(utills.Login, r)
		if rErr != nil {
			utills.ErrorManagement(w, rErr)
			return
		}

		//------------------------------------------

		invalidEmail := utills.CheckRegexEmail(user.Email)
		if invalidEmail != user.Email {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidEmail})
			return
		}

		//------------------------------------------

		exUser, cErr := repository.FindUserByEmail(db, user.Email)
		if cErr.Type != utills.ExistEmail {
			utills.ErrorManagement(w, cErr)
			return
		}

		//------------------------------------------

		pErr := bcrypt.CompareHashAndPassword([]byte(exUser.GetPassword()), []byte(user.GetPassword()))
		if pErr != nil {
			utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.Pass})
			fmt.Println(pErr)
			return
		}

		//------------------------------------------

		res := utills.CreateToken(exUser)
		fmt.Fprintf(w, `{ "response" : %+v }`, res)
	}

}
