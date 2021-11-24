package protocol

import (
	"hexagonal-template/config"
	"hexagonal-template/internal/core/service"
	"hexagonal-template/pkg/logger"
	"hexagonal-template/pkg/validators"
)

var app *application

type application struct {
	// svc stand for service
	svc *service.Service
	// pkg stand for package
	pkg packages
}

type packages struct {
	validator validators.Validator
}

func init() {
	logger.Init()
	config.Init()
	packages := packages{
		validator: validators.New(),
	}
	//todo: inject repository into the service
	app = &application{
		svc: service.New(nil),
		pkg: packages,
	}
}
