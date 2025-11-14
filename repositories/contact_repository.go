package repositories

import (
	"log"
	"tpIRSO/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	CreateContact(contact *models.Contact) error
	GetAllContacts() ([]models.Contact, error)
	GetContactByID(id string) (*models.Contact, error)
	DeleteContact(id string) error
}

type contactRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewContactRepository(db *gorm.DB, logger *log.Logger) ContactRepository {
	return &contactRepository{
		db:     db,
		logger: logger,
	}
}

func (r *contactRepository) CreateContact(contact *models.Contact) error {
	if err := r.db.Create(contact).Error; err != nil {
		r.logger.Printf("Error creating contact: %v", err)
		return err
	}
	r.logger.Printf("Contact created successfully with ID: %d", contact.ID)
	return nil
}

func (r *contactRepository) GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	if err := r.db.Find(&contacts).Error; err != nil {
		r.logger.Printf("Error fetching all contacts: %v", err)
		return nil, err
	}
	r.logger.Printf("Fetched %d contacts", len(contacts))
	return contacts, nil
}

func (r *contactRepository) GetContactByID(id string) (*models.Contact, error) {
	var contact models.Contact
	if err := r.db.First(&contact, id).Error; err != nil {
		r.logger.Printf("Error fetching contact with ID %s: %v", id, err)
		return nil, err
	}
	r.logger.Printf("Fetched contact with ID: %s", id)
	return &contact, nil
}

func (r *contactRepository) DeleteContact(id string) error {
	if err := r.db.Delete(&models.Contact{}, id).Error; err != nil {
		r.logger.Printf("Error deleting contact with ID %s: %v", id, err)
		return err
	}
	r.logger.Printf("Deleted contact with ID: %s", id)
	return nil
}
