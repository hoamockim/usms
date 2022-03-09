package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"usms/app/controller"
)

var r *gin.Engine

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {

	appV1 := r.Group("/usms/v1")
	{
		commonService := appV1.Group("/")
		{
			commonService.GET("health-check", controller.UserProfile().HealthCheck)
		}

		gatewayService := appV1.Group("/user-info")
		{
			gatewayService.GET("/:code", controller.UserProfile().GetProfileDetail)
			gatewayService.POST("", controller.UserProfile().SaveUserInfo)
		}
	}

	if err := r.Run("localhost:8081"); err != nil {
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
