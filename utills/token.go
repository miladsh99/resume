package utills

import (
	"github.com/golang-jwt/jwt/v5"
	"project1/dto"
	"project1/entity"
	"time"
)

var jwtKey = []byte("miladooo")

func CreateToken(user *entity.User) dto.LoginResponse {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).UnixNano(),
	})

	tokenString, _ := token.SignedString(jwtKey)

	res := dto.LoginResponse{
		Token: tokenString,
		ID:    user.ID,
	}

	return res
}
