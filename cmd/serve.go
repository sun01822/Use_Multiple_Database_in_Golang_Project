package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"myapp/server"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// Your code here
	var echo_ = echo.New()
	var Server = server.New(echo_)

	Server.Start()
}
