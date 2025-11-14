package services

import (
	"tpIRSO/models"
	"tpIRSO/repositories"
)

type ContactService struct {
	contactRepo repositories.ContactRepository
}

func NewContactService(contactRepo repositories.ContactRepository) *ContactService {
	return &ContactService{
		contactRepo: contactRepo,
	}
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	return s.contactRepo.CreateContact(contact)
}

func (s *ContactService) GetAllContacts() ([]models.Contact, error) {
	return s.contactRepo.GetAllContacts()
}

func (s *ContactService) GetContactByID(id string) (*models.Contact, error) {
	return s.contactRepo.GetContactByID(id)
}

func (s *ContactService) DeleteContact(id string) error {
	return s.contactRepo.DeleteContact(id)
}
