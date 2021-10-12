package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version defines version
const (
	Version = "unversioned"
)

/*
	|--------------------------------------------------------------------------
	| Application's Command
	|--------------------------------------------------------------------------
	|
	| Here is which command you want to provide for your application
	| to use in your application.
	|
*/

// rootCmd is the root of all sub commands in the binary
// it doesn't have a Run method as it executes other sub commands
var rootCommand = &cobra.Command{
	Use:     "user task",
	Short:   "task manages user task",
	Version: Version,
}
var cfgPath string
var env string

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCommand.AddCommand(serveRestCommand)
}
