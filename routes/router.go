package routes

import (
	"github.com/labstack/echo/v4"
	"myapp/controllers"
	"net/http"
)

type Routes struct {
	echo  *echo.Echo
	bqCtr controllers.BQController
}

func NewRoutes(echo *echo.Echo, bqCtr controllers.BQController) *Routes {
	return &Routes{
		echo:  echo,
		bqCtr: bqCtr,
	}
}

func (r *Routes) Init() {
	e := r.echo
	g := e.Group("/api/v1")

	g.GET("/health", Health)
	g.GET("/bq/get", r.bqCtr.Get)
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "I am alive !!")
}
