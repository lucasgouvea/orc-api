package companies

import (
	Shared "orc-api/internal/shared"
)

type Company struct {
	Shared.Model
	Name            string
	Type            int
	IntermediatorID *uint
	Intermediateds  []Company `gorm:"foreignkey:IntermediatorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
	case int(INTERMEDIATED):
		{
			return CONTRACT
		}
	}
	panic(InvalidCompanyTypeErr)
}

func (c Company) toSchema() *CompanySchema {
	var intermediateds []CompanyIntermediatedSchema = make([]CompanyIntermediatedSchema, 0)
	base := CompanyBaseSchema{
		ID:        int(c.ID),
		CreatedAt: c.CreatedAt,
		Name:      c.Name,
		Type:      c.getType().String(),
	}

	for _, i := range c.Intermediateds {
		intermediateds = append(intermediateds, CompanyIntermediatedSchema{
			CompanyBaseSchema{
				ID:        int(i.ID),
				CreatedAt: i.CreatedAt,
				Name:      i.Name,
				Type:      i.getType().String(),
			},
		})
	}

	if c.Type == int(AGGREGATE) {
		return &CompanySchema{
			base,
			intermediateds,
		}
	}

	return &CompanySchema{
		base,
		intermediateds,
	}

}
