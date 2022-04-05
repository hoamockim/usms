package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usms/app/dto"
	"usms/app/service"
	"usms/pkg/errors"
)

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

const (
	USERPROFILE = "usms-profile"
	AUTH        = "usms-auth"
)

var authService service.AuthService
var userService interface {
	service.CommandUserService
	service.QueryUserService
}

func init() {

	authService = service.GetAuthService()
	userService = service.GetUserService()
}

func success(ctx *gin.Context, data interface{}) {
	res := &dto.ResData{Data: data}
	ctx.JSON(http.StatusOK, res)
}

func error(ctx *gin.Context, httpCode int, serviceName, errorCode, errorMessage string) {
	appErr := errors.New(serviceName, errorCode, errorMessage)
	ctx.JSON(httpCode, appErr)
}
