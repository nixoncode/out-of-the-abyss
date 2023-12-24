// Package api
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 20:03
package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
	"net/http"
)

func (s *Server) RegisterUser(ctx echo.Context) error {
	fmt.Println(ctx.FormValue("email"))
	user := sqlc.CreateUserParams{}
	err := ctx.Bind(&user)
	if err != nil {
		return err
	}

	createUser, err := s.DB.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	id, err := createUser.LastInsertId()

	msg := fmt.Sprintf("user registered with ID: %d", id)

	return ctx.String(http.StatusOK, msg)
}
