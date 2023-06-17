package models

import "time"

type Products struct {
	Id          int
	Name        string `json:"product_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
