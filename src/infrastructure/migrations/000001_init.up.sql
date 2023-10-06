CREATE SCHEMA IF NOT EXISTS logistcs_vehicle

CREATE TABLE IF NOT EXISTS logistcs_vehicle.vehicles (
	id varchar(36) PRIMARY KEY,
	driver_id varchar(36) NOT NULL
	model varchar(36),
	make varchar(36),
	plate_number varchar(36),
	image_url text,
	created_at timestamp,
	updated_at timestamp
);

CREATE INDEX IF NOT EXISTS idx_driver_vehicle ON logistcs_vehicle.vehicles (driver_id);