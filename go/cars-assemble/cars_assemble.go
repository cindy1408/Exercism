package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
        return 0.0
    } else if 0 < speed && speed < 5 {
    return 1.0
    } else if 4 < speed && speed < 9 {
    return 0.9
    } else {
    return 0.77
    }
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	carPerHour := 221
    successRate := SuccessRate(speed)
    producedCars := float64(speed) * float64(carPerHour) * successRate
    return producedCars
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	producedCarsPerHour := CalculateProductionRatePerHour(speed)
    return int(producedCarsPerHour/60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	producedCarsPerHour := CalculateProductionRatePerHour(speed)
    if producedCarsPerHour < limit {
        return producedCarsPerHour
    } else {
    	return float64(limit)
    }
}
