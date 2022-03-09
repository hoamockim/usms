package dto

import "time"

type HttpLog struct {
	RequestId   string     `json:"request_id"`
	ServiceName string     `json:"service_name"`
	Detail      ServiceLog `json:"detail"`
}

type ServiceLog struct {
	Request  string    `json:"request"`
	Response string    `json:"response"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type BaseRequest struct {
	Body interface{} `json:"body"`
}

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
