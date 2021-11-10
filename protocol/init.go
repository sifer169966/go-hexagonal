package protocol

import (
	"hexagonal-template/config"
	"hexagonal-template/internal/core/service"
	"hexagonal-template/pkg/logger"
	"hexagonal-template/pkg/validators"
)

// export `CfgPath` to set config path from cmd/root.go
var CfgPath string

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
	config.Init(CfgPath)
	packages := packages{
		validator: validators.New(),
	}
	//todo: inject repository into the service
	app = &application{
		svc: service.New(nil),
		pkg: packages,
	}
}
