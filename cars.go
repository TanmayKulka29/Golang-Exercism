package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate)*(successRate/100.0)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var ans int = int(float64(productionRate)*(successRate/100.0)/60)
    return ans
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var cost uint = 0
    cost += (( uint(carsCount) / 10 ) * 95000)
    cost += (( uint(carsCount) % 10 ) * 10000)
	return cost
	panic("CalculateCost not implemented")
}
