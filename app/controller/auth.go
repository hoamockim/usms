package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usms/app/dto"
)

//SignIn create a jwt token for user's request
func SignIn(ctx *gin.Context) {
	var requestBody dto.JwtTokenBody
	if err := ctx.Bind(&requestBody); err != nil {
		error(ctx, http.StatusBadRequest, AUTH, "parsing-like-request", err.Error())
	}
	if !requestBody.Validate() {
		error(ctx, http.StatusBadRequest, AUTH, "validate", "request data is invalid")
	}
	resData, err := authService.SignIn(&requestBody)
	if err != nil {
		error(ctx, http.StatusInternalServerError, AUTH, "running generate token", err.Error())
	}
	success(ctx, &resData)
}

func SignInByThirdParty(ctx *gin.Context) {

}

// RefreshJwtToken renew a jwt token
func RefreshJwtToken(ctx *gin.Context) {

}
