package api_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramonamaltan/go-api/internal/api"
	"github.com/ramonamaltan/go-api/internal/db"
	"github.com/ramonamaltan/go-api/internal/models"
	"github.com/stretchr/testify/require"
)

func TestGetPartnerList(t *testing.T) {
	db := db.Init()
	defer db.Close()
	r := api.SetupRoutes(db)

	req, err := http.NewRequest("GET", "/flooring/partners", nil)
	require.NoError(t, err)
	q := req.URL.Query()
	q.Add("long", "13.85229")
	q.Add("lat", "52.55360")
	q.Add("material", "wood")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.NotEmpty(t, responseData)
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetPartner(t *testing.T) {
	db := db.Init()
	defer db.Close()
	r := api.SetupRoutes(db)
	queries := models.New(db)
	t.Run("success", func(t *testing.T) {
		partners, err := queries.ListPartners(context.Background(), models.ListPartnersParams{Servicename: "flooring", Material: "wood"})
		lastID := partners[len(partners)-1].ID

		req, err := http.NewRequest("GET", fmt.Sprintf("/flooring/partners/%v", lastID), nil)
		require.NoError(t, err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		require.NotEmpty(t, responseData)
		require.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/flooring/partners/1500", nil)
		require.NoError(t, err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}
