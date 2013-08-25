package controllers

import "github.com/robfig/revel"

type Blog struct {
  App
}

func (c Blog) Index() revel.Result {
  return c.Render()
}
