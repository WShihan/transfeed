package jwt

import (
	"errors"
	"transfeed/internal/app/model"
	"transfeed/internal/app/store"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(data interface{}) (*model.User, error) {
	token, ok := data.(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}
	user := model.User{}
	id := claims["id"].(float64)
	err := store.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	if user.Username == "" {
		return nil, errors.New("user not exist")
	}
	if *user.Token != token.Raw {
		return nil, errors.New("token not match")
	}
	return &user, nil
}
