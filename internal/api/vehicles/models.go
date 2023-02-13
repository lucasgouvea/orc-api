package vehicles

import (
	"time"
)

type Vehicle struct {
	ID           int       `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time `json:"created_at" `
	Model        string    `json:"model"`
	LicensePlate string    `json:"license_plate"`
}

type Tabler interface {
	TableName() string
}

func (Vehicle) TableName() string {
	return "vehicles"
}
