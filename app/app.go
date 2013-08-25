package app

import (
  "github.com/robfig/revel"
)

var (
  DB string
)

func AppInit() {
  RegisterDB()
}

func RegisterDB() {
  var found bool
  if DB, found = revel.Config.String("blog.db"); !found {
    DB = "blog"
  }
}
