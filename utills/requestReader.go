package utills

import (
	"github.com/gofiber/fiber/v3"
	"project1/dto"
	"project1/entity"
)

func ReadRequest(ReType RequestType, c fiber.Ctx) (*entity.User, *dto.ErrorHandle) {

	switch ReType {
	case Register:
		var req = dto.RegisterRequest{}

		bErr := c.Bind().Body(&req)
		if bErr != nil {
			return nil, &dto.ErrorHandle{Type: Unmarshal}
		}

		user := entity.User{
			Name:  req.Name,
			Email: req.Email,
		}
		user.SetPassword(req.Password)

		return &user, nil

	case Login:
		var req = dto.LoginRequest{}
		bErr := c.Bind().Body(&req)
		if bErr != nil {
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
