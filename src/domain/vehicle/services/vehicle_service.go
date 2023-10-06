package services

import (
	"context"
	"time"

	"github.com/DevShuxat/eater-service/src/infrastructure/rand"
	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/models"
	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/repositories"
	"go.uber.org/zap"
)

type VehicleService interface {
	CreateVehicle(ctx context.Context, driverID, model, make string) (*models.Vehicle, error)
	UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error
	DeleteVehicle(ctx context.Context, vehicleID string) error
	GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error)
	ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error)
}

type vehicleSvcImpl struct {
	vehicleRepo repositories.VehicleRepository
	logger      *zap.Logger
}

func NewVehicleService(
	vehicleRepo repositories.VehicleRepository,
	logger *zap.Logger,
) VehicleService {
	return &vehicleSvcImpl{
		vehicleRepo: vehicleRepo,
		logger:      logger,
	}
}

func (s *vehicleSvcImpl) CreateVehicle(ctx context.Context, driverID, model, make string) (*models.Vehicle, error) {
	now := time.Now()
	vehicle := models.Vehicle{
		ID:          rand.UUID(),
		DriverID:    driverID,
		Model:       model,
		Make:        make,
		PlateNumber: "",
		ImageUrl:    "",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.vehicleRepo.SaveVehicle(ctx, &vehicle); err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *vehicleSvcImpl) UpdateVehicle(ctx context.Context, vehicle *models.Vehicle) error {
	vehicle.Updated()
	if err := s.vehicleRepo.UpdateVehicle(ctx,vehicle); err != nil {
		return err
	}
	return nil
}

func (s *vehicleSvcImpl) DeleteVehicle(ctx context.Context, vehicleID string) error {
	if err := s.vehicleRepo.DeleteVehicle(ctx, vehicleID); err != nil {
		return err
	}
	return nil
}

func (s *vehicleSvcImpl) GetVehicle(ctx context.Context, vehicleID string) (*models.Vehicle, error) {
	vehicle, err := s.vehicleRepo.GetVehicle(ctx, vehicleID)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}

func (s *vehicleSvcImpl) ListVehicleByDriver(ctx context.Context, driverID string) ([]*models.Vehicle, error) {
	vehicles, err := s.vehicleRepo.ListVehicleByDriver(ctx, driverID)
	if err != nil {
		return nil, err
	}
	return vehicles, nil

}
