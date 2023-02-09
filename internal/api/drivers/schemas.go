package drivers

import "strconv"

type DriverPostSchema struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	LicenseA bool   `json:"license_a" binding:"required"`
	LicenseB bool   `json:"license_b" binding:"required"`
	LicenseC bool   `json:"license_c" binding:"required"`
	LicenseD bool   `json:"license_d" binding:"required"`
}

type DriverPatchSchema struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	LicenseA bool   `json:"license_a"`
	LicenseB bool   `json:"license_b"`
	LicenseC bool   `json:"license_c"`
	LicenseD bool   `json:"license_d"`
}

func (d DriverPostSchema) parse() *Driver {
	driver := Driver{}
	driver.Name = d.Name
	driver.LicenseA = d.LicenseA
	driver.LicenseB = d.LicenseB
	driver.LicenseC = d.LicenseC
	driver.LicenseD = d.LicenseD
	return &driver
}

func (d DriverPatchSchema) parse(_id string) (*Driver, error) {
	driver := Driver{}

	id, err := strconv.Atoi(_id)
	if err != nil {
		return nil, err
	}
	driver.ID = id
	driver.Name = d.Name
	driver.LicenseA = d.LicenseA
	driver.LicenseB = d.LicenseB
	driver.LicenseC = d.LicenseC
	driver.LicenseD = d.LicenseD
	return &driver, err
}
