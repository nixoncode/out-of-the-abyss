// Package api
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:36
package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Server struct {
	port    int
	Handler *echo.Echo
}

func New() *Server {
	e := echo.New()

	return &Server{
		Handler: e,
	}
}

func (s *Server) Start(port int) error {
	s.RegisterHandlers()
	return s.Handler.Start(fmt.Sprintf(":%d", port))
}
