package models

import (
  "blog/app"
  "github.com/robfig/revel"
  "labix.org/v2/mgo"
  "reflect"
  "strings"
)

var (
  Collections map[string]string
)

// Empty struct to embed in models that will provide application default funcs.
type Model struct{}

func Collection(m interface{}, s *mgo.Session) *mgo.Collection {
  typ := reflect.TypeOf(m)
  i := strings.LastIndex(typ.String(), ".") + 1
  n := typ.String()[i:len(typ.String())]

  var found bool
  var c string
  if c, found = revel.Config.String("blog.db.collection." + n); !found {
    c = n
  }
  return s.DB(app.DB).C(c)
}
