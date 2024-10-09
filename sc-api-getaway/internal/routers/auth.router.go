package routers

import (
	"github.com/gin-gonic/gin"
	"soul-connect/sc-api-getaway/internal/config"
	"soul-connect/sc-api-getaway/internal/controllers"
)

type authRouter struct {
	authController *controllers.AuthController
	config         *config.Config
}

func newAuthRouter(authController *controllers.AuthController, config *config.Config) *authRouter {
	return &authRouter{authController, config}
}

func (ar *authRouter) setAuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("auth")
	router.POST("/register", ar.authController.Register)
	router.POST("/login", ar.authController.Login)
	router.POST("/logout", ar.authController.Logout)
	router.POST("/logout-from-all-devices", ar.authController.LogoutFromAllDevices)
	router.DELETE("/delete-user/:user_id", ar.authController.DeleteUser)
}
