package handlers

import (
	"encoding/json"
	"net/http"
	"tpIRSO/models"
	"tpIRSO/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(vehicleService services.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		vehicleService: vehicleService,
	}
}

func (vh *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle models.Vehicle

	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Error al decodificar el vehiculo. Asegúrate de que el JSON cumpla con el formato requerido: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := vh.vehicleService.CreateVehicle(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

func (vh *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID := vars["id"]

	// Validar que el ID sea un UUID válido
	_, err := uuid.Parse(vehicleID)
	if err != nil {
		http.Error(w, "ID de vehículo inválido", http.StatusBadRequest)
		return
	}

	var vehicle models.Vehicle
	if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
		http.Error(w, "Error al decodificar los datos del vehículo: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Asignar el ID del path al vehículo
	vehicle.ID, _ = uuid.Parse(vehicleID)

	if err := vh.vehicleService.UpdateVehicle(&vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}

func (vh *VehicleHandler) GetVehicleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID := vars["id"]

	// Crear una instancia de Vehicle con el ID
	vehicle := &models.Vehicle{}
	vehicle.ID.UnmarshalText([]byte(vehicleID))

	// Llamar al servicio con el objeto vehicle
	if err := vh.vehicleService.GetVehicleById(vehicle); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}
