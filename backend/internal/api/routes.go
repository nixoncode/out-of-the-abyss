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
	s.Handler.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
