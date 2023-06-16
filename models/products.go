package models

import "time"

type Products struct {
	Id          int
	Name        string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
