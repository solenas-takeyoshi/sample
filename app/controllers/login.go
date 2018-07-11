package controllers

import (
	"github.com/revel/revel"
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	return c.Render()
}
