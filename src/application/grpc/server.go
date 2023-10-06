package grpc

import (
	"context"
	"errors"
	"github.com/DevShuxat/logistic-vehicle-service/src/application/dtos"
	pb "github.com/DevShuxat/logistic-vehicle-service/src/application/protos/logistic_vehicle"
	vehiclesvc "github.com/DevShuxat/logistic-vehicle-service/src/domain/vehicle/services"
)

type Server struct {
	pb.UnimplementedVehicleServiceServer
	vehicleSvc vehiclesvc.VehicleService
}

func NewServer(vehicleSvc vehiclesvc.VehicleService) Server {
	return Server{
		vehicleSvc: vehicleSvc,
	}
}

func (s *Server) CreateVehicle(c context.Context, r *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	if r.GetDriverID() == "" {
		return nil, errors.New("invalid or missing driver_id")
	}

	vehicle, err := s.vehicleSvc.CreateVehicle(c, r.GetDriverID(), r.GetModel(), r.GetMake())
	if err != nil {
		return nil, err
	}
	return &pb.CreateVehicleResponse{
		Vehicle: dtos.ToVehicle(vehicle),
	}, nil
}

func (s *Server) UpdateVehicle(c context.Context, r *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	vehicle := dtos.ToVehicle(r.GetVehicle())
	if err := s.vehicleSvc.UpdateVehicle(c, vehicle); err != nil {
		return nil, err
	}
	return &pb.UpdateVehicleResponse{}, nil
}

func (s *Server) DeleteVehicle(c context.Context, r *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	if r.GetVehicleId() == "" {
		return nil, errors.New("invalid or missing vehicle_id")
	}

	return *pb.DeleteVehicleResponse{}, nil
}

func (s *Server) GetVehicle(c context.Context, r *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	if r.GetVehicleId() == "" {
		return nil, errors.New("invalid or missing vehicle_id")
	}

	vehicle, err := s.vehicleSvc.GetVehicle(c, r.GetVehicleId())
	if err != nil {
		return nil, err
	}

	return &pb.GetVehicleResponse{
		Vehicle: dtos.ToVehiclePB(vehicle),
	}, nil
}

func (s *Server) ListVehicleByDriver(c context.Context, r *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error) {
	if r.GetDriverId() == "" {
		return nil, errors.New("invalid or missing driver_id")
	}

	vehicles, err := s.vehicleSvc.ListVehicleByDriver(c, r.GetDriverId())
	if err != nil {
		return nil, err
	}

	return *pb.ListVehicleResponse{
		Vehicles: dtos.ToVehiclesPB(vehicles),
	}, nil
}
