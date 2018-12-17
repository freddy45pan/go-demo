package main

import (
	"github.com/freddy45pan/go-demo/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products/:productID", controller.GetProduct)

	router.Run()
}
