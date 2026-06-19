package database

import (
	"fmt"
	"log"
	"os"

	"hotel-go/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=hotel",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "hotel"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Datenbankverbindung fehlgeschlagen: %v", err)
	}

	// Schema anlegen (falls es noch nicht existiert)
	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS hotel").Error; err != nil {
		log.Fatalf("Schema konnte nicht erstellt werden: %v", err)
	}

	// Tabellen automatisch erstellen bzw. aktualisieren
	if err := db.AutoMigrate(
		&model.Hotel{},
		&model.Standort{},
		&model.Zimmer{},
	); err != nil {
		log.Fatalf("Migration fehlgeschlagen: %v", err)
	}

	log.Println("Datenbankverbindung erfolgreich")
	log.Println("Schema und Tabellen erfolgreich initialisiert")

	return db
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
