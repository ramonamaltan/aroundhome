package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramonamaltan/go-api/internal/models"
)

type Partner struct {
	Material string `json:"material"`
	Address  struct {
		Lat  string `json:"lat"`
		Long string `json:"long"`
	} `json:"address"`
	OperatingRadius int `json:"operatingRadius"`
}

func GetPartnerList(queries *models.Queries) func(c *gin.Context) {
	return func(c *gin.Context) {
		partners, err := queries.ListPartners(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		c.JSON(http.StatusOK, partners)
	}
}

func GetPartnerByID(queries *models.Queries) func(c *gin.Context) {
	return func(c *gin.Context) {
		partnerID := c.Param("id")
		id, err := strconv.ParseInt(partnerID, 10, 64)
		partner, err := queries.GetPartner(c.Request.Context(), sql.NullInt64{
			Int64: id,
			Valid: true,
		})
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, partner)
	}
}
