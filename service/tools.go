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

func ReadRegisterRequest(r *http.Request) (*entity.User, *dto.ErrorHandle) {

	var req = dto.RegisterRequest{}

	data, rErr := io.ReadAll(r.Body)
	if rErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.Body}
	}

	mErr := json.Unmarshal(data, &req)
	if mErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.Unmarshal}
	}

	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Time{}.UTC(),
		UpdatedAt: time.Time{}.UTC(),
	}
	pass, pErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if pErr != nil {
		return nil, &dto.ErrorHandle{Type: utills.Other}
	}
	user.SetPassword(string(pass))

	return &user, nil
}

func CheckRegexEmail(input string) string {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	if re.MatchString(input) {
		return input
	} else {
		return ""
	}
}
