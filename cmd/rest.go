package main

import (
	"hexagonal-template/protocol"
	"log"

	"github.com/spf13/cobra"
)

/*
	|--------------------------------------------------------------------------
	| Command to serve REST protocol
	|--------------------------------------------------------------------------
	|
	| Here is the command that makes your application serve the REST protocol
	| for client.
	|
*/

var serveRestCommand = &cobra.Command{
	Use:   "serve-rest",
	Short: "start a http server",
	RunE:  serveRest,
}

func init() {
	serveRestCommand.PersistentFlags().StringVarP(&env, "env", "e", "production", "environment name")
	serveRestCommand.PersistentFlags().StringVarP(&cfgPath, "config", "c", "config.yaml", "config file path")
	log.Println("you are running on environment: ", env)
}

func serveRest(cmd *cobra.Command, args []string) error {
	return protocol.ServeRest(cfgPath, env)
}
