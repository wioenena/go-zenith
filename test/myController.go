package main

import (
	"fmt"

	"github.com/wioenena/go-zenith/controller"
	"github.com/wioenena/go-zenith/http"
)

type myController struct {
	controller.BaseController
}

func NewMyController() controller.Controller {
	c := new(myController)
	c.BaseController = *controller.NewBaseController()
	return c
}

func (m *myController) Setup() {
	m.Get("/", func(ctx http.HttpContext) error {
		fmt.Println("GET /api/")
		return fmt.Errorf("not implemented")
	})
}

func (m *myController) GetRoute() string {
	return "/api"
}
