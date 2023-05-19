package restHandler

import (
	"errors"
	"friendly/docs"
	service2 "friendly/internal/service"
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
)

type RestHandler struct {
	fb *service2.FirebaseAppService
	vk *service2.VkService
	rl *ReteLimiter
	am *AuthMiddleware
	rd *service2.RedisService
}

func NewRestHandler(fb *service2.FirebaseAppService, rl *ReteLimiter, vk *service2.VkService, am *AuthMiddleware, rd *service2.RedisService) *RestHandler {
	return &RestHandler{fb: fb, rl: rl, vk: vk, am: am, rd: rd}
}

func (h *RestHandler) InitHandlers() *gin.Engine {

	router := gin.New()
	docs.Init()

	apiV1 := router.Group("/api/V1", h.rl.RateLimiterMiddleware, h.am.Middleware)
	{
		apiV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		users := apiV1.Group("/users") // Работа с пользователями
		{
			users.GET("/", h.getUserData)                       // получаем информацию о себе
			users.POST("/", h.registrationUser)                 // обновляем свои данные / регистрация
			users.PUT("/", h.updateUserData)                    // обновляем свои данные / регистрация
			users.DELETE("/", h.deleteUser)                     // удаляем аккаунт
			users.GET("/:id", h.getUserData)                    // получаем информацию о пользователе по id
			users.GET("/exists", h.auth)                        // получаем информацию о существовании юзера
			users.POST("/description", h.addUserDescription)    //добавление / изменение описания пользователя
			users.GET("/:id/description", h.getUserDescription) //получаем описание пользователя
			users.GET("/description", h.getUserDescription)     //получаем описание пользователя

			friends := users.Group("/friends") // работа с друзьями
			{
				friends.GET("/", do)         // список наших друзей
				friends.GET("/:id", do)      // список друзей пользователя с id
				friends.GET("/size", do)     // колличесво наших друзей
				friends.GET("/:id/size", do) // колличесво друзей пользователя с id
				friends.PUT("/:id", do)      // добавить друга по id
				friends.DELETE("/:id", do)   // удалить из друзей по id

			}

			points := users.Group("/points", do) // работа с координатами
			{
				points.PUT("/", do)        // отправляем наши координаты
				points.GET("/", do)        // получаем координаты друзей в радиусе вокруг
				points.GET("/connect", do) //создаем подключение через WebSocets, и шлем данные и получаем их
			}

			chat := users.Group("/chat", do) // работа с чатом
			{
				chat.POST("/send", do) // отправляем сообщение
			}
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
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf("Message : { %s } ", message)

	c.AbortWithStatusJSON(statusCode, DefaultResponse{message})
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
