package controllers

import (
  "github.com/alexefish/blog/app/models"
  "github.com/robfig/revel"
  "labix.org/v2/mgo/bson"
  "time"
  "strconv"
)

type Post struct {
  App
}

func (c Post) Index() revel.Result {
  authenticated := c.UserAuthenticated()

  perPage := 1
  page := pageWithString(c.Params.Get("page"))
  posts := models.GetPostsByDate(c.MongoSession, perPage, page)
  total := models.TotalPostCount(c.MongoSession)

  nextPage := page + 1
  prevPage := page - 1

  if page * perPage >= total {
    nextPage = 0
  }

  for i := range posts {
    posts[i].Body = c.MarkdownHTML(posts[i].Body)
  }
  return c.Render(posts, authenticated, nextPage, prevPage)
}

func pageWithString(p string) int {
  var page int
  if len(p) != 0 {
    pi, err := strconv.Atoi(p);
    if err != nil {
      page = 1
    } else {
      page = pi
    }
  } else {
    page = 1
  }

  return page
}

func (c Post) Show(id bson.ObjectId) revel.Result {
  authenticated := c.UserAuthenticated()
  post := models.GetPostByObjectId(c.MongoSession, id)
  post.Body = c.MarkdownHTML(post.Body)
  return c.Render(post, authenticated)
}

func (c Post) Update(post *models.Post) revel.Result {
  if c.UserAuthenticated() {
    post.Validate(c.Validation)
    if c.Validation.HasErrors() {
      c.Validation.Keep()
      c.FlashParams()
      return c.Redirect(Post.GetUpdate)
    }
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

func (c Post) GetUpdate(id bson.ObjectId) revel.Result {
  if c.UserAuthenticated() {
    post := models.GetPostByObjectId(c.MongoSession, id)
    action := "/Post/Update"
    actionButton := "Update"
    return c.Render(action, post, actionButton)
  }
  return c.Redirect(App.Index)
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

func (c Post) PostCreate(post *models.Post) revel.Result {
  if c.UserAuthenticated() {
    post.Validate(c.Validation)
    if c.Validation.HasErrors() {
      c.Validation.Keep()
      c.FlashParams()
      return c.Redirect(Post.PostCreate)
    }
    post.Id = bson.NewObjectId()
    post.Date = time.Now()
    post.Save(c.MongoSession)
  }
  return c.Redirect(App.Index)
}
