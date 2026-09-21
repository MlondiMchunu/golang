package main

import (
	"fmt"
	"math"
)

func calc() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print(`Enter investment amount : `)
	fmt.Scan(&investmentAmount)

	fmt.Print(`enter no. of years : `)
	fmt.Scan(&years)

	fmt.Print(`expected return rate : `)
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	inflationAdjustedValue := futureValue * math.Pow(1+inflationRate/100, years)

	fmt.Printf(`future value = $%.2f`, futureValue)
	fmt.Println()
	fmt.Printf(`inflation adjusted value = $%.2f`, inflationAdjustedValue)
	fmt.Println()
}
