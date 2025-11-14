package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"tpIRSO/database"
	"tpIRSO/handlers"
	"tpIRSO/models"
	"tpIRSO/repositories"
	"tpIRSO/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type App struct {
	Logger   *log.Logger
	Validate *validator.Validate

	UserService *services.UserService
	UserRepo    repositories.UserRepository
	UserModel   *models.User
	UserHandler *handlers.UserHandler

	VehicleHandler *handlers.VehicleHandler
	VehicleModel   *models.Vehicle
	VehicleRepo    repositories.VehicleRepo
	VehicleService *services.VehicleService

	ContactHandler *handlers.ContactHandler
	ContactRepo    repositories.ContactRepository
	ContactService *services.ContactService

	Server *http.Server
}

func NewApp() *App {
	logger := log.New(os.Stdout, "API ", log.LstdFlags)

	db, err := database.ConnectDB()
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}

	err = db.AutoMigrate(&models.Vehicle{}, &models.User{}, &models.Contact{})
	if err != nil {
		logger.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Tablas vehicles, users y contacts creadas exitosamente!")

	vehicleRepo := repositories.NewRepoVehicle(db, logger)
	userRepo := repositories.NewRepoUser(db, logger)
	contactRepo := repositories.NewContactRepository(db, logger)

	vehicleService := services.NewVehicleService(vehicleRepo)
	userService := services.NewUserService(userRepo)
	contactService := services.NewContactService(contactRepo)

	vehicleHandler := handlers.NewVehicleHandler(*vehicleService)
	userHandler := handlers.NewUserHandler(userService)
	contactHandler := handlers.NewContactHandler(contactService)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	cwd, _ := os.Getwd()
	staticPath := filepath.Join(cwd, "static") // tu carpeta con archivos estáticos
	router.Static("/static", staticPath)       // para acceder a CSS/JS desde HTML

	// Servir imágenes desde /images
	router.Static("/images", filepath.Join(staticPath, "images"))

	// Ruta principal - redirige a aboutUs.html
	router.GET("/", func(c *gin.Context) {
		c.File(filepath.Join(staticPath, "aboutUs.html"))
	})

	// Rutas específicas para los archivos HTML
	router.GET("/login", func(c *gin.Context) {
		c.File(filepath.Join(staticPath, "login.html"))
	})

	router.GET("/home", func(c *gin.Context) {
		c.File(filepath.Join(staticPath, "Home.html"))
	})

	router.GET("/contact", func(c *gin.Context) {
		c.File(filepath.Join(staticPath, "contact.html"))
	})

	// Routes - vehicles.
	router.POST("/api/vehicles", vehicleHandler.CreateVehicle)
	router.GET("/api/vehicles/:id", vehicleHandler.GetVehicleById)
	router.PUT("/api/vehicles/:id", vehicleHandler.UpdateVehicle) //fixear
	router.GET("/api/vehicles/getAll", vehicleHandler.GetAllVehicles)
	router.POST("/api/vehicles/:id/upload-images", vehicleHandler.UploadVehicleImages)
	router.PUT("/api/vehicles/:id/update-images", vehicleHandler.UpdateVehicleImages)
	// Routes - users
	router.POST("/api/users", userHandler.CreateUser)
	router.GET("/api/users/:id", userHandler.GetUser)
	router.PUT("/api/users/:id", userHandler.UpdateUser)
	router.POST("/api/forgot-password", userHandler.ForgotPassword)
	router.POST("/api/login", userHandler.LoginUser)
	router.DELETE("/api/users/:id", userHandler.DeleteUserById)

	// Routes - contacts
	router.POST("/api/contact", contactHandler.CreateContact)
	router.GET("/api/contacts", contactHandler.GetAllContacts)
	router.GET("/api/contacts/:id", contactHandler.GetContactByID)

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// create the server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{
		Logger:         logger,
		Validate:       validator.New(),
		UserService:    userService,
		UserRepo:       userRepo,
		UserHandler:    userHandler,
		VehicleHandler: vehicleHandler,
		VehicleModel:   &models.Vehicle{},
		VehicleRepo:    vehicleRepo,
		VehicleService: vehicleService,
		ContactHandler: contactHandler,
		ContactRepo:    contactRepo,
		ContactService: contactService,
		Server:         server,
	}
}

func (app *App) Start(port string) {
	app.Logger.Printf("Starting application on port %s", port)
	app.Server.Addr = ":" + port
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.Logger.Fatalf("Could not listen on %s: %v", port, err)
	}
}
