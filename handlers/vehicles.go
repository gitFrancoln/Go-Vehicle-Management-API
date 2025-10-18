package handlers

import (
	"net/http"
	"strconv"
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
