# Programmierworkshop am 19.6.2026

## Namen

Halima Omer & Alisa Knöbl

## Link zum Git-Repository

https://github.com/halimaomer/workshop/tree/main

## KI-Werkzeuge

* ChatGPT, Claude

### Agenten

* ChatGPT (GPT-5.5)
* Sonnet (4.6)

### Chat-URLs, z.B. https://chatgpt.com

* https://chatgpt.com

* https://claude.ai/

## Frameworks und Bibliotheken

* Go
* Gin
* GORM
* PostgreSQL Driver (`gorm.io/driver/postgres`)
* go-playground/validator/v10

### REST-Schnittstelle (Lesen und Neuanlegen)

Implementierte REST-Endpunkte:

* `GET /hotels` – Alle Hotels abrufen
* `GET /hotels/{id}` – Ein Hotel anhand der ID abrufen
* `POST /hotels` – Neues Hotel anlegen

### Validierung (nur Neuanlegen)

Für das Anlegen eines Hotels wird `go-playground/validator/v10` verwendet.

Validiert werden unter anderem:

* Hotelname (Pflichtfeld)
* Straße (Pflichtfeld)
* Hausnummer (Pflichtfeld)
* PLZ (5 Zeichen)
* Ort (Pflichtfeld)
* Land (Pflichtfeld)
* Zimmernummer (Pflichtfeld)
* Preis (> 0)

Bei ungültigen Eingaben werden entsprechende HTTP-Fehlercodes und Fehlermeldungen zurückgegeben.

### OR-Mapping (für PostgreSQL)

Als ORM wird GORM verwendet.

Die Datenbanktabellen werden automatisch über `AutoMigrate()` erstellt.

Folgende Entitäten wurden implementiert:

* Hotel
* Standort
* Zimmer

### Optional: OIDC mit Keycloak

Nicht implementiert (optional).

### Einfacher Integrationstest

Es wurden Integrationstests für die REST-Schnittstelle erstellt.

Getestet werden:

* `GET /hotels`
* `POST /hotels`
* Validierungsfehler beim Anlegen eines Hotels

Ausführung:

```bash
go test ./tests/... -v
```

## Prompts/Requests an KI-Agent/en

Prompts / Requests an KI-Agenten
Welches Framework eignet sich für die Entwicklung eines REST-Servers in Go?
Welche Standardbibliotheken und Tools werden in Go-Projekten häufig verwendet?
Wie beginne ich ein neues Go-Projekt mit Gin?
Wie sollte eine typische main.go in Go aussehen?
Ist ein src-Verzeichnis in Go notwendig?
Ist meine Projektstruktur für ein Go-Projekt sinnvoll?
Wie werden Entitäten (Structs) in Go erstellt?
Wie modelliert man eine 1:1- und 1 mit GORM?
Was bedeuten Struct Tags wie json, gorm und binding?
Wie funktioniert die Validierung von Requests in Gin?
Wie arbeitet GORM und welche Funktionen sind für CRUD-Operationen wichtig?
Sind meine GORM-Modelle sinnvoll aufgebaut und wie können sie verbessert werden?
Wie sieht eine Datenbankverbindung mit GORM und PostgreSQL aus?
Wie wird ein Repository in Go aufgebaut?
Ist meine main.go korrekt aufgebaut?
Wie implementiert man Handler, Services und Repositories nach Go-Konventionen?
Wie werden DTOs sinnvoll strukturiert und verwendet?
Wie testet man REST-Endpunkte mit PowerShell?
Wie schreibt man Integrationstests in Go?
Welche Werkzeuge gibt es für Codeformatierung, Codeanalyse und Qualitätssicherung (goimports, golangci-lint, go vet, staticcheck)?
Welche Bedeutung haben Befehle wie go build ./... und go test?
Welche Dateien und Konfigurationen werden für ein Go-Projekt zusätzlich benötigt (z. B. .gitignore, Docker, Docker Compose)?
Wie kann ein ASCII-Art-Startbanner für den Server und für Testläufe erstellt werden?
