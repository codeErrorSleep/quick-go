package bootstrap

import (
	"quick-go/routers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine    *gin.Engine
	apiRouter *routers.Router
}

func (s *Server) Start() {
	s.apiRouter.With(s.engine)
	s.engine.Run()
}

func NewServer(engine *gin.Engine, apiRouter *routers.Router) *Server {
	return &Server{
		engine:    engine,
		apiRouter: apiRouter,
	}
}
