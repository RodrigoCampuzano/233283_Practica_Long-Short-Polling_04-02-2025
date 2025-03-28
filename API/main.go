package main

import (
	"log"
	"net/http"
	"os"
	"ApiShortLong/application"
	"ApiShortLong/core" // Importa el paquete core donde está database.go
	"ApiShortLong/infrastructure/handler"
	"ApiShortLong/infrastructure/repo"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Para manejar variables de entorno
)

func main() {
	// Cargar variables de entorno desde archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	}

	// Database connection
	db, err := database.NewMySQLConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	defer db.Close()

	// Repository
	productRepo := repo.NewMySQLRepository(db) // Cambiado el nombre para evitar conflicto

	// Service
	service := service.NewProductService(productRepo)

	// Handler
	productHandler := handler.NewProductHandler(service)

	// Router
	router := mux.NewRouter()
	productHandler.SetupRoutes(router)

	// CORS middleware mejorado
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	}

	port := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, corsMiddleware(router)))
}