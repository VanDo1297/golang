package server

import (
	"demo/internal/resources"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Resources resources.IResources
}

func NewDefaultServer() (*Server, error) {
	defaultResources, err := resources.NewDefaultResources()
	if err != nil {
		return nil, err
	}
	serv := &Server{
		Resources: defaultResources,
	}

	serv.addDemoRoutes()

	return serv, nil
}

func (server Server) addDemoRoutes() {
	r := gin.New()
	r.GET("/demo", func(ctx *gin.Context) {
		server.GetDemo(ctx)
	})
}
