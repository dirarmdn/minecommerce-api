package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IndexOrders(ctx *gin.Context) {
	var orders []models.Orders

	err := config.DB.Preload("Product").Preload("User").Find(&orders).Error
	if err != nil {
		log.Println("Error retrieving orders data")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, orders)

}

func StoreOrder(ctx *gin.Context) {
	var order models.Orders

	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var product models.Products
	result := config.DB.Where("ID = ?", order.ProductId).First(&product)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	newOrder := models.Orders{
		ProductId: int(product.Id),
		UserId:    order.UserId,
		OrderDate: time.Now(),
	}

	err = config.DB.Create(&newOrder).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "data successfuly added")

}
