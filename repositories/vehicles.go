package repositories

import (
	"errors"
	"log"
	"tpIRSO/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// definimos los filtros de búsqueda - por que se busca? --> por marca, modelo, estado, etc.
type Filters struct {
	Brand   string
	Model   string
	MinYear int
	MaxYear int
	State   string
}

type (
	VehicleRepo interface {
		Create(vehicles *models.Vehicle) error                               //definimos la interfaz del repo, definimos al puntero de vehicle llamado vehicles y manejamos los errores.
		GetAll(filters Filters, offset, limit int) ([]models.Vehicle, error) //el getAll contiene los Filtros de búsqueda, para buscar un vehiculo en particular, offset define dónde empieza la paginación. Limit define el límite de cuántos vehiculos traer. El []models.vehicle es el resultado(lista de vehiculos)
		Get(id string) (*models.Vehicle, error)
		Delete(id string) error
		Update(id string, vehicle *models.Vehicle) error
		SaveOrUpdate(vehicle *models.Vehicle) error //verificamos si estamos guardando o actualizando
	}
	//definimos como debe ser el repositorio
	vehicleRepo struct {
		db  *gorm.DB    //conexion a db
		log *log.Logger //registro de eventos
	}
)

// constructor
func NewRepoVehicle(db *gorm.DB, l *log.Logger) VehicleRepo { //constructor, recibe las dependencias necesarias, retorna la interfaz Repository, no la struct
	return &vehicleRepo{ //retorna un puntero a la struct
		db:  db, //asigna la conexión db
		log: l,  //asigna el logger
	}
}

func (r *vehicleRepo) Create(vehicle *models.Vehicle) error {
	// Verificar que es una creación (ID debe estar vacío)
	if vehicle.ID != uuid.Nil {
		r.log.Println("Error: Vehicle already has ID, use Update instead")
		return errors.New("vehicle already has ID, cannot create")
	}

	// El UUID se asigna automáticamente en BeforeCreate del modelo
	if err := r.db.Create(vehicle).Error; err != nil {
		r.log.Println("Error creating vehicle:", err)
		return err
	}
	r.log.Println("Vehicle created with id:", vehicle.ID)
	return nil
}

func (r *vehicleRepo) GetAll(filters Filters, offset, limit int) ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	query := r.db.Model(&models.Vehicle{})

	// Aplicar filtros
	if filters.Brand != "" {
		query = query.Where("brand = ?", filters.Brand)
	}
	if filters.Model != "" {
		query = query.Where("model = ?", filters.Model)
	}
	if filters.MinYear > 0 {
		query = query.Where("year >= ?", filters.MinYear)
	}
	if filters.MaxYear > 0 {
		query = query.Where("year <= ?", filters.MaxYear)
	}
	if filters.State != "" {
		query = query.Where("state = ?", filters.State)
	}

	// Aplicar paginación
	if err := query.Offset(offset).Limit(limit).Find(&vehicles).Error; err != nil {
		r.log.Println("Error getting vehicles:", err)
		return nil, err
	}

	return vehicles, nil
}

func (r *vehicleRepo) Get(id string) (*models.Vehicle, error) {
	var vehicle models.Vehicle
	if err := r.db.Where("id = ?", id).First(&vehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Println("Vehicle not found with id:", id)
			return nil, err
		}
		r.log.Println("Error getting vehicle:", err)
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepo) Delete(id string) error {
	result := r.db.Where("id = ?", id).Delete(&models.Vehicle{})
	if result.Error != nil {
		r.log.Println("Error deleting vehicle:", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		r.log.Println("Vehicle not found for deletion with id:", id)
		return gorm.ErrRecordNotFound
	}
	r.log.Println("Vehicle deleted with id:", id)
	return nil
}

func (r *vehicleRepo) Update(id string, vehicle *models.Vehicle) error {
	// Verificar que el vehículo existe
	var existingVehicle models.Vehicle
	if err := r.db.Where("id = ?", id).First(&existingVehicle).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Println("Vehicle not found for update with id:", id)
			return err
		}
		r.log.Println("Error finding vehicle for update:", err)
		return err
	}

	if err := r.db.Model(&existingVehicle).Where("id = ?", id).Updates(vehicle).Error; err != nil {
		r.log.Println("Error updating vehicle:", err)
		return err
	}

	r.log.Println("Vehicle updated with id:", id)
	return nil
}

func (r *vehicleRepo) SaveOrUpdate(vehicle *models.Vehicle) error {
	// Si el ID está vacío (uuid.Nil), es una creación
	if vehicle.ID == uuid.Nil {
		r.log.Println("ID is empty, creating new vehicle")
		return r.Create(vehicle)
	}

	// Si el ID existe, verificar si el registro existe en la DB
	var existingVehicle models.Vehicle
	err := r.db.Where("id = ?", vehicle.ID).First(&existingVehicle).Error

	if err == gorm.ErrRecordNotFound {
		// El ID existe en el struct pero no en la DB - crear nuevo
		r.log.Println("Vehicle with ID not found in DB, creating new vehicle")
		return r.Create(vehicle)
	} else if err != nil {
		// Error de conexión u otro
		r.log.Println("Error checking vehicle existence:", err)
		return err
	}

	// El registro existe - actualizar
	r.log.Println("Vehicle exists, updating")
	return r.Update(vehicle.ID.String(), vehicle)
}
