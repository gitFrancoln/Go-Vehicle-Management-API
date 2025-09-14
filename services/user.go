package services

import (
	"errors"
	"regexp"
	"strings"
	"tpIRSO/models"
	"tpIRSO/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// ValidateUser valida que los campos del usuario no estén vacíos y tengan formato correcto
func (us *UserService) ValidateUser(user *models.User) error {
	// Validar que el nombre no esté vacío
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("el nombre del usuario no puede estar vacío")
	}

	// Validar que el email no esté vacío y tenga formato válido
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("el email del usuario no puede estar vacío")
	}

	// Validar formato del email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("el formato del email no es válido")
	}

	// Validar que el teléfono no esté vacío
	if strings.TrimSpace(user.Phone) == "" {
		return errors.New("el teléfono del usuario no puede estar vacío")
	}

	// Validar formato del teléfono (solo números, 10 dígitos)
	phoneRegex := regexp.MustCompile(`^\d{10}$`)
	if !phoneRegex.MatchString(user.Phone) {
		return errors.New("el teléfono debe contener exactamente 10 dígitos")
	}

	// Validar longitudes máximas
	if len(user.Name) > 60 {
		return errors.New("el nombre no puede exceder 60 caracteres")
	}

	if len(user.Email) > 120 {
		return errors.New("el email no puede exceder 120 caracteres")
	}

	return nil
}

// CreateUser crea un nuevo usuario con validaciones
func (us *UserService) CreateUser(user *models.User) error {
	// Validar los datos del usuario
	if err := us.ValidateUser(user); err != nil {
		return err
	}

	// Verificar que el email no esté ya registrado
	existingUser, _ := us.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return errors.New("ya existe un usuario con este email")
	}

	// Crear el usuario
	return us.userRepo.Create(user)
}

// UpdateUser actualiza un usuario existente con validaciones
func (us *UserService) UpdateUser(user *models.User) error {
	// Validar los datos del usuario
	if err := us.ValidateUser(user); err != nil {
		return err
	}

	// Verificar que el usuario existe
	existingUser, err := us.userRepo.GetByID(user.ID.String())
	if err != nil {
		return errors.New("usuario no encontrado")
	}

	// Verificar que el email no esté ya registrado por otro usuario
	userWithEmail, _ := us.userRepo.GetByEmail(user.Email)
	if userWithEmail != nil && userWithEmail.ID != user.ID {
		return errors.New("ya existe otro usuario con este email")
	}

	// Mantener las fechas de creación
	user.CreatedAt = existingUser.CreatedAt

	// Actualizar el usuario
	return us.userRepo.Update(user)
}

// GetUserByID obtiene un usuario por su ID
func (us *UserService) GetUserByID(id string) (*models.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("el ID del usuario no puede estar vacío")
	}

	return us.userRepo.GetByID(id)
}

// GetUserByEmail obtiene un usuario por su email
func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	if strings.TrimSpace(email) == "" {
		return nil, errors.New("el email no puede estar vacío")
	}

	return us.userRepo.GetByEmail(email)
}
