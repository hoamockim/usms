module usms

go 1.15

require (
	github.com/Shopify/sarama v1.30.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-kit/kit v0.10.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/ugorji/go v1.1.13 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	google.golang.org/api v0.3.1
	google.golang.org/genproto v0.0.0-20211222154725-9823f7ba7562 // indirect
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.0
