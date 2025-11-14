package handlers

import (
	"net/http"
	"tpIRSO/models"
	"tpIRSO/services"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	contactService *services.ContactService
}

func NewContactHandler(contactService *services.ContactService) *ContactHandler {
	return &ContactHandler{
		contactService: contactService,
	}
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var contact models.Contact

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al guardar el mensaje de contacto",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Mensaje enviado correctamente. ¡Gracias por contactarnos!",
		"contact": contact,
	})
}

func (h *ContactHandler) GetAllContacts(c *gin.Context) {
	contacts, err := h.contactService.GetAllContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener los mensajes de contacto",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}

func (h *ContactHandler) GetContactByID(c *gin.Context) {
	id := c.Param("id")

	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Mensaje de contacto no encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contact": contact,
	})
}
