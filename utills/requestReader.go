package utills

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"project1/dto"
	"project1/entity"
)

func ReadRequest(ReType RequestType, r *http.Request) (*entity.User, *dto.ErrorHandle) {

	data, rErr := io.ReadAll(r.Body)
	if rErr != nil {
		return nil, &dto.ErrorHandle{Type: Body}
	}

	switch ReType {
	case Register:
		var req = dto.RegisterRequest{}
		mErr := json.Unmarshal(data, &req)
		if mErr != nil {
			return nil, &dto.ErrorHandle{Type: Unmarshal}
		}
		user := entity.User{
			Name:  req.Name,
			Email: req.Email,
		}
		pass, pErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if pErr != nil {
			return nil, &dto.ErrorHandle{Type: Other}
		}
		user.SetPassword(string(pass))

		return &user, nil

	case Login:
		var req = dto.LoginRequest{}
		mErr := json.Unmarshal(data, &req)
		if mErr != nil {
			return nil, &dto.ErrorHandle{Type: Unmarshal}
		}
		user := entity.User{
			Email: req.Email,
		}
		user.SetPassword(req.Password)

		return &user, nil
	}
	return nil, &dto.ErrorHandle{Type: Other}
}
