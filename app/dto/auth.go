package dto

import (
	"regexp"
	"strings"
)

type SignInType int8

const (
	UserPass SignInType = iota
	Email
	Facebook
	PhoneNumber
	Refresh
)

const (
	emailPattern = "(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)"
	phonePattern = "^[0-9]{10,11}$"
)

type JwtTokenRequestBody struct {
	SignInType   SignInType `json:"sign_in_type"`
	UserName     string     `json:"user_name"`
	PassWord     string     `json:"pass_word"`
	Email        string     `json:"email"`
	Facebook     string     `json:"facebook"`
	Phone        string     `json:"phone"`
	RefreshToken string     `json:"refresh_token"`
}

type JwtTokenResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateJwtTokenValidate interface {
	ValidateCreateJwtToken() bool
}

type RefreshJwtTokenValidate interface {
	ValidateRefreshJwtToken() bool
}

func (bodyReq *JwtTokenRequestBody) ValidateCreateJwtToken() bool {
	switch bodyReq.SignInType {
	case UserPass:
		return strings.TrimSpace(bodyReq.UserName) == "" && strings.TrimSpace(bodyReq.PassWord) == ""
	case Email:
		m, err := regexp.MatchString(emailPattern, bodyReq.Email)
		return m || err != nil
	case PhoneNumber:
		m, err := regexp.MatchString(phonePattern, bodyReq.Phone)
		return m || err != nil
	case Refresh:
		return strings.TrimSpace(bodyReq.RefreshToken) == ""
	default:
		return false
	}
}

func (bodyReq *JwtTokenRequestBody) ValidateRefreshJwtToken() bool {
	return true
}
