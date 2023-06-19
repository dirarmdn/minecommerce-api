package models

import "time"

type Users struct {
	Id        int
	Username  string
	Email     string
	Password  string
	FullName  string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
