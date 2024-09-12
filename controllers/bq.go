package controllers

import (
	"github.com/labstack/echo/v4"
	"myapp/domain"
	"net/http"
	"strconv"
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
	limit, limitErr := strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	if limitErr != nil {
		return c.JSON(http.StatusBadRequest, "Invalid limit")
	}
	getData, getDataErr := bc.bqSvc.GetDataBQ(limit)
	if getDataErr != nil {
		return c.JSON(http.StatusNotFound, getDataErr)
	}
	return c.JSON(http.StatusOK, getData)
}
