package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"resume/repository"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "invalid request method")
		return
	}

	data, _ := io.ReadAll(r.Body)
	req := LoginRequest{}
	json.Unmarshal(data, &req)

	user, lErr := FindUserByEmail(req.Email, repository.ConnectDatabase())
	if lErr != nil {
		fmt.Fprintf(w, lErr.Error())
	}

	pErr := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(req.Password))
	if pErr != nil {
		fmt.Fprintf(w, "Wrong Password")
		return
	} else {
		res := Token(user)
		fmt.Fprintf(w, `{ "response" : %+v }`, res)
		return
	}

}

//		// ارسال توکن به کاربر
//		http.SetCookie(w, &http.Cookie{
//			Name:    "token",
//			Value:   tokenString,
//			Expires: expirationTime,
//		})
//		http.Redirect(w, r, "/data", http.StatusSeeOther)
//		return
//	}
//
//	w.WriteHeader(http.StatusUnauthorized)
//}
