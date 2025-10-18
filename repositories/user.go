package repositories

import (
	"errors"
	"fmt"
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
		DeleteUserById(userID string) error
		ForgotPasswordByID(user *models.User) error
		UpdatePassword(email string, hashedPassword string) error
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
		return errors.New("user already has ID, cannot create")
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

func (ur *userRepo) DeleteUserById(userID string) error {
	// Iniciar una transacción para asegurar la atomicidad
	tx := ur.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error al iniciar la transacción: %w", tx.Error)
	}

	// Verificar si el usuario existe antes de intentar eliminarlo
	var user models.User
	if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
		tx.Rollback() // Revertir la transacción
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuario no encontrado")
		}
		return fmt.Errorf("error al buscar el usuario: %w", err)
	}

	// Eliminar el usuario
	if err := tx.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		tx.Rollback() // Revertir en caso de error en la eliminación
		return fmt.Errorf("error al eliminar el usuario: %w", err)
	}

	// Si todo fue bien, confirmar la transacción
	return tx.Commit().Error
}

func (ur *userRepo) ForgotPasswordByID(user *models.User) error {
	if err := ur.db.Model(user).Update("password", user.Password).Error; err != nil {
		ur.log.Printf("error updating password : %v", err)
		return err
	}
	ur.log.Println("password updated succesfully")
	return nil
}

func (ur *userRepo) UpdatePassword(email string, hashedPassword string) error {
	// Iniciar una transacción
	tx := ur.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error al iniciar la transacción: %w", tx.Error)
	}

	// Verificar que el usuario existe y obtenerlo
	var user models.User
	if err := tx.Where("email = ?", email).First(&user).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("usuario no encontrado")
		}
		return fmt.Errorf("error al buscar el usuario: %w", err)
	}

	// Actualizar solo el campo de contraseña
	if err := tx.Model(&user).Update("password", hashedPassword).Error; err != nil {
		tx.Rollback()
		ur.log.Printf("Error actualizando contraseña: %v", err)
		return fmt.Errorf("error al actualizar la contraseña: %w", err)
	}

	// Confirmar la transacción
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error al confirmar la transacción: %w", err)
	}

	ur.log.Printf("Contraseña actualizada exitosamente para el usuario con email: %s", email)
	return nil
}
