package models

import (
  "github.com/robfig/revel"
  "time"
)

type Post struct {
  PostId    int
  Title     string
  Body      string
  Date      time.Time
}
