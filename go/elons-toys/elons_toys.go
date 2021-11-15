package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car)Drive() Car {
	if car.batteryDrain > car.battery {
		return *car
	}
	car.battery = car.battery - car.batteryDrain
	car.distance = car.speed
	return *car
}
// TODO: define the 'DisplayDistance() string' method
func(car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func(car *Car) DisplayBattery() string {
	message := fmt.Sprintf("Battery at %d", car.battery)
	return message + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func(car *Car) CanFinish(trackDistance int) bool {
	if trackDistance/car.speed <= car.battery/car.batteryDrain {
		return true
	} else {
		return false
	}
}