package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"usms/app/dto"
	"usms/pkg/errors"
)

type baseController struct {
}

func (ctrl *baseController) handleErr(c *gin.Context, errMeta *errors.ErrorMeta, serviceName string) {
	requestId := c.GetHeader("X-Request-ID")
	var appErr *errors.AppError
	var reqBody []byte
	_, _ = c.Request.Body.Read(reqBody)
	_ = c.Request.Body.Close()
	zap.S().Infof("x-request-id :%v, request body: %v", requestId, string(reqBody))
	if errMeta == nil {
		panic("error meta is nil")
	}
	appErr = errors.New(errMeta, serviceName)
	c.JSON(errMeta.HttpCode, appErr)
}

func (ctrl *baseController) success(c *gin.Context, data interface{}) {
	requestId := c.GetHeader("X-Request-ID")
	if data != nil {
		zap.S().Infof("x-request-id :%v, code: %v, message: %v", requestId, http.StatusOK, "success")
		jsonRes, _ := json.Marshal(data)
		zap.S().Infof("x-request-id :%v, response data: %v", requestId, string(jsonRes))
	}
	c.JSON(http.StatusOK, data)
}

func (ctrl *baseController) logDebug(data string) {
	_, _ = fmt.Fprintf(os.Stdout, "[GIN-debug] [Request Info] "+data)
}

func (ctrl *baseController) HealthCheck(ctx *gin.Context) {
	ctrl.success(ctx, dto.BaseResponse{Meta: dto.Meta{Code: strconv.Itoa(http.StatusOK), Message: "Running"}})
}
