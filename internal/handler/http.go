package handler

import (
	"hexagonal-template/internal/core/domain"
	"hexagonal-template/internal/core/port"
	"hexagonal-template/pkg/validators"

	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	svc        port.Service
	validators validators.Validators
}

func NewHttp(svc port.Service, validators validators.Validators) *HttpHandler {
	return &HttpHandler{
		svc:        svc,
		validators: validators,
	}
}

/*
	|--------------------------------------------------------------------------
	| The Handler Adaptor
	|--------------------------------------------------------------------------
	|
	| An Adapter will initiate the interaction with the Application through
	| a Port, using specific technology that means you can choose
	| any technology you want for your application protocol.
	|
*/

func (hdl *HttpHandler) SomeHandler(c *fiber.Ctx) error {
	var request domain.BusinessLogicRequest

	/*
		|--------------------------------------------------------------------------
		| Request Body Serialization
		|--------------------------------------------------------------------------
		|
		| Here you can parse the body from incoming request into the structure/model
		| to use in your application.
		|
	*/

	/*
		|--------------------------------------------------------------------------
		| Data Validation
		|--------------------------------------------------------------------------
		|
		| Here you may specify which part of the incoming request body you want to validate
		| before putting them into the business logic.
		|
	*/

	err := hdl.svc.SomeBusinessLogic(request)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
