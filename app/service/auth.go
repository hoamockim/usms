package service

import (
	"usms/app/dto"
)

type AuthService interface {
	RequestToken(reqBody *dto.JwtTokenRequestBody) (dto.JwtTokenResponseData, error)
	RefreshToken(reqBody dto.JwtTokenRequestBody) (dto.JwtTokenResponseData, error)
	DeleteToken(accessToken string) (bool, error)
}

type defaultAuthService struct {
	userService interface {
		QueryUserService
	}
}

var defaultAuth AuthService

func NewAuthService() AuthService {
	return defaultAuth
}

func init() {
	defaultAuth = &defaultAuthService{getQueryService()}
}

func (srv *defaultAuthService) RequestToken(reqBody *dto.JwtTokenRequestBody) (resData dto.JwtTokenResponseData, err error) {
	return
}

func (srv *defaultAuthService) RefreshToken(reqBody dto.JwtTokenRequestBody) (resData dto.JwtTokenResponseData, err error) {
	return
}

func (srv *defaultAuthService) DeleteToken(accessToken string) (bool, error) {
	return true, nil
}
