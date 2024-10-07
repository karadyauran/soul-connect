package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"soul-connect/sc-api-getaway/internal/generated"
)

type AuthController struct {
	client generated.AuthServiceClient
}

func NewAuthController(authClient generated.AuthServiceClient) *AuthController {
	return &AuthController{
		client: authClient,
	}
}

func (c *AuthController) Register(gc *gin.Context) {
	var req generated.RegisterUserRequest

	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()

	response, err := c.client.Register(ctx, &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, response)
}
