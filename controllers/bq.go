package controllers

import (
	"github.com/labstack/echo/v4"
	"myapp/domain"
	"net/http"
)

type BQController struct {
	bqSvc domain.BQUseCase
}

func NewBQController(bqSvc domain.BQUseCase) BQController {
	return BQController{
		bqSvc: bqSvc,
	}
}

func (bc *BQController) Get(c echo.Context) error {
	// Your code here
	var payload domain.Payload
	uuid := c.QueryParam("uuid")
	date := c.QueryParam("date")
	limit := c.QueryParam("limit")

	payload.UUID = uuid
	payload.Date = date
	payload.Limit = limit

	getData, getDataErr := bc.bqSvc.GetDataBQ(payload)
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
