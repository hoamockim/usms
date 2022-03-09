package jwt

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
	"time"
	"usms/pkg/configs"
)

const (
	RS256 = "RS256"
)

type jwtKeys struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

var keys jwtKeys

func init() {
	priBytes, pubBytes := readKeys(configs.GetJwtKeys()) //readKeys("private.key", "public.key")
	priKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priBytes)
	pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if priKey == nil || pubKey == nil {
		panic("Cannot parse jwt key from PEM")
	}
	keys.privateKey = priKey
	keys.publicKey = pubKey
}

func readKeys(priKeyPath, pubKeyPath string) (priBytes, pubBytes []byte) {
	priBytes, _ = ioutil.ReadFile(priKeyPath)
	pubBytes, _ = ioutil.ReadFile(pubKeyPath)

	if len(priBytes) == 0 || len(pubBytes) == 0 {
		panic("jwt is invalid")
	}
	return
}

func GenerateJwtToken(userCode string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &UserClaims{
		UserCode: userCode,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  int64(time.Now().Second()),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(RS256), claims)
	tokenString, err := token.SignedString(keys.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func verifyToken(accessToken string, claims *UserClaims) (httpCode int) {
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return keys.publicKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			httpCode = http.StatusUnauthorized
			return
		}
		httpCode = http.StatusBadRequest
		return
	}
	if !token.Valid {
		httpCode = http.StatusUnauthorized
		return
	}
	return http.StatusOK
}

func RefreshToken(accessToken, refresh string) (newAccessToken string, httpCode int) {
	var claims UserClaims
	httpCode = verifyToken(accessToken, &claims)
	if httpCode != http.StatusOK {
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return newAccessToken, http.StatusBadRequest
	}
	newAccessToken, _ = GenerateJwtToken(claims.UserCode)
	return
}
