package server

import (
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"example.com/gotraining/go-hexagonal_http_api-course/internal/platform/server/handler/health"
	"example.com/gotraining/go-hexagonal_http_api-course/kit/bus"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	bus bus.Bus
}

func New(host string, port uint, bus bus.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		bus: bus,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.bus))
	s.engine.GET("/courses", courses.GetHandler(s.bus))
}