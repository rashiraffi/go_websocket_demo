package main

import (
	"externalserver/internal/handler"
	"externalserver/internal/router"
	"externalserver/internal/service"
	"externalserver/pkg/logger"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log, err := logger.New(os.Getenv("ENV"))
	if err != nil {
		log.Fatal("Error loading logger")
	}

	log.Sugar().Info("Loger setup done")

	service := service.New()
	handler := handler.New(service)

	app := router.New(handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3001"
	}
	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Sugar().Fatalw("Error starting server", err.Error())
	}
}
