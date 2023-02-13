package companies

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type CompanySchema struct {
	ID        int       `json:"id" `
	CreatedAt time.Time `json:"created_at" `
	Name      string    `json:"name"`
	Type      string    `json:"type"`
}

type CompanyPostSchema struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required_company_type"`
}

type CompanyPatchSchema struct {
	Name *string `json:"name"`
	Type *string `json:"type" binding:"optional_company_type"`
}

func (c CompanyPostSchema) parse() *Company {
	company := Company{}
	company.Name = c.Name
	company.Type = int(toCompanyType(c.Type))
	return &company
}

func (c CompanyPatchSchema) parse(id int) map[string]any {
	var m map[string]any = make(map[string]any)

	if c.Name != nil {
		m["name"] = *c.Name
	}
	if c.Type != nil {
		m["type"] = int(toCompanyType(*c.Type))
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
