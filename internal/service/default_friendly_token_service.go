package service

import (
	"errors"
	"friendly/internal/utils"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type DefaultFriendlyTokenService struct {
	secretKey  string
	activeTime time.Duration
}

func NewDefaultFriendlyTokenService(secretKey string, activeTime time.Duration) *DefaultFriendlyTokenService {
	return &DefaultFriendlyTokenService{secretKey: secretKey, activeTime: activeTime}
}

func (d *DefaultFriendlyTokenService) ParseToken(tokenString string) (string, int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(d.secretKey), nil
	})

	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["id"].(string)
		unix := int64(claims["exp"].(float64))
		return userId, unix, nil
	}
	utils.Log("Incorrect token", nil)

	return "", 0, errors.New("invalid token")
}
func (d *DefaultFriendlyTokenService) CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(d.activeTime).Unix(),
	})

	signedToken, err := token.SignedString([]byte(d.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
func (d *DefaultFriendlyTokenService) GetFriendsToken(uid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
	})

	signedToken, err := token.SignedString([]byte(d.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, err
}
func (d *DefaultFriendlyTokenService) ParseFriendsToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(d.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["id"].(string)
		return userId, nil
	}
	utils.Log("Incorrect token", nil)

	return "", errors.New("invalid token")
}
