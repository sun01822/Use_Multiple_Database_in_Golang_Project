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
	getData, getDataErr := bc.bqSvc.Get()
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
