package main

import (
	"fmt"
)

func main() {
	prices   := []float64{ 10, 20, 30 }
	taxRates := []float64{ 0.1, 0.25, 0.5, 0.7 }

	result := make( map[float64][]float64 )

	for _, taxRate := range taxRates {
		pricesPerTaxRate := make([]float64, len( prices ) )
		for priceIdx, price := range prices {
			pricesPerTaxRate[ priceIdx ] = price * ( 1 + taxRate )
		}

		result[ taxRate ] = pricesPerTaxRate
	}

	fmt.Println( result )
}