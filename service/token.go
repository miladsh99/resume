package service

import (
	"github.com/golang-jwt/jwt/v5"
	"resume/dto"
	"resume/entity"
	"time"
)

var jwtKey = []byte("miladooo")

func Token(user entity.User) dto.LoginResponse {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).UnixNano(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(jwtKey)

	res := dto.LoginResponse{
		Token: tokenString,
		ID:    user.ID,
	}

	return res
}
