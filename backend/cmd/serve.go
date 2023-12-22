// Package cmd
// Created by GoLand.
// User: nixon
// Date: 22/12/2023
// Time: 10:02
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func NewServeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "start the web server on localhost:8080",
		Run: func(cmd *cobra.Command, args []string) {
			// todo start the server
			log.Println("server should run here")
		},
	}

	return command
}
