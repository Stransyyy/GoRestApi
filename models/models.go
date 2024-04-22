package models

import "time"

type Users struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	TS       time.Time `json:"ts"`
	Role     []Roles   `json:"role"`
}

type Roles struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
