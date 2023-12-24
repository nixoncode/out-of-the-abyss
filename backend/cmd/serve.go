// Package cmd
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:02
package cmd

import (
	"github.com/nixoncode/out-of-the-abyss/internal/api"
	"github.com/nixoncode/out-of-the-abyss/internal/db/sqlc"
	"github.com/spf13/cobra"
)

func NewServeCommand(db *sqlc.Queries) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "start the web api on 0.0.0.0:8080",
		Run: func(cmd *cobra.Command, args []string) {
			server := api.New(db)

			err := server.Start(8080)
			if err != nil {
				panic(err)
			}
		},
	}

	return command
}
