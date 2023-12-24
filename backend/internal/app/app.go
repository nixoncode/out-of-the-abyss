// Package app
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:17
package app

import (
	"database/sql"
	"github.com/nixoncode/out-of-the-abyss/cmd"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
	"github.com/spf13/cobra"
)

type App struct {
	RootCmd *cobra.Command
	DB      *sqlc.Queries
}

func (app *App) Start() error {
	app.RootCmd.AddCommand(cmd.NewServeCommand(app.DB))

	return app.Execute()
}

func (app *App) Execute() error {
	return app.RootCmd.Execute()
}

func New(db *sql.DB) *App {
	return &App{
		RootCmd: &cobra.Command{
			Use:   "out-of-the-abyss",
			Short: "Out of the abyss CLI",
		},
		DB: sqlc.New(db),
	}
}
