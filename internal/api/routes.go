package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonamaltan/go-api/internal/models"
)

// SetupRoutes creates a gin router with default middleware and sets up all routes
func SetupRoutes(db *sql.DB) *gin.Engine {
	queries := models.New(db)
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	rg := r.Group("/:serviceName/partners")
	rg.GET("", GetPartnerList(queries))
	rg.GET("/:id", GetPartnerByID(queries))

	return r
}
