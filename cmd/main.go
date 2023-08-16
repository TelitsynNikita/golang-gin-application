package main

import (
	"log"

	todo "github.com/TelitsynNikita/golang-gin-application"
	"github.com/TelitsynNikita/golang-gin-application/pkg/handler"
	"github.com/TelitsynNikita/golang-gin-application/pkg/repository"
	"github.com/TelitsynNikita/golang-gin-application/pkg/service"
)

func main() {
	repository := repository.NewRepository()
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("server failed: %s", err.Error())
	}
}
