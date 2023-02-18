package vehicles

import (
	Shared "orc-api/internal/shared"
	"time"
)

type Vehicle struct {
	Shared.Model
	CreatedAt        time.Time `json:"created_at" `
	ModelDescription string    `json:"model"`
	LicensePlate     string    `json:"license_plate"`
}

type Tabler interface {
	TableName() string
}

func (Vehicle) TableName() string {
	return "vehicles"
}
