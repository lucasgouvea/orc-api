package drivers

type DriverSchema struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	LicenseA  bool   `json:"license_a"`
	LicenseB  bool   `json:"license_b"`
	LicenseC  bool   `json:"license_c"`
	LicenseD  bool   `json:"license_d"`
	LicenseE  bool   `json:"license_e"`
}

type DriverPostSchema struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required,min=18,max=100"`
	LicenseA *bool  `json:"license_a" binding:"required"`
	LicenseB *bool  `json:"license_b" binding:"required"`
	LicenseC *bool  `json:"license_c" binding:"required"`
	LicenseD *bool  `json:"license_d" binding:"required"`
	LicenseE *bool  `json:"license_e" binding:"required"`
}

type DriverPatchSchema struct {
	Name     *string `json:"name"`
	Age      *int    `json:"age" binding:"omitempty,min=18,max=100"`
	LicenseA *bool   `json:"license_a"`
	LicenseB *bool   `json:"license_b"`
	LicenseC *bool   `json:"license_c"`
	LicenseD *bool   `json:"license_d"`
	LicenseE *bool   `json:"license_e"`
}

func (d DriverPostSchema) parse() *Driver {
	driver := Driver{}
	driver.Name = d.Name
	driver.Age = d.Age
	driver.LicenseA = *d.LicenseA
	driver.LicenseB = *d.LicenseB
	driver.LicenseC = *d.LicenseC
	driver.LicenseD = *d.LicenseD
	driver.LicenseE = *d.LicenseE
	return &driver
}

func (d DriverPatchSchema) parse() map[string]any {
	var m map[string]any = make(map[string]any)

	if d.Name != nil {
		m["name"] = *d.Name
	}
	if d.Age != nil {
		m["age"] = *d.Age
	}
	if d.LicenseA != nil {
		m["license_a"] = *d.LicenseA
	}
	if d.LicenseB != nil {
		m["license_b"] = *d.LicenseB
	}
	if d.LicenseC != nil {
		m["license_c"] = *d.LicenseC
	}
	if d.LicenseD != nil {
		m["license_d"] = *d.LicenseD
	}
	if d.LicenseE != nil {
		m["license_e"] = *d.LicenseE
	}

	return m
}
