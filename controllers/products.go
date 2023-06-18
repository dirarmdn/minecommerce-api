package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Products
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} models.Products
// @Router /products [get]
func IndexProduct(ctx *gin.Context) {
	var products []models.Products

	if err := config.DB.Find(&products).Error; err != nil {
		log.Println("error retrieving products", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve products",
		})
		return
	}

	ctx.JSON(http.StatusOK, products)

}

// @Summary Get Product by id
// @Description Show Product by ID
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Products
// @Router /product/{id} [get]
func ShowProduct(ctx *gin.Context) {
	var product models.Products

	if err := config.DB.Where("ID = ?", ctx.Param("id")).First(&product).Error; err != nil {
		log.Println("error retrieving product", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "failed to retrieve product",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)

}

type StoreProductInput struct {
	Name        string `json:"product_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// @Summary Store Product
// @Description Create new Product
// @Tags Product
// @Produce json
// @Param data body StoreProductInput true "product data"
// @Success 200 {string} data successfully added
// @Router /product [post]
func StoreProduct(ctx *gin.Context) {
	var input StoreProductInput

	if err := ctx.BindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	data := models.Products{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	if err := config.DB.Create(&data).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to store data",
		})
		return
	}

	ctx.JSON(http.StatusOK, "data successfully added")

}

type UpdateProductInput struct {
	Name        string `json:"product_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Param data body UpdateProductInput true "Product data"
// @Success 200 {string} string data successfully updated
// @Router /product/{id} [patch]
func UpdateProduct(ctx *gin.Context) {
	var product models.Products
	var input UpdateProductInput

	if err := config.DB.Where("ID = ?", ctx.Param("id")).First(&product).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Model(&product).Updates(input).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to update data",
		})
		return
	}

	ctx.JSON(http.StatusOK, "data successfully updated")

}

// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {string} data successfully deleted
// @Router /product/{id} [delete]
func DeleteProduct(ctx *gin.Context) {
	var product models.Products

	if err := config.DB.Where("ID = ?", ctx.Param("id")).First(&product).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	if err := config.DB.Delete(product).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "data successfully deleted")

}
