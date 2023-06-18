package app

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"friendly/internal/apiserver"
	"friendly/internal/cache"
	"friendly/internal/config"
	restHandler "friendly/internal/handler/rest"
	"friendly/internal/repository/user_description_repository"
	"friendly/internal/repository/user_repository"
	"friendly/internal/service"
	"friendly/internal/service/auth_service"
	"friendly/internal/service/user_description_service"
	"friendly/internal/service/user_service"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func Start(configPath, pathToFirebaseSecret, pathsConfig string) {

	conf, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
	level, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
	logrus.SetLevel(level)

	ptConf, err := config.NewRestPathsConfig(pathsConfig)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}

	app, err := getFirebaseApp(pathToFirebaseSecret)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server.Cache.Host,
		Password: conf.Server.Cache.Password,
	})
	redisService := cache.NewRedisService(client, conf.Server.Cache.Duration)

	userRepo := user_repository.NewFirebaseUserRepository(app)
	userDescriptionRepo := user_description_repository.NewFirebaseDescriptionRepository(app)
	//friendlyRepo := user_friends_repository.NewFirebaseFriendsRepository(app)

	userService := user_service.NewDefaultUserService(userRepo, redisService)
	userDescriptionService := user_description_service.NewDefaultUserDescriptionService(userDescriptionRepo, redisService)
	authService := auth_service.NewFirebaseAuthService(app)
	friendlyTokenService := service.NewDefaultFriendlyTokenService(conf.Server.Jwt.SecretKey, conf.Server.Jwt.ValidTime)

	rateLimiter := restHandler.NewRateLimiter(10, redisService)
	authMiddleware := restHandler.NewAuthMiddleware(*ptConf, conf.Server.Jwt.RecreateTime, friendlyTokenService, userService, []auth_service.AuthService{authService})
	//vkService := service2.NewVkService("https://api.vk.com/method/users.get", "5.131")

	handler := restHandler.NewRestHandler(userService, userDescriptionService, rateLimiter, authMiddleware)
	handles := handler.InitHandlers()

	server := apiserver.NewApiServer(conf, handles)
	err = server.Run()
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
}

func getFirebaseApp(pathToSecretFirebase string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(pathToSecretFirebase)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	return app, err
}
