// Package api
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:51
package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) RegisterHandlers() {
	s.Handler.GET("/", s.Landing)
	s.Handler.POST("/users/register", s.RegisterUser)
}

func (s *Server) Landing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
