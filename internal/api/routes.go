package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRoutes creates a gin router with default middleware and sets up all routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	rg := r.Group("/customers")
	rg.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "customerList",
		})
	})
	rg.POST("")

	return r
}
