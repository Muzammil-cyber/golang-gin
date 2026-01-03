package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials entity.LoginCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return ""
	}

	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if !isAuthenticated {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return ""
	}

	token := c.jwtService.GenerateToken(credentials.Username, credentials.Username == "admin")
	return token
}
