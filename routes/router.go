package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers"
	"myapp/methodutils"
	"net/http"
)

type Routes struct {
	echo  *echo.Echo
	bqCtr controllers.BQController
	pgCtr controllers.PostgresController
}

func NewRoutes(echo *echo.Echo, bqCtr controllers.BQController, pgCtr controllers.PostgresController) *Routes {
	return &Routes{
		echo:  echo,
		bqCtr: bqCtr,
		pgCtr: pgCtr,
	}
}

func (r *Routes) Init() {
	e := r.echo
	e.Use(methodutils.LogAPICall)
	g := e.Group("/api/v1")

	g.GET("/health", Health)
	g.GET("/bq/get", r.bqCtr.Get)
	g.GET("/pg/get", r.pgCtr.Get)
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "I am alive !!")
}
