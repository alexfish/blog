package controllers

import (
  "github.com/robfig/revel"
  "github.com/alexefish/revmgo"
  "github.com/knieriem/markdown"
  "strings"
  "bytes"
)

type App struct {
	*revel.Controller
  revmgo.MongoController
}

func (c App) Index() revel.Result {
  return c.Redirect(Post.Index)
}

func (c App) UserAuthenticated() bool {
  if _, ok := c.Session["user"]; ok {
    return true
  }
  return false
}

func (c App) MarkdownHTML(s string) string {
  parser := markdown.NewParser(nil)
  src := strings.NewReader(s)
  dst := new(bytes.Buffer)
  parser.Markdown(src, markdown.ToHTML(dst))
  return dst.String()
}
