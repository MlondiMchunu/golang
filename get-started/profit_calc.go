package main

import (
	"fmt"
)

func profit() {
	var revenue, expenses, taxRate float64

	fmt.Print(`revenue : `)
	fmt.Scan(&revenue)

	fmt.Print(`expenses : `)
	fmt.Scan(&expenses)

	fmt.Print(`tax rate :`)
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(`ebt = `, ebt)
	fmt.Println(`profit = `, profit)
	fmt.Println(`ratio = `, ratio)
}
