package main

import (
	"candystore/candystore"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/customers/top", topCustomers)

	router.Run("localhost:8181")
}

func topCustomers(c *gin.Context) {
	topCustomers := candystore.TopCustomers()

	c.IndentedJSON(http.StatusOK, topCustomers)
}
