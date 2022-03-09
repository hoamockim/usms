package jwt

import (
	"fmt"
	"testing"
)

func Test_GenerateToken(t *testing.T) {
	accessToken, err := GenerateJwtToken("C00004")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(accessToken)
}

func Test_VerifyToken(t *testing.T) {
	accessToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2NvZGUiOiJDMDAwMDQiLCJleHAiOjE2MzEwOTc1ODMsImlhdCI6NDN9.EtgW_45c26EBYMz2T4CsgH5R2FQZCetmSnHcu2-CMWYdQaaS0QhtT-pW9cpgkbS33Po1msBGO_cR4bbk1-LxF6iE65xE-grzFhU7YwTk5N-ivy8v1qWQSxteGStPGuNaFaBCNuqqwHh3FiPZMkthjk5cHSQrnyakH4LLiM5o1f8vUrV-NTEHfx7DO82Db_Odi2H5IUHq8jykCjOCIUmVJW4sjlZfwDxoYiB7llL7f37rIrGpft6ypyXdturmvQ1sRcXatDOH0i64v-BZ9IZZfrMcpLdFVuEq7xdSlTYbUxtD20JauldJ9Tr05McnwpTaI8LRmk8uq0sHnm0sbzRiWA"
	var claims UserClaims
	code := verifyToken(accessToken, &claims)
	fmt.Println("Httpcode: ", code)
	fmt.Println("User code: ", claims.UserCode)
}
