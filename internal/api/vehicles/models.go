package vehicles

import (
	Shared "orc-api/internal/shared"
	"strconv"
	"time"
)

type Vehicle struct {
	Shared.Model
	CreatedAt        time.Time `json:"created_at" `
	ModelDescription string    `json:"model"`
	LicensePlate     string    `json:"license_plate"`
}

func (Vehicle) TableName() string {
	return "vehicles"
}

func (v Vehicle) Schema() VehicleSchema {
	return VehicleSchema{
		Id:               strconv.FormatUint(uint64(v.ID), 10),
		CreatedAt:        v.CreatedAt.String(),
		UpdatedAt:        v.UpdatedAt.String(),
		ModelDescription: v.ModelDescription,
		LicensePlate:     v.LicensePlate,
	}
}
