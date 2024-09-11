package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"myapp/conn"
	"myapp/controllers"
	"myapp/repositories"
	"myapp/routes"
	"myapp/server"
	"myapp/services"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// Your code here

	// Clients
	bqClient := conn.NewBqDBClient()

	// Repositories
	bqRepo := repositories.NewBQRepository(bqClient)

	// Services
	bqSvc := services.NewBQService(bqRepo)

	// Controllers
	bqCtr := controllers.NewBQController(bqSvc)

	var echo_ = echo.New()
	var Routes = routes.NewRoutes(echo_, bqCtr)
	var Server = server.New(echo_)

	Routes.Init()
	Server.Start()
}
