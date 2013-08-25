package controllers

import (
  "github.com/robfig/revel"
  "github.com/alexfish/blog/models/model"
)

type Blog struct {
  App
}

func (c Blog) Index() revel.Result {
  posts := models.GetPostsByDate(c.MongoSession, 10)
  return c.Render(posts)
}
