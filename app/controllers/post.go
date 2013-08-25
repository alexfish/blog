package controllers

import (
  "blog/app/models"
  "github.com/robfig/revel"
  "labix.org/v2/mgo/bson"
  "time"
)

type Post struct {
  App
}

func (c Post) Index() revel.Result {
  return c.Redirect(Blog.Index)
}

func (c Post) GetCreate() revel.Result {
  post := models.Post{}
  action := "/post/new"
  actionButton := "Create"
  return c.Render(action, post, actionButton)
}

func (c Post) PostCreate(post *models.Post) revel.Result {
  post.Id = bson.NewObjectId()
  post.Date = time.Now()
  post.Save(c.MongoSession)

  return c.Redirect(App.Index)
}
