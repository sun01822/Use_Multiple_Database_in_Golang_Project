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
	var payload domain.Payload
	uuid := c.QueryParam("uuid")
	date := c.QueryParam("date")
	limit := c.QueryParam("limit")

	payload.UUID = uuid
	payload.Date = date
	payload.Limit = limit

	getData, getDataErr := pc.postgresSvc.GetFromPG(payload)
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
