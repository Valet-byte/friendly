package restHandler

import (
	"errors"
	"fmt"
	"friendly/docs"
	"friendly/internal/model"
	"friendly/internal/service/friends_service"
	"friendly/internal/service/user_description_service"
	"friendly/internal/service/user_service"
	"friendly/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	TokenHeader         = "Token"
	UserData            = "UserData"
	UserUid             = "UserUid"
)

type RestHandler struct {
	userService            user_service.UserService
	userDescriptionService user_description_service.UserDescriptionService
	userFriendsService     friends_service.FriendsService
	rateLimiter            *ReteLimiter
	authMiddleware         *AuthMiddleware
}

func (h *RestHandler) InitHandlers() *gin.Engine {

	docs.SwaggerInfo.BasePath = "/api/V1"

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("/api/V1", h.rateLimiter.RateLimiterMiddleware, h.authMiddleware.Middleware)
	{
		users := apiV1.Group("/users") // Работа с пользователями
		{
			users.GET("/", h.getUserData)                       // получаем информацию о себе
			users.POST("/", h.updateUserData)                   // обновляем свои данные / регистрация
			users.DELETE("/", h.deleteUser)                     // удаляем аккаунт
			users.GET("/:id", h.getUserData)                    // получаем информацию о пользователе по id
			users.POST("/description", h.addUserDescription)    //добавление / изменение описания пользователя
			users.GET("/:id/description", h.getUserDescription) //получаем описание пользователя
			users.GET("/description", h.getUserDescription)     //получаем описание пользователя

			friends := users.Group("/friends") // работа с друзьями
			{
				friends.GET("/", h.getFriends)                 // список наших друзей
				friends.GET("/:id", h.getFriends)              // список друзей пользователя с id
				friends.GET("/token", h.generateFriendToken)   // создать токен друзей
				friends.PUT("/:id", h.addFriendsByParam)       // добавить друга по id
				friends.PUT("/", h.addFriendsByToken)          // добавить пользователя по токену
				friends.DELETE("/:id", h.deleteFriendsByParam) // удалить из друзей по id
			}

		}

		events := apiV1.Group("/events") // события
		{
			events.GET("/", do)
			events.GET("/search", do)
			events.POST("/")
			events.PUT("/:id")
		}

	}

	return router
}

// @Summary Получить приветствие
// @Description Возвращает строку "Привет, мир!"
// @Tags hello
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Привет, мир!"
// @Router /hello [get]
func do(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, "Fot this function isn't realization")
}

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	utils.Log(fmt.Sprintf("Message : { %s } ", message), nil)

	c.AbortWithStatusJSON(http.StatusOK, DefaultResponse{
		Code:    statusCode,
		Message: message,
	})
}

func GetUserToken(ctx *gin.Context) (string, error) {
	header := ctx.GetHeader(AuthorizationHeader)

	if header == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return "", errors.New("invalid auth header")
	}

	return headerParts[1], nil
}

func GetUserData(ctx *gin.Context) (model.User, error) {
	data, exist := ctx.Get(UserData)
	if !exist {
		logrus.Error("Not found user data! user_handler.deleteUser")
		NewErrorResponse(ctx, http.StatusInternalServerError, "not found user data")
		return model.User{}, errors.New("failed get user data from request")
	}
	return data.(model.User), nil
}
