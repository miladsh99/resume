package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"resume/service"
)

func main() {
	http.HandleFunc("/auth/register", service.RegisterUser)
	http.HandleFunc("/auth/login", service.LoginUser)

	err := http.ListenAndServe("localhost:7775", nil)
	if err != nil {
		fmt.Println(err)
	}
}
