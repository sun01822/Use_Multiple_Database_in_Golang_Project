package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"myapp/config"
	"myapp/conn"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use: "myapp",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	config.LoadConfig()
	conn.ConnectBQ()
	conn.ConnectPostgres()
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
