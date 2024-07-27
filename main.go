package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"project1/repository"
	"project1/service"
)

func main() {

	dsn := "root:dalim123@tcp(127.0.0.1:3306)/project1?parseTime=true"
	db, _ := repository.ConnectDB(dsn)
	defer db.Close()

	http.HandleFunc("/auth/register", service.RegisterUser(db))
	http.HandleFunc("/auth/login", service.LoginUser(db))

	err := http.ListenAndServe("localhost:7775", nil)
	if err != nil {
		fmt.Println(err)
	}

}
