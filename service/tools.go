package service

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"regexp"
	"resume/dto"
	"resume/entity"
	"resume/utills"
	"time"
)

func ReadRequest(r *http.Request, w http.ResponseWriter) entity.User {
	data, dErr := io.ReadAll(r.Body)
	if dErr != nil {
		utills.ErrorManagement(w, utills.Body)
		return entity.User{}
	}
	req := dto.RegisterRequest{}
	uErr := json.Unmarshal(data, &req)
	if uErr != nil {
		utills.ErrorManagement(w, utills.Unmarshal)
		return entity.User{}
	}
	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	pass, pErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if pErr != nil {
		utills.ErrorManagement(w, utills.Other)
		return entity.User{}
	}
	user.SetPassword(string(pass))
	return user
}

func CheckRegexEmail(input string, w http.ResponseWriter) string {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	if re.MatchString(input) {
		return input
	} else {
		utills.ErrorManagement(w, utills.Invalid)
		return ""
	}
}
