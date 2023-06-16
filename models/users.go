package models

import "time"

type Users struct {
	Id        int
	FullName  string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
