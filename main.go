package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"resume/service"
)

func main() {

	http.HandleFunc("/auth/login", service.LoginUser)
	http.HandleFunc("/auth/register", service.RegisterUser)

	err := http.ListenAndServe("localhost:6665", nil)
	if err != nil {
		fmt.Println(err)
	}
}
