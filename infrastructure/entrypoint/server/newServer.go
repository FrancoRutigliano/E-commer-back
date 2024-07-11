package server

import (
	"e-commer/infrastructure/entrypoint/server/routes"

	"github.com/labstack/echo/v4"
)

type Server struct{}

// Constructor de servidor
func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := echo.New()

	routes.Init(app)

	return app.Start(":8080")
}
