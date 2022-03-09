package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"usms/app"
	"usms/app/controller"
)

var r *gin.Engine

func main() {
	appV1 := r.Group("/auth/v1")
	{
		commonService := appV1.Group("/")
		{
			commonService.GET("health-check", controller.GetAuthController().HealthCheck)
		}

		jwtTokenService := appV1.Group("/token/")
		{
			jwtTokenService.POST("request", controller.GetAuthController().CreateJwtToken)
			jwtTokenService.POST("refresh", controller.GetAuthController().RefreshJwtToken)
		}
	}

	if err := r.Run("localhost:8082"); err != nil {
		os.Exit(1)
	}
}

func init() {
	r = gin.New()
	app.InitLogger()
}
