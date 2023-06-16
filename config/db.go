package config

import (
	"fmt"
	"log"
	"minecommerce-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	conn := "host=localhost user=postgres password=123456 port=5432 dbname=db_ecommerce sslmode=disable"
	
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect", err.Error())
	} else {
		fmt.Println("DB connected")
		DB = db
	}

	db.AutoMigrate(&models.Products{}, &models.Orders{}, &models.Users{})

}