package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	"log"
	"project1/config"
	"project1/middleware"
	"project1/repository"
	"project1/service"
)

func main() {

	app := fiber.New()

	conFile, rErr := config.ReadConfig("config/database.yaml")
	if rErr != nil {
		log.Fatalf("Error reading config file: %v", rErr)
	}

	dsn := config.GenerateConfig(conFile)
	db, _ := repository.ConnectDB(dsn)
	defer db.Close()

	app.Post("/auth/register", service.RegisterUser(db), middleware.ValidateRegister)
	app.Get("/auth/login", service.LoginUser(db))

	app.Listen(":7775")

}
