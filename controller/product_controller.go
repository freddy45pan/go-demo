package controller

import (
	"net/http"

	"github.com/freddy45pan/go-demo-core/product"
	"github.com/gin-gonic/gin"
)

// GetProduct get product
func GetProduct(c *gin.Context) {
	product := new(product.Product)
	product.ID = c.Param("productID")
	err := product.Read()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
