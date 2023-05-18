package app

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"friendly/internal/apiserver"
	"friendly/internal/config"
	restHandler "friendly/internal/handler/rest"
	service "friendly/internal/service/token"
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
	app, err := getFirebaseApp(pathToFirebaseSecret)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
	fbApp := service.NewFirebaseApp(app)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server.Cache.Host,
		Password: conf.Server.Cache.Password,
	})
	ptConf, err := config.NewRestPathsConfig(pathsConfig)
	if err != nil {
		logrus.Fatalf("Fatal error! error : {%s}", err.Error())
	}
	redisService := service.NewRedisService(client, conf.Server.Cache.Duration)
	friendlyTS := service.NewFriendlyTokenService(conf.Server.Jwt.SecretKey, redisService, fbApp, conf.Server.Jwt.ValidTime)
	rl := restHandler.NewRateLimiter(10, redisService)
	am := restHandler.NewAuthMiddleware(ptConf, conf.Server.Jwt.RecreateTime, friendlyTS, fbApp)
	vkService := service.NewVkService("https://api.vk.com/method/users.get", "5.131")
	handler := restHandler.NewRestHandler(fbApp, rl, vkService, am, redisService)
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
