package handler

import (
	"net/http"
	"strconv"

	"hotel-go/internal/model"
	"hotel-go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HotelHandler struct {
	service  *service.HotelService
	validate *validator.Validate
}

func NewHotelHandler(service *service.HotelService) *HotelHandler {
	return &HotelHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *HotelHandler) RegisterRoutes(router *gin.Engine) {
	hotels := router.Group("/hotels")

	hotels.GET("", h.FindAll)
	hotels.GET("/:id", h.FindByID)
	hotels.POST("", h.Create)
}

func (h *HotelHandler) FindAll(c *gin.Context) {
	hotels, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Hotels konnten nicht geladen werden.",
		})
		return
	}

	c.JSON(http.StatusOK, hotels)
}

func (h *HotelHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ungültige Hotel-ID.",
		})
		return
	}

	hotel, err := h.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Hotel wurde nicht gefunden.",
		})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func (h *HotelHandler) Create(c *gin.Context) {
	var req CreateHotelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Ungültige Anfrage. Bitte JSON-Format und Datentypen prüfen.",
			"details": err.Error(),
		})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validierung fehlgeschlagen.",
			"details": err.Error(),
		})
		return
	}

	hotel := model.Hotel{
		Name: req.Name,
		Standort: &model.Standort{
			Strasse:    req.Standort.Strasse,
			Hausnummer: req.Standort.Hausnummer,
			PLZ:        req.Standort.PLZ,
			Ort:        req.Standort.Ort,
			Land:       req.Standort.Land,
		},
	}

	for _, z := range req.Zimmer {
		hotel.Zimmer = append(hotel.Zimmer, model.Zimmer{
			Zimmernummer: z.Zimmernummer,
			Preis:        z.Preis,
		})
	}

	if err := h.service.Create(&hotel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Hotel konnte nicht gespeichert werden.",
			"details": err.Error(),
		})
		return
	}

	c.Header("Location", "/hotels/"+strconv.Itoa(int(hotel.ID)))
	c.JSON(http.StatusCreated, hotel)
}
