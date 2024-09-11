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
	postgresClient := conn.NewPostgresClient()

	// Repositories
	bqRepo := repositories.NewBQRepository(bqClient)
	postgresRepo := repositories.NewPostgresRepository(postgresClient)

	// Services
	bqSvc := services.NewBQService(&bqRepo)
	pgSvc := services.NewPostgresService(postgresRepo)

	// Controllers
	bqCtr := controllers.NewBQController(bqSvc)
	pgCtr := controllers.NewPostgresController(pgSvc)

	var echo_ = echo.New()
	var Routes = routes.NewRoutes(echo_, bqCtr, pgCtr)
	var Server = server.New(echo_)

	Routes.Init()
	Server.Start()
}
