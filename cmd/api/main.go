package main

import (
	"log"
	"net/http"

	"hotel-go/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Fehler beim Zugriff auf die Datenbank: %v", err)
	}

	defer sqlDB.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server konnte nicht gestartet werden: %v", err)
	}
}
package main

import (
	"log"
	"net/http"

	"hotel-go/internal/database"
	"hotel-go/internal/handler"
	"hotel-go/internal/repository"
	"hotel-go/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
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

	// Hotel-Routen
	hotelHandler.RegisterRoutes(router)

	// Server starten
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server konnte nicht gestartet werden: %v", err)
	}
}