package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get Orders
// @Description Get list of all available Orders
// @Tags Order
// @Produce json
// @Success 200 {array} models.Orders
// @Router /orders [get]
func IndexOrders(ctx *gin.Context) {
	var orders []models.Orders

	err := config.DB.Preload("Product").Preload("User").Find(&orders).Error
	if err != nil {
		log.Println("Error retrieving orders data")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, orders)

}

// @Summary Post Order
// @Description Create new Order
// @Tags Order
// @Produce json
// @Param data body models.Orders true "Order Data"
// @Success 200 {string} data successfully added
// @Router /orders [post]
func StoreOrder(ctx *gin.Context) {
	var order models.Orders

	//  attempts to bind the JSON data from the HTTP request body to the order variable
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if product exist or no
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

	// store data
	if err := config.DB.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "data successfully added")

}
