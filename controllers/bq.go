package controllers

import "myapp/services"

type BQController struct {
	bqSvc services.BQService
}

func NewBQController(bqSvc services.BQService) BQController {
	return BQController{
		bqSvc: bqSvc,
	}
}
