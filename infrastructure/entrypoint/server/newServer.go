package server

import "github.com/labstack/echo/v4"

type Server struct{}

// Constructor de servidor
func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return app.Start(":8080")
}
