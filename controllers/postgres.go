package controllers

import (
	"github.com/labstack/echo/v4"
	"myapp/domain"
	"net/http"
)

type PostgresController struct {
	postgresSvc domain.PostUseCase
}

func NewPostgresController(postgresSvc domain.PostUseCase) PostgresController {
	return PostgresController{
		postgresSvc: postgresSvc,
	}
}

func (pc *PostgresController) Get(c echo.Context) error {
	// Your code here
	getData, getDataErr := pc.postgresSvc.Get()
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
