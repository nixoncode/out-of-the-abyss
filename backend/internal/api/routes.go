// Package api
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:51
package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/out-of-the-abyss/internal/core/users"
	"net/http"
)

func (s *Server) RegisterHandlers() {
	s.Handler.GET("/", s.Landing)

	userHandlers := users.New(s.DB, s.Handler)
	userHandlers.RegisterHandlers()
}

func (s *Server) Landing(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
