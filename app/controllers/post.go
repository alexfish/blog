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
  if c.UserAuthenticated() {
    post := models.Post{}
    action := "/post/new"
    actionButton := "Create"
    return c.Render(action, post, actionButton)
  } else {
    return c.Redirect(App.Index)
  }
}

func (c Post) Show(id bson.ObjectId) revel.Result {
  authenticated := c.UserAuthenticated()
  post := models.GetPostByObjectId(c.MongoSession, id)
  return c.Render(post, authenticated)
}

func (c Post) Update(post *models.Post) revel.Result {
  if c.UserAuthenticated() {
    post.Date = time.Now()
    post.Save(c.MongoSession)
  }
  return c.Redirect(App.Index)
}

func (c Post) GetUpdate(id bson.ObjectId) revel.Result {
  if c.UserAuthenticated() {
    post := models.GetPostByObjectId(c.MongoSession, id)
    action := "/Post/Update"
    actionButton := "Update"
    return c.Render(action, post, actionButton)
  }
  return c.Redirect(App.Index)
}

func (c Post) PostCreate(post *models.Post) revel.Result {
  if c.UserAuthenticated() {
    post.Id = bson.NewObjectId()
    post.Date = time.Now()
    post.Save(c.MongoSession)
  }
  return c.Redirect(App.Index)
}

func (c Post) Delete(id bson.ObjectId) revel.Result {
  if c.UserAuthenticated() {
    post := models.GetPostByObjectId(c.MongoSession, id)
   post.Delete(c.MongoSession)
  }

  return c.Redirect(App.Index)
}
