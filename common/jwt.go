package common

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

}

//GetTokenAuth Returns the JWTAuth object
func GetTokenAuth() *jwtauth.JWTAuth {
	tokenAuth = jwtauth.New("HS256", []byte("private_key"), "publickey")
	return tokenAuth
}
