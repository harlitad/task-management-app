package auth

import (
	"errors"
	"time"

	"example.com/task-management-app/config"
	"example.com/task-management-app/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(claims jwt.MapClaims) (string, error) {
	exp := time.Now().Add(60 * time.Minute).Unix()
	claims["exp"] = exp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(config.JWTSECRETKEY))
	if err != nil {
		return "", err
	}
	return jwtToken, err
}

func VerifyToken(token string) (bool, jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSECRETKEY), nil
	})

	if err != nil {
		return false, nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return false, nil, errors.New("token expired")
		}
	}

	if !jwtToken.Valid {
		return false, nil, errors.New("token invalid")
	}

	return true, jwtToken.Claims.(jwt.MapClaims), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(reqPassword string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(reqPassword))

	if err != nil {
		return false, err
	}

	return true, nil

}

func GetUser(user any) model.UserAuthenticated {
	v, ok := user.(model.UserAuthenticated)
	if ok {
		return v
	}
	return model.UserAuthenticated{}
}
