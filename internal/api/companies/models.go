package companies

import (
	"time"
)

type Company struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	Name      string
	Type      int
}

type Tabler interface {
	TableName() string
}

func (Company) TableName() string {
	return "companies"
}

func (c Company) getType() companyType {
	switch c.Type {
	case int(AGGREGATE):
		{
			return AGGREGATE
		}
	case int(CONTRACT):
		{
			return CONTRACT
		}
	}
	panic(InvalidCompanyTypeErr)
}

func (c Company) toSchema() CompanySchema {
	return CompanySchema{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		Name:      c.Name,
		Type:      c.getType().String(),
	}
}
