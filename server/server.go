package server

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Echo *echo.Echo
}


func accessControl(server *Server) {
	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}

func New() *Server {
	var e *echo.Echo

	e = echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
    })
    	
	s := &Server{
		Echo: e,
	}

	accessControl(s)

	return s
}