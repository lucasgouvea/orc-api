package vehicles

type VehicleSchema struct {
	Id               string `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	ModelDescription string `json:"model_description"`
	LicensePlate     string `json:"license_plate"`
}

type VehiclePostSchema struct {
	Model        string `json:"model" binding:"required"`
	LicensePlate string `json:"license_plate" binding:"required"`
}

type VehiclePatchSchema struct {
	Model        *string `json:"model"`
	LicensePlate *string `json:"license_plate"`
}

func (v VehiclePostSchema) parse() *Vehicle {
	vehicle := Vehicle{}
	vehicle.ModelDescription = v.Model
	vehicle.LicensePlate = v.LicensePlate
	return &vehicle
}

func (v VehiclePatchSchema) parse(id int) map[string]any {
	var m map[string]any = make(map[string]any)

	if v.Model != nil {
		m["model"] = *v.Model
	}
	if v.LicensePlate != nil {
		m["license_plate"] = *v.LicensePlate
	}

	return m
}
