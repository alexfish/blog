package models

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "time"
)

type Post struct {
  Model
  Id        bson.ObjectId `bson:"_id,omitempty"`
  Title     string        `bson:"Title"`
  Body      string        `bson:"Body"`
  Date      time.Time     `bson:"Date"`
}

func GetPostsByDate(s *mgo.Session, limit int) []*Post {
  posts := []*Post{}
  a := new(Post)
  query := Collection(a, s).Find(nil).Sort("{'Date': 1}").Limit(limit)
  query.All(&posts)

  return posts
}
