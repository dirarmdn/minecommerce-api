package main

import (
	"minecommerce-api/config"
	"minecommerce-api/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnect()
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "It's working!")
	})

	// Products
	g.GET("/products", controllers.IndexProduct)

	// Orders
	g.GET("/orders", controllers.IndexOrders)
	g.POST("/orders", controllers.StoreOrder)

	g.Run(":8001")
}
