package vehicle

import (
	"context"

	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/models"
	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/repositories"
	"gorm.io/gorm"
)

const (
	tableVehicle = "logistics_vehicle.vehicles"
)

type vehicleRepoImpl struct {
	db *gorm.DB
}

func NewVehicleService(db *gorm.DB) repositories.VehicleRepository {
	return &vehicleRepoImpl{
		db: db,
	}
}

func (r *vehicleRepoImpl) SaveVehicle(ctx context.Context, vehicle *models.Vehicle) error {
	result := r.db.WithContext(ctx).Table(tableVehicle).Create(vehicle)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *vehicleRepoImpl) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error {
	result := r.db.WithContext(ctx).Table(tableVehicle).Save(vehicle)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *vehicleRepoImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {
	result := r.db.WithContext(ctx).Table(tableVehicle).Delete(&models.Vehicle{}, "id = ?", vehicleID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *vehicleRepoImpl) GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error) {
	var vehicle *models.Vehicle
	result := r.db.WithContext(ctx).Table(tableVehicle).First(&vehicle, "id = ?", vehicleID)
	if result.Error != nil {
		return nil, result.Error
	}
	return vehicle, nil
}

func (r *vehicleRepoImpl) ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle
	result := r.db.WithContext(ctx).Table(tableVehicle).Where("driver_id = ?", driverID).Find(&vehicles)
	if result.Error != nil {
		return nil, result.Error
	}
	return vehicles, nil
}
