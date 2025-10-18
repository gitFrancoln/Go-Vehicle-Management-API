package services

import (
	"errors"
	"regexp"
	"strings"
	"tpIRSO/models"
	"tpIRSO/repositories"
)

type VehicleService struct {
	vehicleRepo repositories.VehicleRepo
}

func NewVehicleService(vehicleRepo repositories.VehicleRepo) *VehicleService {
	return &VehicleService{
		vehicleRepo: vehicleRepo,
	}
}

func (vs *VehicleService) validateVehicle(vehicle *models.Vehicle) error {
	if strings.TrimSpace(vehicle.Brand) == "" {
		return errors.New("la marca no puede estar vacía")
	}
	if strings.TrimSpace(vehicle.Model) == "" {
		return errors.New("el modelo no puede estar vacío")
	}
	if strings.TrimSpace(vehicle.State) == "" {
		return errors.New("es necesario que especifica la condición del vehículo")
	}
	if strings.TrimSpace(vehicle.Titulo) == "" {
		return errors.New("la publicación debe contener un título")
	}
	if strings.TrimSpace(vehicle.Version) == "" {
		return errors.New("la publicación debe contener la versión del vehículo")
	}
	if vehicle.Price <= 0 {
		return errors.New("el precio debe ser mayor a 0")
	}
	if vehicle.Year < 1900 || vehicle.Year > 2025 {
		return errors.New("el año debe estar entre 1900 y 2025")
	}
	urlPattern := regexp.MustCompile(`^https?://.*\.(jpg|jpeg|png|gif|webp)$`)
	for _, imageURL := range vehicle.Image {
		if strings.TrimSpace(imageURL) == "" {
			return errors.New("las URLs de imágenes no pueden estar vacías")
		}
		if !urlPattern.MatchString(strings.ToLower(imageURL)) {
			return errors.New("las URLs de imágenes deben ser válidas y terminar en jpg, jpeg, png, gif o webp")
		}
	}
	if len(vehicle.Titulo) > 100 {
		return errors.New("el título no puede exceder 100 caracteres")
	}
	if len(vehicle.Brand) > 30 {
		return errors.New("la marca no puede exceder 30 caracteres")
	}
	if len(vehicle.Model) > 30 {
		return errors.New("el modelo no puede exceder 30 caracteres")
	}
	if len(vehicle.Version) > 30 {
		return errors.New("la versión no puede exceder 30 caracteres")
	}
	if len(vehicle.State) > 30 {
		return errors.New("el estado no puede exceder 30 caracteres")
	}

	return nil
}

func (v *VehicleService) CreateVehicle(vehicles *models.Vehicle) error {
	if err := v.validateVehicle(vehicles); err != nil {
		return err
	}

	existingVehicle, _ := v.vehicleRepo.Get(vehicles.ID.String())

	if existingVehicle != nil {
		return errors.New("ya existe un vehiculo con este id")
	}

	return v.vehicleRepo.Create(vehicles)
}

func (v *VehicleService) UpdateVehicle(vehicles *models.Vehicle) error {
	if err := v.validateVehicle(vehicles); err != nil {
		return err
	}

	existingVehicle, err := v.vehicleRepo.Get(vehicles.ID.String())

	if err != nil {
		return errors.New("vehiculo no encontrado")
	}

	vehicles.CreatedAt = existingVehicle.CreatedAt

	return v.vehicleRepo.Update(vehicles.ID.String(), vehicles)
}

func (v *VehicleService) GetVehicleById(vehicles *models.Vehicle) error {
	if vehicles.ID.String() == "" {
		return errors.New("el ID del vehículo no puede estar vacío")
	}

	existingVehicle, err := v.vehicleRepo.Get(vehicles.ID.String())
	if err != nil {
		return errors.New("vehículo no encontrado")
	}

	*vehicles = *existingVehicle
	return nil
}

func (v *VehicleService) GetAll(filters repositories.Filters, offset, limit int) ([]models.Vehicle, error) {
	vehicles, err := v.vehicleRepo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, errors.New("no se pudieron obtener los vehículos")
	}
	return vehicles, nil
}
