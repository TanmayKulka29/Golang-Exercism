package elon
import ("fmt")
// TODO: define the 'Drive()' method
func (car *Car) Drive() {
    if car.battery >= car.batteryDrain {
        car.distance += car.speed
        car.battery -= car.batteryDrain
    }
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
    return float64(trackDistance-car.distance)/float64(car.speed) <= float64(car.battery)/float64(car.batteryDrain)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.





package elon

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int

	battery  int
	distance int
}

// NewCar creates a new car with given specifications.
func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}
