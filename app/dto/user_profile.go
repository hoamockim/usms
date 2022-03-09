package dto

import (
	"usms/pkg/util"
)

type UserInfoReq struct {
	FullName         string `json:"full_name"`
	Address          string `json:"address,omitempty"`
	PhoneNumber      string `json:"phone_number"`
	ExtraPhoneNumber string `json:"extra_phone_number,omitempty"`
	Email            string `json:"email"`
	Gender           int8   `json:"gender,omitempty"`
	Birthday         string `json:"birthday,omitempty"`
}

type UserInfoRes struct {
	Code     string `json:"code"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

func (req *UserInfoReq) ValidateBeforeCreating() (matched bool, message string) {
	if matched, message = util.IsEmail(req.Email); !matched {
		return
	}
	return
}
