package controllers

import (
  "github.com/robfig/revel"
)

type Post struct {
  App
}

func (c Post) Index() revel.Result {
  return c.Redirect(Blog.Index)
}
