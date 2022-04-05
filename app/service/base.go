package service

import (
	"usms/db/repositories"
)

type serviceImpl struct {
	UserRepo interface {
		repositories.UserInfoRepository
		repositories.UserAttributeRepository
	}
}

var userService interface {
	QueryUserService
	CommandUserService
}

var srv *serviceImpl

func init() {
	initService(repositories.New())
	initJwtParse()
}

// initService
func initService(userInterface interface {
	repositories.UserInfoRepository
	repositories.UserAttributeRepository
}) {
	srv = &serviceImpl{UserRepo: userInterface}
}

// GetUserService get user service
func GetUserService() interface {
	QueryUserService
	CommandUserService
} {
	return userService
}

// GetAuthService get authentication service
func GetAuthService() AuthService {
	return srv
}
