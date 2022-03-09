package controller

import (
	"github.com/gin-gonic/gin"
	"usms/app/dto"
	"usms/app/service"
	"usms/pkg/errors"
)

type AuthController struct {
	baseController
	authService service.AuthService
}

var (
	authController AuthController
)

func init() {
	authController.authService = service.NewAuthService()
}

func GetAuthController() *AuthController {
	return &authController
}

//CreateJwtToken create a jwt token for user's request
func (ctrl *AuthController) CreateJwtToken(ctx *gin.Context) {
	var requestBody dto.JwtTokenRequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "create-jwt-token")
	}
	if !requestBody.ValidateCreateJwtToken() {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "create-jwt-token")
	}
	resData, err := ctrl.authService.RequestToken(&requestBody)
	if err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "create-jwt-token")
	}
	ctrl.success(ctx, &resData)
}

func (ctrl *AuthController) RefreshJwtToken(ctx *gin.Context) {
	var requestBody dto.JwtTokenRequestBody
	if err := ctx.Bind(&requestBody); err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "refresh-jwt-token")
	}
	if !requestBody.ValidateRefreshJwtToken() {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "refresh-jwt-token")
	}
}
