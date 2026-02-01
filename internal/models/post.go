package models

import "time"

type PostStatus string

const (
  PostPending  PostStatus = "pending"
  PostApproved PostStatus = "approved"
  PostRemoved  PostStatus = "removed"
)

type Post struct {
  ID        string     `json:"id"`
  AuthorID  string     `json:"authorId"`
  Title     string     `json:"title"`
  Content   string     `json:"content"`
  Status    PostStatus `json:"status"`
  CreatedAt time.Time  `json:"createdAt"`
}