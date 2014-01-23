package controllers

import (
  "code.google.com/p/go.crypto/bcrypt"
  "github.com/alexfish/blog/app/models"
  "github.com/robfig/revel"
)

type User struct {
  App
}

func (c User) GetLogin() revel.Result {
  if c.UserAuthenticated() == false {
    return c.Render()
  }

  // User already logged in bounce to main page
  return c.Redirect(App.Index)
}

func (c User) PostLogin(Email, Password string) revel.Result {
  user := models.GetUserByEmail(c.MongoSession, Email)

  if user.Email != "" {
    err := bcrypt.CompareHashAndPassword(user.Password, []byte(Password))
    if err == nil {
      c.Session["user"] = Email
      c.Flash.Success("Welcome, " + user.Name)
      return c.Redirect(App.Index)
    }
  }

  c.Flash.Out["mail"] = Email
  c.Flash.Error("Incorrect email address or password.")
  return c.Redirect(User.GetLogin)
}

func (c User) GetLogout() revel.Result {
  for k := range c.Session {
    delete(c.Session, k)
  }
  return c.Redirect(App.Index)
}
