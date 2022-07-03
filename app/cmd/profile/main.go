package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"usms/app/controller"
	"usms/pkg/configs"
)

var r *gin.Engine

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {

	appV1 := r.Group("/usms/v1")
	{
		commonService := appV1.Group("sys/")
		{
			commonService.GET("health-check", controller.HealthCheck)
		}

		profile := appV1.Group("/profile")
		{
			profile.GET("/:code", controller.GetProfileDetail)
			profile.POST("sign-up", controller.SignUp)
		}
	}

	if err := r.Run(configs.AppURL()); err != nil {
		os.Exit(1)
	}

	/*lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcServer.Serve(lis)*/
}

func init() {
	r = gin.New()
}
