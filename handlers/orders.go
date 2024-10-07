package handlers

import (
	"net/http"
	"wb/cache"

	"github.com/gin-gonic/gin"
)

func GetOrderByIdHandler(c *gin.Context) {
	id := c.Param("id")

	value, exists := cache.Cache.GetCache(id)
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, value)
}
