package main

import (
	"minecommerce-api/config"
	"minecommerce-api/controllers"

	"net/http"

	_ "minecommerce-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mini E-commerce REST API
// @description Mini E-commerce REST API Documentation
// @version 0.1
// @host: localhost:8001

func main() {
	config.DBConnect()
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "It's working!")
	})

	// Products
	g.GET("/products", controllers.IndexProduct)
	g.GET("/product/:id", controllers.ShowProduct)
	g.POST("/product", controllers.StoreProduct)
	g.PATCH("/product/:id", controllers.UpdateProduct)
	g.DELETE("/product/:id", controllers.DeleteProduct)

	// Orders
	g.GET("/orders", controllers.IndexOrders)
	g.GET("/order/:id", controllers.ShowOrder)
	g.POST("/order", controllers.StoreOrder)
	g.PATCH("/order/:id", controllers.UpdateOrder)
	g.DELETE("order/:id", controllers.DeleteOrder)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":8001")

}
