package model

import (
	"context"

	"github.com/go-chi/jwtauth"
)

type JWT struct{}

type JWTBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateToken(tokenAuth *jwtauth.JWTAuth, id int, name string) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"id": id, "name": name})
	return tokenString
}

func DecodeToken(tokenAuth *jwtauth.JWTAuth, ctx context.Context) (JWTBody, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return JWTBody{}, err
	}
	return JWTBody{
		ID:   int(claims["id"].(float64)),
		Name: claims["name"].(string),
	}, nil
}
