package restHandler

import (
	"fmt"
	"friendly/internal/config"
	"friendly/internal/model"
	"friendly/internal/service/auth_service"
	"friendly/internal/service/friendly_token_service"
	service "friendly/internal/service/user_service"
	"friendly/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AuthMiddleware struct {
	pathConfig        config.RestPathsConfig
	recreateTokenTime time.Duration
	tokenService      friendly_token_service.FriendlyTokenService
	userService       service.UserService
	authServices      []auth_service.AuthService
}

func NewAuthMiddleware(pathConfig config.RestPathsConfig,
	recreateTokenTime time.Duration,
	tokenService friendly_token_service.FriendlyTokenService,
	userService service.UserService,
	authServices []auth_service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		pathConfig:        pathConfig,
		recreateTokenTime: recreateTokenTime,
		tokenService:      tokenService,
		userService:       userService,
		authServices:      authServices,
	}
}

func (s *AuthMiddleware) Middleware(ctx *gin.Context) {
	path := ctx.Request.URL.Path
	method := ctx.Request.Method

	conf, err := s.pathConfig.FindConfig(path, method)
	var userData model.User

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

		uid, exp, err := s.tokenService.ParseToken(token)

		if err != nil {
			for _, authService := range s.authServices {
				uid, err = authService.GetUserUid(ctx, token)
				if err != nil {
					utils.Log(fmt.Sprintf("Invalid token for { %v }", authService), err)
					continue
				}
			}

			if err != nil {
				utils.Log("Invalid token, err: ", err)
				NewErrorResponse(ctx, http.StatusBadRequest, "Invalid token!")
				ctx.Abort()
				return
			}

			friendlyToken, err := s.tokenService.CreateToken(uid)
			if err != nil {
				NewErrorResponse(ctx, http.StatusInternalServerError, "Failed recreate user!")
				logrus.Error("AuthMiddleware. Failed create token with uid:", uid, "err :", err)
				ctx.Abort()
				return
			}

			ctx.Header(TokenHeader, friendlyToken)
		} else {
			if (exp - time.Now().Unix()) < int64(s.recreateTokenTime.Seconds()) {
				newToken, err := s.tokenService.CreateToken(uid)
				if err != nil {
					NewErrorResponse(ctx, http.StatusInternalServerError, "Failed recreate user!")
					logrus.Error("AuthMiddleware. Failed create token with uid:", uid, "err :", err)
					ctx.Abort()
					return
				}

				ctx.Header(TokenHeader, newToken)
			}
		}

		userData, err = s.userService.GetUserByUid(ctx, uid)

		if err != nil {
			ctx.Set(UserUid, uid)
		}

		ctx.Set(UserData, userData)

		ctx.Next()
		return

	}
	ctx.Next()
	return
}
