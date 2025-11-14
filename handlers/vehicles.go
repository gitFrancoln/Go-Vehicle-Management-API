package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"tpIRSO/models"
	"tpIRSO/repositories"
	"tpIRSO/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(vehicleService services.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		vehicleService: vehicleService,
	}
}

func (vh *VehicleHandler) CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el vehiculo. Asegúrate de que el JSON cumpla con el formato requerido: " + err.Error()})
		return
	}

	if err := vh.vehicleService.CreateVehicle(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, vehicle)
}

func (vh *VehicleHandler) UpdateVehicle(c *gin.Context) {
	vehicleID := c.Param("id")

	// Validar que el ID sea un UUID válido
	_, err := uuid.Parse(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vehículo inválido"})
		return
	}

	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del vehículo: " + err.Error()})
		return
	}

	// Asignar el ID del path al vehículo
	vehicle.ID, _ = uuid.Parse(vehicleID)

	if err := vh.vehicleService.UpdateVehicle(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (vh *VehicleHandler) GetVehicleById(c *gin.Context) {
	vehicleID := c.Param("id")

	// Crear una instancia de Vehicle con el ID
	vehicle := &models.Vehicle{}
	vehicle.ID.UnmarshalText([]byte(vehicleID))

	// Llamar al servicio con el objeto vehicle
	if err := vh.vehicleService.GetVehicleById(vehicle); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vehicle)
}

func (vh *VehicleHandler) GetAllVehicles(c *gin.Context) {
	var filters repositories.Filters

	filters.Brand = c.Query("brand")
	filters.Model = c.Query("model")
	filters.State = c.Query("state")

	if minYearStr := c.Query("minYear"); minYearStr != "" {
		if minYear, err := strconv.Atoi(minYearStr); err == nil {
			filters.MinYear = minYear
		}
	}
	if maxYearStr := c.Query("maxYear"); maxYearStr != "" {
		if maxYear, err := strconv.Atoi(maxYearStr); err == nil {
			filters.MaxYear = maxYear
		}
	}

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	offset := (page - 1) * limit

	vehicles, err := vh.vehicleService.GetAll(filters, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si no hay vehículos, podemos devolver un array vacío
	if len(vehicles) == 0 {
		c.JSON(http.StatusOK, []models.Vehicle{})
		return
	}

	c.JSON(http.StatusOK, vehicles)
}

// UploadVehicleImages maneja la subida de imágenes para un vehículo
func (vh *VehicleHandler) UploadVehicleImages(c *gin.Context) {
	vehicleID := c.Param("id")

	// Validar que el ID sea un UUID válido
	parsedID, err := uuid.Parse(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vehículo inválido"})
		return
	}

	// Verificar que el vehículo existe
	vehicle := &models.Vehicle{}
	vehicle.ID = parsedID
	if err := vh.vehicleService.GetVehicleById(vehicle); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehículo no encontrado"})
		return
	}

	// Obtener el formulario multipart
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el formulario: " + err.Error()})
		return
	}

	files := form.File["images"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se han subido imágenes"})
		return
	}

	// Crear el directorio de imágenes si no existe
	cwd, _ := os.Getwd()
	uploadDir := filepath.Join(cwd, "static", "images")
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear directorio de imágenes"})
		return
	}

	var imageURLs []string

	// Procesar cada archivo
	for _, file := range files {
		// Validar tipo de archivo
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de imagen no válido. Solo se permiten: jpg, jpeg, png, gif, webp"})
			return
		}

		// Generar nombre único para el archivo
		timestamp := time.Now().Unix()
		newFilename := fmt.Sprintf("%s_%d%s", vehicleID, timestamp, ext)
		savePath := filepath.Join(uploadDir, newFilename)

		// Abrir el archivo subido
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir archivo subido"})
			return
		}
		defer src.Close()

		// Crear el archivo en el servidor
		dst, err := os.Create(savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar archivo"})
			return
		}
		defer dst.Close()

		// Copiar el contenido
		if _, err = io.Copy(dst, src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al copiar archivo"})
			return
		}

		// Agregar la URL al array
		imageURL := "/images/" + newFilename
		imageURLs = append(imageURLs, imageURL)
	}

	// Actualizar el vehículo con las nuevas imágenes
	vehicle.Image = append(vehicle.Image, imageURLs...)
	if err := vh.vehicleService.UpdateVehicle(vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar vehículo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Imágenes subidas exitosamente",
		"vehicle_id": vehicleID,
		"images":     imageURLs,
		"total":      len(imageURLs),
	})
}

// UpdateVehicleImages reemplaza todas las imágenes de un vehículo
func (vh *VehicleHandler) UpdateVehicleImages(c *gin.Context) {
	vehicleID := c.Param("id")

	// Validar que el ID sea un UUID válido
	parsedID, err := uuid.Parse(vehicleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de vehículo inválido"})
		return
	}

	// Verificar que el vehículo existe
	vehicle := &models.Vehicle{}
	vehicle.ID = parsedID
	if err := vh.vehicleService.GetVehicleById(vehicle); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehículo no encontrado"})
		return
	}

	var requestBody struct {
		Images []string `json:"images"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar JSON: " + err.Error()})
		return
	}

	// Actualizar las imágenes
	vehicle.Image = requestBody.Images
	if err := vh.vehicleService.UpdateVehicle(vehicle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar vehículo: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Imágenes actualizadas exitosamente",
		"vehicle_id": vehicleID,
		"images":     vehicle.Image,
	})
}
