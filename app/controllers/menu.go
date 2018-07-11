package controllers

import (
	"github.com/revel/revel"
)

type Menu struct {
	*revel.Controller
}

func (c Menu) Index() revel.Result {
	return c.Render()
}
