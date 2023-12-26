// Package users
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 20:03
package users

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
	"net/http"
)

func (u *Users) register(ctx echo.Context) error {
	fmt.Println(ctx.FormValue("email"))
	user := sqlc.CreateUserParams{}
	err := ctx.Bind(&user)
	if err != nil {
		return err
	}

	createUser, err := u.DB.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	id, err := createUser.LastInsertId()

	msg := fmt.Sprintf("users registered with ID: %d", id)

	return ctx.String(http.StatusOK, msg)
}
