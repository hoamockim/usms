package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HealthCheck(ctx *gin.Context) {
	success(ctx, BaseResponse{Meta: Meta{Code: strconv.Itoa(http.StatusOK), Message: "Running"}})
}
