package main

import (
	"log"
	"net/http"

	"hotel-go/internal/database"
	"hotel-go/internal/handler"
	"hotel-go/internal/repository"
	"hotel-go/internal/service"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

func main() {

	figure.NewFigure("HOTEL", "doom", true).Print()

	log.Println("────────────────────────────────────────────")
	log.Println(" Server läuft auf http://localhost:8080")
	log.Println(" Datenbank verbunden")
	log.Println(" Umgebung: Development")
	log.Println("────────────────────────────────────────────")
	// Datenbank verbinden
	db := database.Connect()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Fehler beim Zugriff auf die Datenbank: %v", err)
	}
	defer sqlDB.Close()

	// Repository -> Service -> Handler
	hotelRepository := repository.NewHotelRepository(db)
	hotelService := service.NewHotelService(hotelRepository)
	hotelHandler := handler.NewHotelHandler(hotelService)

	// Router
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "Hotel API",
			"version": "1.0.0",
			"status":  "running",
		})
	})
	// Hotel-Routen
	hotelHandler.RegisterRoutes(router)

	// Server starten
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server konnte nicht gestartet werden: %v", err)
	}
}
