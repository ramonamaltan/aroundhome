package api_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramonamaltan/go-api/internal/api"
	"github.com/ramonamaltan/go-api/internal/db"
	"github.com/stretchr/testify/require"
)

func TestGetPartnerList(t *testing.T) {
	db := db.Init()
	defer db.Close()
	r := api.SetupRoutes(db)

	req, _ := http.NewRequest("GET", "/partners", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	fmt.Println(responseData)
	//require.Equal(t, mockResponse, string(responseData))
	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetPartner(t *testing.T) {
	db := db.Init()
	defer db.Close()
	r := api.SetupRoutes(db)

	req, _ := http.NewRequest("GET", "/partners/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	fmt.Println(responseData)
	//require.Equal(t, mockResponse, string(responseData))
	require.Equal(t, http.StatusOK, w.Code)
}
