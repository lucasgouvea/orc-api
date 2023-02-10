package drivers

import (
	"time"
)

type Driver struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	LicenseA  *bool     `json:"license_a" gorm:"default:false"`
	LicenseB  *bool     `json:"license_b" gorm:"default:false"`
	LicenseC  *bool     `json:"license_c" gorm:"default:false"`
	LicenseD  *bool     `json:"license_d" gorm:"default:false"`
	LicenseE  *bool     `json:"license_e" gorm:"default:false"`
}

type Tabler interface {
	TableName() string
}

func (Driver) TableName() string {
	return "drivers"
}
