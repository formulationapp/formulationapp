package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/formulationapp/formulationapp/internal/controller"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/repository"
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file: %s", err)
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		logrus.Fatalf("Error opening database: %s", err)
	}

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				var appError dto.AppError
				switch {
				case errors.As(err, &appError):
					return echo.NewHTTPError(400, err.Error())
				}
			}
			return err
		}
	})

	config := dto.Config{SigningSecret: os.Getenv("SIGNING_SECRET")}

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e.HideBanner = true
	e.HidePort = true
	logrus.Infof("Starting server on port %s", port)
	err = e.Start(fmt.Sprintf(":%s", port))
	logrus.Fatal(err)
}
