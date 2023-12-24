// Package api
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:36
package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
)

type Server struct {
	Handler *echo.Echo
	DB      *sqlc.Queries
}

func New(db *sqlc.Queries) *Server {
	e := echo.New()

	return &Server{
		Handler: e,
		DB:      db,
	}
}

func (s *Server) Start(port int) error {
	s.RegisterHandlers()
	return s.Handler.Start(fmt.Sprintf(":%d", port))
}
