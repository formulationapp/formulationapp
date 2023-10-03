package main

import (
	"embed"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"net/http"
	"os"

	"github.com/formulationapp/formulationapp/internal/controller"
	"github.com/formulationapp/formulationapp/internal/dto"
	"github.com/formulationapp/formulationapp/internal/repository"
	"github.com/formulationapp/formulationapp/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed public
var public embed.FS

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatalf("Error opening database: %s", err)
	}

	e := echo.New()
	e.Use(middleware.CORS())
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

	e.GET("/", echo.NotFoundHandler, middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public/landing",
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		Filesystem: http.FS(public),
	}))
	e.GET("/*", echo.NotFoundHandler, middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "public",
		Index:      "index.html",
		Browse:     false,
		HTML5:      true,
		Filesystem: http.FS(public),
	}))

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
