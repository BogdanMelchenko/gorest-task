package model

import (
	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Role int    `json:"role" xml:"role"`
}
