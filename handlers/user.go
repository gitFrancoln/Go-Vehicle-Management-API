package handlers

import (
	"net/http"
	"strings"
	"tpIRSO/dto"
	"tpIRSO/models"
	"tpIRSO/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	// Decodificar el JSON del request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del usuario"})
		return
	}

	if err := uh.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos del usuario"})
		return
	}

	existingUser, err := uh.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	user.ID = existingUser.ID

	if err := uh.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	// Obtener el usuario por ID
	user, err := uh.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro email es requerido"})
		return
	}

	user, err := uh.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	err := uh.userService.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}

func (uh *UserHandler) ForgotPassword(c *gin.Context) {
	var request dto.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos. Se requiere email y nueva contraseña"})
		return
	}

	// Normalizar aliases: permitir user_email/user_password
	email := strings.TrimSpace(request.Email)
	if email == "" {
		email = strings.TrimSpace(request.UserEmail)
	}
	newPass := strings.TrimSpace(request.NewPassword)
	if newPass == "" {
		newPass = strings.TrimSpace(request.UserPassword)
	}

	if email == "" || newPass == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email y nueva contraseña son requeridos"})
		return
	}

	// Intenta restablecer la contraseña
	if err := uh.userService.ForgotPassword(email, newPass); err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contraseña restablecida correctamente"})
}

func (uh *UserHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido o campos faltantes"})
		return
	}

	user, err := uh.userService.LoginUser(req.UserEmail, req.UserPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	user.Password = "" // no devolver contraseña

	c.JSON(http.StatusOK, gin.H{"message": "Login exitoso", "user": user})
}
