package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"tpIRSO/database"
	"tpIRSO/handlers"
	"tpIRSO/models"
	"tpIRSO/repositories"
	"tpIRSO/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	logger := log.New(os.Stdout, "API ", log.LstdFlags)

	// Inicializar la conexión a la base de datos
	db, err := database.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate de tablas
	err = db.AutoMigrate(&models.Vehicle{}, &models.User{})
	if err != nil {
		logger.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Tablas vehicles y users creadas exitosamente!")

	// Inicializar repositorios
	vehicleRepo := repositories.NewRepoVehicle(db, logger)
	userRepo := repositories.NewRepoUser(db, logger)

	// Inicializar servicios
	vehicleService := services.NewVehicleService(vehicleRepo)
	userService := services.NewUserService(userRepo)

	// Inicializar handlers
	vehicleHandler := handlers.NewVehicleHandler(*vehicleService)
	userHandler := handlers.NewUserHandler(userService)

	// Crear el router
	router := mux.NewRouter()

	// Configurar las rutas de vehicles
	router.HandleFunc("/api/vehicles", vehicleHandler.CreateVehicle).Methods("POST")
	router.HandleFunc("/api/vehicles/{id}", vehicleHandler.GetVehicleById).Methods("GET")
	router.HandleFunc("/api/vehicles/{id}", vehicleHandler.UpdateVehicle).Methods("PUT")

	// Configurar las rutas de users
	router.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PUT")

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         300,
	})

	// Crear el servidor
	server := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler.Handler(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Iniciar el servidor
	logger.Println("Servidor iniciando en http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("Error iniciando el servidor:", err)
	}
}
