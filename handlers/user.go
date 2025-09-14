package handlers

import (
	"encoding/json"
	"net/http"
	"tpIRSO/models"
	"tpIRSO/services"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser maneja la creación de un nuevo usuario
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decodificar el JSON del request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error al decodificar los datos del usuario", http.StatusBadRequest)
		return
	}

	// Crear el usuario usando el servicio (incluye validaciones)
	if err := uh.userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Responder con el usuario creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// UpdateUser maneja la actualización de un usuario existente
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.User

	// Decodificar el JSON del request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error al decodificar los datos del usuario", http.StatusBadRequest)
		return
	}

	// Obtener el usuario existente para mantener el ID
	existingUser, err := uh.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Mantener el ID del usuario existente
	user.ID = existingUser.ID

	// Actualizar el usuario usando el servicio (incluye validaciones)
	if err := uh.userService.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Responder con el usuario actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUser maneja la obtención de un usuario por ID
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	// Obtener el usuario por ID
	user, err := uh.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con el usuario encontrado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUserByEmail maneja la obtención de un usuario por email
func (uh *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "El parámetro email es requerido", http.StatusBadRequest)
		return
	}

	// Obtener el usuario por email
	user, err := uh.userService.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Responder con el usuario encontrado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
