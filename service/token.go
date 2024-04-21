package service

import (
	"github.com/golang-jwt/jwt/v5"
	"resume/entity"
	"time"
)

type TokenResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

var jwtKey = []byte("miladooo")

func Token(user entity.User) TokenResponse {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).UnixNano(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(jwtKey)

	res := TokenResponse{
		Token: tokenString,
		ID:    user.ID,
	}

	return res
}
