package api

import (
	"database/sql"
	"errors"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramonamaltan/go-api/internal/models"
	"github.com/umahmood/haversine"
)

type PartnerResponse struct {
	ID           int64   `json:"id"`
	Material     string  `json:"material"`
	Address      Address `json:"address"`
	Radius       int32   `json:"operating_radius"`
	Rating       float64 `json:"rating"`
	KmToCustomer float64 `json:"km_to_customer"`
}

type Address struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func GetPartnerList(queries *models.Queries) func(c *gin.Context) {
	return func(c *gin.Context) {
		serviceName := c.Param("serviceName")
		material := c.Query("material")
		longitude := c.Query("long")
		latitude := c.Query("lat")
		if material == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing query param serviceName or material"})
			return
		}
		long, err := strconv.ParseFloat(longitude, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "missing or invalid long value"})
			return
		}
		lat, err := strconv.ParseFloat(latitude, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "missing or invalid lat value"})
			return
		}
		partners, err := queries.ListPartners(c.Request.Context(), models.ListPartnersParams{Servicename: serviceName, Material: "%" + material + "%"})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		customerCoord := haversine.Coord{Lat: lat, Lon: long}

		var partnerList []PartnerResponse
		for _, partner := range partners {
			partnerLat := partner.Latitude
			partnerLong := partner.Longitude
			partnerCoord := haversine.Coord{Lat: partnerLat, Lon: partnerLong}
			_, km := haversine.Distance(customerCoord, partnerCoord)
			// check distance to customer is within partner radius
			if km <= float64(partner.Radius) {
				partnerMatch := PartnerResponse{
					ID:       partner.ID,
					Material: partner.Material,
					Address: Address{
						Lat:  partner.Latitude,
						Long: partner.Longitude,
					},
					Radius:       partner.Radius,
					Rating:       partner.Rating,
					KmToCustomer: km,
				}
				partnerList = append(partnerList, partnerMatch)
			}
		}

		// sort by rating
		sort.Slice(partnerList, func(i, j int) bool { return partnerList[i].Rating > partnerList[j].Rating })

		c.JSON(http.StatusOK, partnerList)
	}
}

func GetPartnerByID(queries *models.Queries) func(c *gin.Context) {
	return func(c *gin.Context) {
		partnerID := c.Param("id")
		id, err := strconv.ParseInt(partnerID, 10, 64)
		partner, err := queries.GetPartner(c.Request.Context(), id)
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "partner not found"})
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get partner"})
			return
		}

		c.JSON(http.StatusOK, partner)
	}
}
