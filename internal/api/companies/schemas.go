package companies

import (
	"github.com/go-playground/validator/v10"
)

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
	company.Type = toCompanyType(c.Type).Int()
	return &company
}

func (c CompanyPatchSchema) parse(id int) map[string]any {
	var m map[string]any = make(map[string]any)

	if c.Name != nil {
		m["name"] = *c.Name
	}
	if c.Type != nil {
		m["type"] = toCompanyType(*c.Type).Int()
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
