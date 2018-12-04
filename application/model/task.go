package model

import (
	_ "github.com/lib/pq"
)

type Task struct {
	ID          int    `json:"id" xml:"id"`
	Title       string `json:"title"  xml:"title"`
	Description string `json:"description" xml:"description"`
	Done        bool   `json:"done" xml:"done"`
	OwnerID     int    `json:"owner_id" xml:"owner_id"`
}
