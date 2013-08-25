package models

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type User struct {
  Model
  Id          bson.ObjectId `bson:"_id,omitempty"`
  Name        string        `bson:"Name"`
  Email       string        `bson:"Email"`
  Password    []byte        `bson:"Password"`
}

type Password struct {
  Pass        string
  PassConfirm string
}

func GetUserByObjectId(s *mgo.Session, Id bson.ObjectId) *User {
  u := new(User)
  query := Collection(u, s).FindId(Id)
  query.One(u)
  return u
}

func GetUserByEmail(s *mgo.Session, Email string) *User {
  acct := new(User)
  query := Collection(acct, s).Find(bson.M{"Email": Email})
  query.One(acct)

  return acct
}
