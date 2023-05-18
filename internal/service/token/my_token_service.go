package service

import (
	"errors"
	"friendly/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type FriendlyTokenService struct {
	secretKey    string
	redisService *RedisService
	fbApp        *FirebaseAppService
	activeTime   time.Duration
}

func NewFriendlyTokenService(secretKey string, redisService *RedisService, fbApp *FirebaseAppService, activeTime time.Duration) *FriendlyTokenService {
	return &FriendlyTokenService{secretKey: secretKey, redisService: redisService, fbApp: fbApp, activeTime: activeTime}
}

func (s *FriendlyTokenService) CreateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(s.activeTime).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *FriendlyTokenService) ValidateToken(tokenString string) (string, int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["id"].(string)
		unix := int64(claims["exp"].(float64))
		return userId, unix, nil
	}

	return "", 0, errors.New("invalid token")
}

func (s *FriendlyTokenService) GetUserData(userId string, ctx *gin.Context) (*model.User, error) {
	userData, err := s.redisService.GetUser(userId)

	if err != nil {

		userData, err = s.fbApp.GetUserData(ctx, userId)

		if userData == nil {
			return nil, errors.New("not found user data")
		}

		_, err := s.redisService.PutUser(userData.Uid, *userData)
		if err != nil {
			logrus.Error("Error save user to redis! FriendlyTokenService(62), error :", err)
		}

		if err != nil {
			return nil, err
		}

	}

	return userData, nil
}
