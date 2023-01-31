package main

type BuildProcess interface {
    SetWheels() BuildProcess
    SetSeats() BuildProcess
    SetStructure() BuildProcess
    GetVehicle() VehicleProduct
}

type Director interface {
    SetBuilder(b BuildProcess)
    Construct()
}

type VehicleProduct struct {
    Wheels    int
    Seats     int
    Structure string
}

type manufacturingDirector struct {
    builder BuildProcess
}

func (m *manufacturingDirector) Construct() {
    m.builder.SetSeats().SetStructure().SetWheels()
}

func (m *manufacturingDirector) SetBuilder(b BuildProcess) {
    m.builder = b
}

func NewDirector() Director {
    return &manufacturingDirector{}
}
