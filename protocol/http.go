package protocol

import (
	"hexagonal-template/config"
	"hexagonal-template/internal/core/service"
	"hexagonal-template/internal/handler"
	"hexagonal-template/internal/repository"
	"hexagonal-template/pkg/validators"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

/*
	|--------------------------------------------------------------------------
	| Application Protocol
	|--------------------------------------------------------------------------
	|
	| Here you can choose which protocol your application wants to interact
	| with the client for instance HTTP, gRPC etc.
	|
*/

// The example to serve REST
func ServeREST() error {
	app := fiber.New(fiber.Config{
		DisableKeepalive: false,
	})

	repo := repository.NewPostgres()
	svc := service.New(repo)
	vld := validators.New()
	hdl := handler.NewHttp(svc, vld)

	// example
	app.Get("/", hdl.SomeHandler)

	// gracefully shuts down  ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down ...")
		err := app.Shutdown()
		if err != nil {
			log.Println(err)
		}
		os.Exit(0)
	}()
	err := app.Listen(":" + config.GetConfig().App.HTTPPort)
	if err != nil {
		return err
	}
	return nil
}
