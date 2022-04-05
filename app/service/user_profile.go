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
	VerifyByEmail(email, password string) (userInfo *dto.UserInfoRes, err error)
}

type CommandUserService interface {
	CreateUser(userInfoReq dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
}

// CreateUser create new user for system
func (srv *serviceImpl) CreateUser(userInfoReq dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error) {
	if isValid, mess := userInfoReq.ValidateBeforeCreating(); !isValid {
		err = errors.New(mess)
		return
	}

	entity := models.UserInfo{
		Code:           util.UUID8(),
		PriEmail:       userInfoReq.Email,
		PriMobilePhone: userInfoReq.PhoneNumber,
		FullName:       userInfoReq.FullName,
	}
	if err = srv.UserRepo.SaveUserInfo(&entity); err != nil {
		return nil, err
	}
	userInfo = &dto.UserInfoRes{
		Code:     entity.Code,
		FullName: entity.FullName,
	}
	return
}

//GetUserInfo get user's information
func (srv *serviceImpl) GetUserInfo(userCode string) (*dto.UserInfoRes, error) {
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

func (srv *serviceImpl) VerifyByEmail(email, password string) (userInfo *dto.UserInfoRes, err error) {
	var res *dto.UserInfoRes
	entity, err := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.EmailType,
		Email:     email,
		PassWord:  password,
	})
	if err != nil {
		return res, err
	}
	res = &dto.UserInfoRes{}
	res.Code = entity.Code
	res.FullName = entity.FullName
	return res, nil
}
