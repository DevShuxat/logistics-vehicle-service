package dtos

import (
	"time"
	pb "github.com/DevShuxat/logistic-vehicle-service/src/application/protos/logistic_vehicle"
	"github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/models"
)

func ToVehiclePB(vehicle *models.Vehicle) *pb.Vehicle {
	return &pb.Vehicle{
		Id:          vehicle.ID,
		DriverID:    vehicle.DriverID,
		Model:       vehicle.Model,
		Make:        vehicle.Make,
		PlateNumber: vehicle.PlateNumber,
		Image:       vehicle.ImageUrl,
		CreatedAt:   vehicle.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   vehicle.UpdatedAt.Format(time.RFC3339),
	}
}
func ToVehicle(vehicle *pb.Vehicle) *models.Vehicle {
	return &models.Vehicle{
		ID:          vehicle.GetId(),
		DriverID:    vehicle.GetDriverID(),
		Model:       vehicle.GetModel(),
		Make:        vehicle.GetMake(),
		PlateNumber: vehicle.GetPlateNumber(),
		ImageUrl:    vehicle.GetImage(),
		CreatedAt:   toTime(vehicle.GetCreateAt()),
		UpdatedAt:   toTime(vehicle.GetUpdateAt()),
	}
}
func ToVehiclesPB(vehicles []*models.Vehicle) []*pb.Vehicle {
	result := make([]*pb.Vehicle, len(vehicles))
	for i := range vehicles {
		result[i] = ToVehiclePB(vehicles[i])
	}
	return result
}
