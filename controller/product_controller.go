package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProduct get product
func GetProduct(c *gin.Context) {
	productID := c.Param("productID")

	c.JSON(http.StatusOK, gin.H{
		"id":    productID,
		"title": "prod_" + productID,
		"price": 50,
	})
}
