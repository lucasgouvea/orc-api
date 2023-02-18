package companies

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CompanyBaseSchema struct {
	ID        int       `json:"id" `
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name"`
	Type      string    `json:"type"`
}

type CompanySchema struct {
	CompanyBaseSchema
	Intermediateds []CompanyIntermediatedSchema `json:"intermediateds,omitempty"`
}

type CompanyIntermediatedSchema struct {
	CompanyBaseSchema
}

type CompanyPostSchema struct {
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required_company_type"`
	Intermediateds []struct {
		Name string `json:"name"`
	} `json:"intermediateds"`
}

type CompanyPatchSchema struct {
	Name *string `json:"name"`
	Type *string `json:"type" binding:"optional_company_type"`
}

func (c CompanyPostSchema) parse() (*Company, error) {
	cType := NewCompanyType(c.Type)

	company := Company{}
	company.Name = c.Name
	company.Type = int(cType)

	if cType == CONTRACT && len(c.Intermediateds) > 0 {
		return nil, ContractIntermediatedErr
	}

	if cType == AGGREGATE && len(c.Intermediateds) == 0 {
		return nil, MissingIntermediatedErr
	} else {
		for _, i := range c.Intermediateds {
			company.Intermediateds = append(company.Intermediateds, Company{Name: i.Name, Type: int(INTERMEDIATED)})
		}
	}

	return &company, nil
}

func (c CompanyPatchSchema) parse(id int) map[string]any {
	var m map[string]any = make(map[string]any)

	if c.Name != nil {
		m["name"] = *c.Name
	}
	if c.Type != nil {
		m["type"] = int(NewCompanyType(*c.Type))
	}

	return m
}

func ValidateCompanyType(fl validator.FieldLevel) bool {
	ct := fl.Field().Interface().(string)

	if isCompanyType(ct) {
		return true
	}

	return false
}

func ValidateOptionalCompanyType(fl validator.FieldLevel) bool {
	ct := fl.Field().Interface().(string)

	if isCompanyType(ct) {
		return true
	}

	return false
}
