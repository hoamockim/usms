package controller

import (
	"github.com/gin-gonic/gin"
	"strings"
	"usms/app/dto"
	"usms/app/service"
	"usms/pkg/errors"
)

type UserProfileController struct {
	baseController
	userService interface {
		service.CommandUserService
		service.QueryUserService
	}
}

var userProfileController UserProfileController

func init() {
	userProfileController.userService = service.NewUserService()
}

func UserProfile() *UserProfileController {
	return &userProfileController
}

//add or update teacher info
//SaveUserInfo
func (ctrl *UserProfileController) SaveUserInfo(ctx *gin.Context) {
	var reqBody dto.UserInfoReq
	if err := ctx.Bind(&reqBody); err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "create-user-info")
	}
	userInfoRes, err := ctrl.userService.CreateUser(reqBody)
	if err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "create-user-info")
	}
	ctrl.success(ctx, userInfoRes)
}

//GetProfileDetail
func (ctrl *UserProfileController) GetProfileDetail(ctx *gin.Context) {
	userCode := ctx.Param("code")
	if strings.TrimSpace(userCode) == "" {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "user-info")
	}

	userInfoRes, err := ctrl.userService.GetUserInfo(userCode)
	if err != nil {
		ctrl.handleErr(ctx, &errors.ErrorMeta{HttpCode: 400}, "user-info")
	}
	ctrl.success(ctx, userInfoRes)
}
