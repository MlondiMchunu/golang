package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	fmt.Print(`Enter investment amount : `)
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Printf(`future value = $%.2f`, futureValue)
	fmt.Println()
	fmt.Printf(`inflation adjusted value = $%.2f`, inflationAdjustedValue)
	fmt.Println()
}
