package service

import (
	"errors"
	"usms/app/dto"
	"usms/db/models"
	"usms/db/repositories"
	"usms/pkg/util"
)

type QueryUserService interface {
	GetUserInfo(userCode string) (userInfo *dto.UserInfoRes, err error)
}

type CommandUserService interface {
	CreateUser(userInfoReq dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
}

var userService interface {
	QueryUserService
	CommandUserService
}

type defaultUserService struct {
	UserRepo interface {
		repositories.UserInfoRepository
		repositories.UserAttributeRepository
	}
}

func init() {
	userService = initUserService(repositories.New())
}

func initUserService(userInterface interface {
	repositories.UserInfoRepository
	repositories.UserAttributeRepository
}) interface {
	QueryUserService
	CommandUserService
} {
	userService = &defaultUserService{UserRepo: userInterface}
	return userService
}

func getQueryService() QueryUserService {
	return userService
}

func NewUserService() interface {
	QueryUserService
	CommandUserService
} {
	return userService
}

//CreateUser
func (srv *defaultUserService) CreateUser(userInfoReq dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error) {
	if isValid, mess := userInfoReq.ValidateBeforeCreating(); !isValid {
		err = errors.New(mess)
		return
	}

	userModel := models.UserInfo{
		Code:           util.UUID8(),
		PriEmail:       userInfoReq.Email,
		PriMobilePhone: userInfoReq.PhoneNumber,
		FullName:       userInfoReq.FullName,
	}
	if err = srv.UserRepo.SaveUserInfo(&userModel); err != nil {
		return nil, err
	}
	userInfo = &dto.UserInfoRes{
		Code:     userModel.Code,
		FullName: userModel.FullName,
	}
	return
}

//GetUserInfo
func (srv *defaultUserService) GetUserInfo(userCode string) (*dto.UserInfoRes, error) {
	var res *dto.UserInfoRes
	entity, err := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.CodeType,
		UserCode:  userCode,
	})
	if err != nil {
		return res, err
	}
	res = &dto.UserInfoRes{}
	res.Code = entity.Code
	res.FullName = entity.FullName
	return res, nil
}
