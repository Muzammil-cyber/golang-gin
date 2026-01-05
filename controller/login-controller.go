package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzammil-cyber/golang-gin/dto"
	"github.com/muzammil-cyber/golang-gin/entity"
	"github.com/muzammil-cyber/golang-gin/service"
	"github.com/muzammil-cyber/golang-gin/utils"

	_ "github.com/muzammil-cyber/golang-gin/docs"
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

// Login godoc
// @Summary User Login
// @Description Authenticate user with username and password to receive a JWT token for accessing protected endpoints
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body entity.LoginCredentials true "User login credentials (username and password)"
// @Success 200 {object} dto.LoginResponse "Successfully authenticated, returns JWT token"
// @Failure 400 {object} dto.ValidationErrorResponse "Invalid request format or missing required fields"
// @Failure 401 {object} dto.ErrorResponse "Authentication failed - invalid username or password"
// @Router /auth/login [post]
func (c *loginController) Login(ctx *gin.Context) string {
	var credentials entity.LoginCredentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ValidationErrorResponse{Errors: utils.FormatValidationError(err)})
		return ""
	}

	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if !isAuthenticated {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Invalid username or password"})
		return ""
	}

	token := c.jwtService.GenerateToken(credentials.Username, credentials.Username == "admin")
	return token
}
