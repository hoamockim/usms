package jwt

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	UserCode string `json:"user_code"`
	Role     string `json:"role"`
	*jwt.StandardClaims
}
