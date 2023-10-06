package repositories

import (
	"context"

	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/models"
)

type VehicleRepository interface {
	SaveVehicle(ctx context.Context, vehicle *models.Vehicle) error
	UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error
	DeleteVehicle(ctx context.Context, vehicleID string) error
	GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error)
	ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error)
}

