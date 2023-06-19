package models

import "time"

type Orders struct {
	Id           int
	ProductId    int `json:"product_id"`
	UserId       int `json:"buyer_id"`
	BuyerAddress string
	BuyerEmail   string
	OrderDate    time.Time
	UpdatedAt    time.Time
	Product      Products
	User         Users
}
