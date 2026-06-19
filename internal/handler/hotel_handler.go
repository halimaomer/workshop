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
			"error": "Hotels konnten nicht geladen werden",
		})
		return
	}

	c.JSON(http.StatusOK, hotels)
}

func (h *HotelHandler) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ungültige ID",
		})
		return
	}

	hotel, err := h.service.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Hotel nicht gefunden",
		})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func (h *HotelHandler) Create(c *gin.Context) {
	var req CreateHotelRequest

	// JSON einlesen
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validieren
	if err := h.validate.Struct(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Request -> Entity
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

	// Speichern
	if err := h.service.Create(&hotel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Location-Header setzen
	c.Header("Location", "/hotels/"+strconv.Itoa(int(hotel.ID)))

	c.JSON(http.StatusCreated, hotel)
}
