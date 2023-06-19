package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"minecommerce-api/services"
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

	if err := config.DB.Preload("Product").Preload("User").Find(&orders).Error; err != nil {
		log.Println("error retrieving orders data", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve orders",
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)

}

// @Summary Get Order by ID
// @Description Get Order by ID
// @Tags Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Orders
// @Router /order/{id} [get]
func ShowOrder(ctx *gin.Context) {
	var order models.Orders

	err := config.DB.Preload("Product").Preload("User").Where("ID = ?", ctx.Param("id")).Find(&order).Error
	if err != nil {
		log.Println("error retrieving order data", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "failed to retrieve order",
		})
		return
	}

	ctx.JSON(http.StatusOK, order)

}

type StoreOrderInput struct {
	ProductId int `json:"product_id"`
	UserId    int `json:"buyer_id"`
}

// @Summary Post Order
// @Description Store new Order
// @Tags Order
// @Produce json
// @Param data body StoreOrderInput true "Order Data"
// @Success 200 {string} data successfully added
// @Router /order [post]
func StoreOrder(ctx *gin.Context) {
	var input StoreOrderInput
	var product models.Products
	var user models.Users

	//  attempts to bind the JSON data from the HTTP request body to the order variable
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if product exist or no
	result := config.DB.Where("ID = ?", input.ProductId).First(&product)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	// check if user exist or no
	res := config.DB.Where("ID = ?", input.UserId).First(&user)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	newOrder := models.Orders{
		ProductId:    int(product.Id),
		UserId:       input.UserId,
		BuyerAddress: user.Address,
		BuyerEmail:   user.Email,
		OrderDate:    time.Now(),
	}

	// store data
	if err := config.DB.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "failed to store data",
		})
		return
	}

	services.SendMail(newOrder.BuyerEmail, newOrder.BuyerAddress, product.Name)

	ctx.JSON(http.StatusOK, "data successfully added")

}

type UpdateOrderInput struct {
	ProductId int `json:"product_id"`
}

// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Produce json
// @Param id path int true "Order ID"
// @Param data body UpdateOrderInput true "Order data"
// @Success 200 {string} data successfully updated
// @Router /order/{id} [patch]
func UpdateOrder(ctx *gin.Context) {
	var order models.Orders
	var input UpdateOrderInput

	if err := config.DB.Where("ID = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "order not found",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Model(&order).Updates(input).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, "data successfully updated")

}

// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {string} data successfully deleted
// @Router /order/{id} [delete]
func DeleteOrder(ctx *gin.Context) {
	var order models.Orders

	if err := config.DB.Where("ID = ?", ctx.Param("id")).First(&order).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "order not found",
		})
		return
	}

	if err := config.DB.Delete(order).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, "data successfully deleted")

}
