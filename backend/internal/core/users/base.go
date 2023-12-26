// Package users
// Created by GoLand.
// User: nixon
// Date: 24/12/2023
// Time: 09:47
package users

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
)

type Users struct {
	DB *sqlc.Queries
	H  *echo.Echo
}

func New(db *sqlc.Queries, h *echo.Echo) *Users {
	return &Users{DB: db, H: h}
}

func (u *Users) RegisterHandlers() {
	h := u.H.Group("/users")
	h.POST("/register", u.register)
}
