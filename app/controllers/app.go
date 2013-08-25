package controllers

import (
  "github.com/robfig/revel"
  "github.com/jgraham909/revmgo"
)

type App struct {
	*revel.Controller
  revmgo.MongoController
}

func (c App) Index() revel.Result {
  return c.Redirect(Blog.Index)
}

func (c App) UserAuthenticated() bool {
  if _, ok := c.Session["user"]; ok {
    return true
  }
  return false
}
