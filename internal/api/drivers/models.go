package drivers

import (
	"time"
)

type Driver struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name" binding:"required"`
	Age       string    `json:"age" binding:"required"`
	LicenseA  bool      `json:"license_a" binding:"required"`
	LicenseB  bool      `json:"license_b" binding:"required"`
	LicenseC  bool      `json:"license_c" binding:"required"`
	LicenseD  bool      `json:"license_d" binding:"required"`
	LicenseE  bool      `json:"license_e" binding:"required"`
}

type Tabler interface {
	TableName() string
}

func (Driver) TableName() string {
	return "drivers"
}
