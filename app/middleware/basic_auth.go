package middleware

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"usms/pkg/util"
)

const (
	HeaderAuthorizationKey = "Authorization"
)

type Accounts map[string]string

func BasicAuth(accounts Accounts) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			basicAuthData = c.GetHeader(HeaderAuthorizationKey)
			splits        = strings.Split(basicAuthData, " ")
		)

		if len(splits) != 2 && splits[0] != "Basic" {
			zap.S().Errorf("Authorization header is invalid format")
			//utils.handleErr(c, errors.New(errors.ErrUnauthorized))
			c.Abort()
			return
		}

		decodedString, err := base64.StdEncoding.DecodeString(splits[1])
		if err != nil || strings.TrimSpace(string(decodedString)) == "" {
			zap.S().Errorw(fmt.Sprintf("Decoding header fail"))

			c.Abort()
			return
		}
		userAndPwd := strings.Split(string(decodedString), ":")

		if len(userAndPwd) != 2 {
			zap.S().Errorf("Cannot parse username and password")
			c.Abort()
			return
		}
		isAuthorized := accounts.Authorized(userAndPwd[0], userAndPwd[1])
		if !isAuthorized {
			zap.S().Errorf("Unauthorized")
			//utils.handleErr(c, errors.New(errors.ErrUnauthorized))
			c.Abort()
			return
		}
	}
}

func (accounts Accounts) Authorized(username, password string) bool {
	pwd, ok := accounts[username]
	if !ok {
		return false
	}
	return strings.Compare(util.Sha256(password), pwd) == 0
}
