package speed

// TODO: define the 'Car' type struct
type Car struct { 
    battery int
    batteryDrain int
	speed int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    new_car := Car {
    	speed : speed,
    	batteryDrain : batteryDrain,
        battery:     100, 
        distance:    0,  
    }
    return new_car
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
    ntrack := Track {
    	distance : distance,
    }
    return ntrack
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery>=car.batteryDrain {
        car.distance += car.speed
        car.battery -= car.batteryDrain  
    }
    return car
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    val := (car.battery / car.batteryDrain) * car.speed
	if val < track.distance {
		return false
	}
	return true
	panic("Please implement the CanFinish function")
}
