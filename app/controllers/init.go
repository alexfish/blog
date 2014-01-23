package controllers

import (
  "github.com/alexfish/revmgo"
  "github.com/robfig/revel"
  "time"
)

func init() {
    revmgo.ControllerInit()
    revel.TemplateFuncs["year"] = func() string {
      t := time.Now()
      return t.Format("2006")
    }
}
