package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotel-go/internal/database"
	"hotel-go/internal/handler"
	"hotel-go/internal/repository"
	"hotel-go/internal/service"

	"github.com/gin-gonic/gin"
)

func setupRouter(t *testing.T) *gin.Engine {
	t.Helper()

	db := database.Connect()

	repo := repository.NewHotelRepository(db)
	svc := service.NewHotelService(repo)
	h := handler.NewHotelHandler(svc)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	h.RegisterRoutes(router)

	return router
}

func TestGetHotels(t *testing.T) {
	router := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/hotels", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Erwartet 200, erhalten %d", w.Code)
	}
}

func TestCreateHotel(t *testing.T) {
	router := setupRouter(t)

	body := map[string]any{
		"name": "Testhotel",
		"standort": map[string]any{
			"strasse":    "Teststraße",
			"hausnummer": "1",
			"plz":        "76133",
			"ort":        "Karlsruhe",
			"land":       "Deutschland",
		},
		"zimmer": []map[string]any{
			{"zimmernummer": "101", "preis": 99.90},
		},
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/hotels", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Erwartet 201, erhalten %d — Body: %s", w.Code, w.Body.String())
	}
}

func TestCreateHotel_ValidationError(t *testing.T) {
	router := setupRouter(t)

	// Name fehlt → Validierung muss scheitern
	body := map[string]any{
		"name": "",
		"standort": map[string]any{
			"strasse":    "Teststraße",
			"hausnummer": "1",
			"plz":        "76133",
			"ort":        "Karlsruhe",
			"land":       "Deutschland",
		},
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/hotels", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Erwartet 422, erhalten %d", w.Code)
	}
}
