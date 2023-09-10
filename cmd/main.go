package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/formulationapp/formulationapp/internal/controller"
	"github.com/formulationapp/formulationapp/internal/repository"
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	e := echo.New()

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.HideBanner = true
	e.HidePort = true
	log.Infof("Starting server on port %s", port)
	err = e.Start(fmt.Sprintf(":%s", port))
	log.Fatal(err)
}
