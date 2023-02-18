package drivers

import (
	Shared "orc-api/internal/shared"
)

type Driver struct {
	Shared.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	LicenseA bool   `json:"license_a" gorm:"default:false"`
	LicenseB bool   `json:"license_b" gorm:"default:false"`
	LicenseC bool   `json:"license_c" gorm:"default:false"`
	LicenseD bool   `json:"license_d" gorm:"default:false"`
	LicenseE bool   `json:"license_e" gorm:"default:false"`
}

type Tabler interface {
	TableName() string
}

func (Driver) TableName() string {
	return "drivers"
}
