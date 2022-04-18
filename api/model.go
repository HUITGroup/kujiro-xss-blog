package main

import "time"

type Post struct {
	ID       uint64    `json:"post_id,omitempty" db:"id"`
	Title    string    `json:"title" db:"title"`
	Content  string    `json:"content" db:"content"`
	CreateAt time.Time `json:"timestamp,omitempty" db:"create_at"`
}

type Comment struct {
	ID       uint64    `json:"comment_id,omitempty" db:"id"`
	Name     string    `json:"name" db:"poster"`
	Content  string    `json:"content" db:"content"`
	PostID   uint64    `json:"post_id" db:"post_id"`
	CreateAt time.Time `json:"timestamp,omitempty" db:"create_at"`
}
