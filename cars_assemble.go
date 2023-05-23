package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	return float64(productionRate) * float64(successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)

}

// CalculateCost works out the cost of producing the given number of cars.

/* Check if the carsCount is divisible by 10. 
This step determines if there are any complete groups of ten cars that can be produced.

	If there are complete groups of ten cars, 
	calculate the cost by multiplying the number of groups by the cost of producing a group of ten cars, which is $95,000.

If there are any remaining cars that are less than a group of ten, calculate the cost by multiplying the number of remaining cars by the cost of producing a single car, which is $10,000.

Add together the costs from steps 2 and 3 to get the total cost of producing the given number of cars. 

Return the total cost as the result. */
func CalculateCost(carsCount int) uint {
	costPerGroup := uint(95000)      // Cost of producing 10 cars together
    costPerCar := uint(10000)        // Cost of producing an individual car
    groupCount := carsCount / 10     // Number of groups of 10 cars
    remainingCars := carsCount % 10  // Number of individual cars remaining

    totalCost := uint(groupCount) * costPerGroup + uint(remainingCars) * costPerCar
    return totalCost
}
