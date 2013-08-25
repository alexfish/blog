package models

import (
  "github.com/robfig/revel"
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
  query := Collection(a, s).Find(nil).Sort("{'Date': -1}").Limit(limit)
  query.All(&posts)

  return posts
}

func GetPostByObjectId(s *mgo.Session, Id bson.ObjectId) *Post {
  post := new(Post)

  query := Collection(post, s).FindId(Id)
  query.One(post)
  return post
}

func (post *Post) Save(s *mgo.Session) error {
  coll := Collection(post, s)
  _, err := coll.Upsert(bson.M{"_id": post.Id}, post)
  if err != nil {
    revel.WARN.Printf("Unable to save post: %v error %v", post, err)
  }
  return err
}

func (post *Post) Delete(s *mgo.Session) error {
  coll := Collection(post, s)
  err := coll.RemoveId(post.Id)
  if err != nil {
    revel.WARN.Printf("Undable to delete post: %v error %v", post, err)
  }
  return err
}
