package repositories

import (
	"errors"
	"log"
	"tpIRSO/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *models.User) error
		Update(user *models.User) error
		GetByID(id string) (*models.User, error)
		GetByEmail(email string) (*models.User, error)
		SaveOrUpdate(user *models.User) error
	}

	userRepo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepoUser(db *gorm.DB, l *log.Logger) UserRepository {
	return &userRepo{
		db:  db,
		log: l,
	}
}

func (ur *userRepo) Create(user *models.User) error {

	if user.ID != uuid.Nil {
		ur.log.Println("Error: User already has ID, use Update instead")
		return errors.New("User already has ID, cannot create")
	}

	if err := ur.db.Create(user).Error; err != nil {
		ur.log.Printf("Error creating user: %v", err)
		return err
	}
	ur.log.Println("User created successfully")
	return nil
}

func (ur *userRepo) Update(user *models.User) error {
	if err := ur.db.Save(user).Error; err != nil {
		ur.log.Printf("Error updating user: %v", err)
		return err
	}
	ur.log.Println("User updated successfully")
	return nil
}

func (ur *userRepo) GetByID(id string) (*models.User, error) {
	var userId models.User
	if err := ur.db.Where("id = ?", id).First(&userId).Error; err != nil {
		ur.log.Printf("Error getting user by id: %v", err)
		return nil, err
	}
	return &userId, nil
}

func (ur *userRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		ur.log.Printf("Error getting user by email: %v", err)
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) SaveOrUpdate(user *models.User) error {
	if user.ID == uuid.Nil {
		ur.log.Println("ID is empty, creating new user")
		return ur.Create(user)
	}
	var existingUser models.User
	err := ur.db.Where("id = ?", user.ID).First(&existingUser).Error

	if err == gorm.ErrRecordNotFound {
		ur.log.Println("User with ID not found in DB, creating user")
		return ur.Create(user)
	} else if err != nil {
		ur.log.Println("Error checking vehicle existence", err)
		return err
	}
	ur.log.Println("User already exist, updating")
	return ur.Update(user)
}
