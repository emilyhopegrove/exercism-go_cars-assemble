package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
/* 
Check if the carsCount is divisible by 10. If there are complete groups of ten cars,
	calculate the cost by multiplying the number of groups by the cost of producing a group of ten cars, which is $95,000.
If there are any remaining cars that are less than a group of ten,
	calculate the cost by multiplying the number of remaining cars by the cost of producing a single car, which is $10,000.
Add together the costs from the groups and the remaining cars to get the total cost of producing the given number of cars.
Return the total cost as the result.
*/
func CalculateCost(carsCount int) uint {
	const (
		costPerGroup = uint(95000) // Cost of producing 10 cars together
		costPerCar   = uint(10000) // Cost of producing an individual car
	)

	return uint(carsCount/10)*costPerGroup + uint(carsCount%10)*costPerCar
}
