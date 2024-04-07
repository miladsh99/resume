package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"resume/entity"
	"resume/repository"
	"resume/service"
	"time"
)

func main() {
	db := repository.ConnectDatabase()
	defer db.Close()

	http.HandleFunc("/auth/register", register)
	http.ListenAndServe(":7777", nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	db := repository.ConnectDatabase()
	defer db.Close()
	// name, email , password

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid request method")
		return
	}

	data, _ := io.ReadAll(r.Body)

	req := RegisterRequest{}
	json.Unmarshal(data, &req)

	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 7)
	user.SetPassword(string(pass))

	registeredUser, err := service.RegisterUser(db, user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": registeredUser.ID,
		"email":   registeredUser.Email,
		"exp":     time.Now().Add(time.Hour * 24).UnixNano(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString([]byte("test"))

	res := RegisterResponse{
		Token: tokenString,
		ID:    registeredUser.ID,
	}
	fmt.Fprintf(w, fmt.Sprintf(`{ "response" : %+v }`, res))

}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string
	ID    uint
}
