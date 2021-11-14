package main

import (
	"fmt"
	"math/big"
)

func main() {
	numerator, denominator := big.NewInt(3), big.NewInt(2)
	count := 0
	for i := 0; i < 1000; i++ {
		numerator, denominator = nextExpression(numerator, denominator)
		nStr := numerator.Text(10)
		dStr := denominator.Text(10)
		if len(nStr) > len(dStr) {
			count++
		}
	}

	fmt.Printf("There are %d fractions that contain a numerator with more digits than the denominator", count)
}

func nextExpression(numerator *big.Int, denominator *big.Int) (*big.Int, *big.Int) {
	// 1 + n/d
	x, y := numerator.Add(numerator, denominator), denominator
	// 1/x
	x, y = y, x
	// final expression
	x, y = x.Add(x, y), y
	return x, y
}
