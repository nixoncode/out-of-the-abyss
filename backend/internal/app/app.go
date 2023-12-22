// Package app
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:17
package app

import (
	"github.com/nixoncode/out-of-the-abyss/cmd"
	"github.com/spf13/cobra"
)

type App struct {
	RootCmd *cobra.Command
}

func (app *App) Start() error {
	app.RootCmd.AddCommand(cmd.NewServeCommand())

	return app.Execute()
}

func (app *App) Execute() error {
	return app.RootCmd.Execute()
}

func New() *App {
	return &App{
		RootCmd: &cobra.Command{
			Use:   "out-of-the-abyss",
			Short: "Out of the abyss CLI",
		},
	}
}
