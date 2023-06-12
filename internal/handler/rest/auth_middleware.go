package restHandler

import (
	"friendly/internal/config"
	"friendly/internal/model"
	service2 "friendly/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	friendlyToken = "friendly_token"
	firebaseToken = "firebase_token"
)

type AuthMiddleware struct {
	pathConfig        *config.RestPathsConfig
	recreateTokenTime time.Duration
	frTokenService    *service2.FriendlyTokenService
	fbApp             *service2.FirebaseAppService
}

func NewAuthMiddleware(pathConfig *config.RestPathsConfig, recreateTokenTime time.Duration, frTokenService *service2.FriendlyTokenService, fbApp *service2.FirebaseAppService) *AuthMiddleware {
	return &AuthMiddleware{pathConfig: pathConfig, recreateTokenTime: recreateTokenTime, frTokenService: frTokenService, fbApp: fbApp}
}

func (s *AuthMiddleware) Middleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method

	conf, err := s.pathConfig.FindConfig(path, method)
	var userData *model.User

	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Incorrect path or method!")
		ctx.Abort()
		return
	}

	if !conf.IsOpen {
		token, err := GetUserToken(ctx)

		if err != nil {
			ctx.Abort()
			return
		}

		if conf.AuthMethod == friendlyToken {
			uid, exp, err := s.frTokenService.ValidateToken(token)

			if err != nil {
				logrus.Error("Invalid token, err: ", err)
				NewErrorResponse(ctx, http.StatusBadRequest, "Invalid token!")
				ctx.Abort()
				return
			}

			userData, err = s.frTokenService.GetUserData(uid, ctx)

			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
				logrus.Error("AuthMiddleware. Not found user with uid:", uid, "err :", err)
				ctx.Abort()
				return
			}

			ctx.Set(UserData, userData)

			if (exp - time.Now().Unix()) < int64(s.recreateTokenTime.Seconds()) {
				newToken, err := s.frTokenService.CreateToken(uid)
				if err != nil {
					NewErrorResponse(ctx, http.StatusInternalServerError, "Failed recreate user!")
					logrus.Error("AuthMiddleware. Failed create token with uid:", uid, "err :", err)
					ctx.Abort()
					return
				}

				ctx.Header(TokenHeader, newToken)
			}

			ctx.Next()
			return
		}

		if conf.AuthMethod == firebaseToken {
			uid, err := s.fbApp.GetUserUid(ctx, token)

			if err != nil {
				NewErrorResponse(ctx, http.StatusBadRequest, "Invalid token!")
				ctx.Abort()
				return
			}

			userData, err = s.frTokenService.GetUserData(uid, ctx)

			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
				logrus.Error("AuthMiddleware. Not found user with uid:", uid, "err :", err)
				ctx.Abort()
				return
			}

			ctx.Set(UserData, userData)

			newToken, err := s.frTokenService.CreateToken(uid)
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "Failed recreate user!")
				logrus.Error("AuthMiddleware. Failed create token with uid:", uid, "err :", err)
				ctx.Abort()
				return
			}

			ctx.Header(TokenHeader, newToken)

			ctx.Next()
			return
		}
	}
	ctx.Next()
	return
}
