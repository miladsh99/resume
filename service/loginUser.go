package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"resume/dto"
	"resume/repository"
	"resume/utills"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidMethod})
		return
	}

	//------------------------------------------

	var req = dto.LoginRequest{}

	data, dErr := io.ReadAll(r.Body)
	if dErr != nil {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.Body})
		return
	}

	jErr := json.Unmarshal(data, &req)
	if jErr != nil {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.Unmarshal})
		return
	}

	//------------------------------------------

	invalidEmail := CheckRegexEmail(req.Email)
	if invalidEmail != req.Email {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.InvalidEmail})
		return
	}

	//------------------------------------------

	user, cErr := repository.CheckUserByEmail(req.Email)
	if cErr.Type != utills.ExistEmail {
		utills.ErrorManagement(w, cErr)
		return
	}

	//------------------------------------------

	pErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password))
	if pErr != nil {
		utills.ErrorManagement(w, &dto.ErrorHandle{Type: utills.Pass})
		return
	}

	//------------------------------------------

	res := CreateToken(user)
	fmt.Fprintf(w, `{ "response" : %+v }`, res)

}
