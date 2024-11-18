package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("SldUQW5hbmRhQVBNS0VZ")

type JWTToken struct{
	Username string
	jwt.RegisteredClaims
}

func Generate(username string)(string,error){
	exp := time.Now().Add(2 * time.Hour)

	jwttoken := &JWTToken{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512,jwttoken)
	
	
	return token.SignedString(key)
}

func ValidateJWT(tokenReq string)(*JWTToken,error){

	jwttoken := &JWTToken{}

	token, err:= jwt.ParseWithClaims(tokenReq,jwttoken,func(token *jwt.Token) (interface{}, error){
		return key, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return jwttoken,nil
}