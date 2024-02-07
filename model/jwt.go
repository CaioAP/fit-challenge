package model

import "github.com/go-chi/jwtauth"

type JWT struct{}

func CreateToken(tokenAuth *jwtauth.JWTAuth, identifier int) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"identifier": identifier})
	return tokenString
}
