package controllers

import (
	"log"
	"minecommerce-api/config"
	"minecommerce-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexProduct(ctx *gin.Context) {
	var products []models.Products

	err := config.DB.Find(&products).Error
	if err != nil {
		log.Println("Error retrieving products")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, products)

}
