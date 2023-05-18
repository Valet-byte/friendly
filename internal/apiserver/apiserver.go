package apiserver

import (
	"context"
	"friendly/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiServer struct {
	server *http.Server
}

func NewApiServer(conf *config.FriendlyConfig, handler *gin.Engine) *ApiServer {
	return &ApiServer{&http.Server{
		Handler:      handler,
		Addr:         conf.Server.Host + ":" + conf.Server.Port,
		ReadTimeout:  conf.Server.Timeout.Read,
		WriteTimeout: conf.Server.Timeout.Write,
	}}
}

func (s *ApiServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *ApiServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
