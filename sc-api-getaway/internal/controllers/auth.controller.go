package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"soul-connect/sc-api-getaway/internal/generated"
	"strings"
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

	if req.Username == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	if req.Email == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if req.Password == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
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

func (c *AuthController) Login(gc *gin.Context) {
	var req generated.LoginUserRequest

	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()

	response, err := c.client.Login(ctx, &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, response)
}

func (c *AuthController) Logout(gc *gin.Context) {
	authHeader := gc.GetHeader("Authorization")
	if authHeader == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
		return
	}

	ctx := context.Background()

	_, err := c.client.Logout(ctx, &generated.LogoutUserRequest{
		Token: token,
	})
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, "Successful logged out for user")
}

func (c *AuthController) LogoutFromAllDevices(gc *gin.Context) {
	var req generated.LogoutUserFromAllDevicesRequest
	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	ctx := context.Background()

	_, err := c.client.LogoutFromAllDevices(ctx, &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, "Successful logged out for user from all devices")
}

func (c *AuthController) DeleteUser(gc *gin.Context) {
	fmt.Println(gc.Params)

	userId := gc.Param("user_id")
	fmt.Println(userId)

	ctx := context.Background()

	_, err := c.client.DeleteUser(ctx, &generated.DeleteUserRequest{
		UserId: userId,
	})

	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, "Successful deleted user")
}
