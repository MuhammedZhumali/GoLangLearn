package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var res float64
	res = float64(productionRate) * (successRate / 100)
	return res
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var res int
	res = int(float64(productionRate) * (successRate / 100) / 60)
	return res
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var firstDigit int
	firstDigit = carsCount / 10
	var secondDigit int
	secondDigit = carsCount % 10
	var res uint
	res = uint(firstDigit*95000 + secondDigit*10000)
	return res

}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(426, 80))
	fmt.Println(CalculateCost(21))
}
