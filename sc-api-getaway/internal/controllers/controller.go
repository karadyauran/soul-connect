package controllers

import "soul-connect/sc-api-getaway/internal/generated"

type Controller struct {
	AuthController *AuthController
}

func NewController(authClient generated.AuthServiceClient) *Controller {
	return &Controller{
		AuthController: NewAuthController(authClient),
	}
}
