package controllers

import (
	"github.com/labstack/echo/v4"
	"myapp/domain"
	"net/http"
	"strconv"
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
	limit, limitErr := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if limitErr != nil {
		return c.JSON(http.StatusBadRequest, "Invalid limit")
	}
	getData, getDataErr := pc.postgresSvc.GetFromPG(limit)
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
