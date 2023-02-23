package drivers

import (
	Shared "orc-api/internal/shared"
	"strconv"
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

func (Driver) TableName() string {
	return "drivers"
}

func (d Driver) Schema() DriverSchema {
	return DriverSchema{
		Id:        strconv.FormatUint(uint64(d.ID), 10),
		CreatedAt: d.CreatedAt.String(),
		UpdatedAt: d.CreatedAt.String(),
		Name:      d.Name,
		Age:       d.Age,
		LicenseA:  d.LicenseA,
		LicenseB:  d.LicenseB,
		LicenseC:  d.LicenseC,
		LicenseD:  d.LicenseD,
		LicenseE:  d.LicenseE,
	}
}
