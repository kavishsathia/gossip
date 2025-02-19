package helpers

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func VerifyJWT(tokenString string) (*User, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT secret key not found in environment")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("userId claim not found or invalid")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("username claim not found or invalid")
	}

	return &User{
		Username: username,
		UserID:   int(userID),
	}, nil
}

func GenerateJWT(userId int, username string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", errors.New("JWT private key not found in environment")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		print(err.Error())
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}

func GetUserInfo(c *gin.Context) (interface{}, *User, error) {
	user, ok := c.Get("user")
	if !ok {
		return nil, nil, fmt.Errorf("failed to get user")
	}

	userInfo, ok := user.(*User)
	if !ok {
		return user, nil, fmt.Errorf("failed to cast userInfo")
	}

	return user, userInfo, nil
}
