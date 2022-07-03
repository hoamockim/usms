package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"usms/app"
	"usms/app/controller"
	"usms/pkg/configs"
)

var r *gin.Engine

func main() {
	appV1 := r.Group("/usms/v1")
	{
		commonService := appV1.Group("/")
		{
			commonService.GET("health-check", controller.HealthCheck)
		}

		auth := appV1.Group("/auth/")
		{
			auth.POST("sign-in", controller.SignIn)

			auth.POST("refresh", controller.RefreshJwtToken)
		}
	}

	if err := r.Run(configs.AppURL()); err != nil {
		os.Exit(1)
	}
}

func init() {
	r = gin.New()
	app.InitLogger()
	app.InitAuth()
}
