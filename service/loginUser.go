package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"resume/dto"
	"resume/entity"
	"resume/repository"
	"resume/utills"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utills.ErrorManagement(w, utills.InvalidMethod)
		return
	}

	data, dErr := io.ReadAll(r.Body)
	if dErr != nil {
		utills.ErrorManagement(w, utills.Body)
		return
	}

	req := dto.LoginRequest{}
	jErr := json.Unmarshal(data, &req)
	if jErr != nil {
		utills.ErrorManagement(w, utills.Unmarshal)
		return
	}

	invalidEmail := CheckRegexEmail(req.Email, w)
	if invalidEmail != req.Email {
		return
	}

	user := FindUserByEmail(entity.User{Email: req.Email}, repository.ConnectDatabase(), w)
	if user.ID == 0 {
		return
	}

	pErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password))
	if pErr != nil {
		utills.ErrorManagement(w, utills.Pass)
		return
	} else {
		res := Token(user)
		fmt.Fprintf(w, `{ "response" : %+v }`, res)
		return
	}

}
