package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} models.Products
// @Router /products [get]
func IndexProduct(ctx *gin.Context) {
	var products []models.Products

	if err := config.DB.Find(&products).Error; err != nil {
		log.Println("Error retrieving products")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)

}

// @Summary Post Product
// @Description Create new Product
// @Tags Product
// @Produce json
// @Param data body models.Products true "product data"
// @Success 200 {string} data successfully added
// @Router /products [post]
func StoreProduct(ctx *gin.Context) {
	var product models.Products

	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	data := models.Products{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "data successfully added")

}
