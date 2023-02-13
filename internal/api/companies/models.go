package companies

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Company struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name"`
	Type      int       `json:"type"`
}

type Tabler interface {
	TableName() string
}

func (Company) TableName() string {
	return "companies"
}

func (c *companyType) Scan(value any) error {
	_, ok := value.([]byte)
	if !ok {
		return errors.New("Failed scan for companyType")
	}
	return nil
}

func (c *companyType) Value() (driver.Value, error) {
	return 1, nil
}
