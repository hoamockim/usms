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
	Register(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error)
}

// Register create new user for system
func (srv *serviceImpl) Register(req dto.UserInfoReq) (userInfo *dto.UserInfoRes, err error) {
	if isValid := req.ValidateBeforeCreating(); !isValid {
		err = errors.New("email or password is not correct")
		return
	}
	//1: check account exist with the current information
	//if existed: return error with message user is existed
	var entity *models.UserInfo
	entity = srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.EmailType,
		UserCode:  req.Email,
	})

	entity = &models.UserInfo{
		Code:           util.UUID8(),
		PriEmail:       req.Email,
		PriMobilePhone: req.PhoneNumber,
		FullName:       req.FullName,
	}

	// 2: create new account if it's not existed
	// create verify link and send via email
	// tracking register
	if err = srv.UserRepo.SaveUserInfo(entity); err != nil {
		return nil, err
	}
	userInfo = &dto.UserInfoRes{
		Code:     entity.Code,
		FullName: entity.FullName,
	}
	return
}

// GetUserInfo get user's information
func (srv *serviceImpl) GetUserInfo(userCode string) (*dto.UserInfoRes, error) {
	var res *dto.UserInfoRes
	entity := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
		InputType: repositories.CodeType,
		UserCode:  userCode,
	})

	res = &dto.UserInfoRes{}
	res.Code = entity.Code
	res.FullName = entity.FullName
	return res, nil
}

// VerifyByEmail verify an account by email
func (srv *serviceImpl) VerifyByEmail(email, password string) (userInfo *dto.UserInfoRes, err error) {
	var res *dto.UserInfoRes
	entity := srv.UserRepo.GetUserInfo(&repositories.UserFilter{
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
