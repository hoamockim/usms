package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//GetProfileDetail get user info
func GetProfileDetail(ctx *gin.Context) {
	userCode := ctx.Param("code")
	if strings.TrimSpace(userCode) == "" {
		error(ctx, http.StatusBadRequest, USERPROFILE, "parsing-request", "user code is empty")
	}

	userInfoRes, err := userService.GetUserInfo(userCode)
	if err != nil {
		error(ctx, http.StatusInternalServerError, USERPROFILE, "create-user", err.Error())
	}
	success(ctx, userInfoRes)
}

func SignUp(ctx *gin.Context) {

}
