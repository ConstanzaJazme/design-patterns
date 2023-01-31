package vehicle_builder

import "fmt"

func Start() {
	manufacturingComplex := NewDirector()

    carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	fmt.Printf("Number of wheels %d\n", car.Wheels)
	fmt.Printf("Car type %s\n", car.Structure)
	fmt.Printf("Number of seats %d\n", car.Seats)

    bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

    fmt.Printf("Number of wheels %d\n", motorbike.Wheels)
    fmt.Printf("Car type %s\n", motorbike.Structure)
    fmt.Printf("Number of seats %d\n", motorbike.Seats)
}
