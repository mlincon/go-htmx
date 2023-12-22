/*
Complete the code in the function calculate to return a slice consisting of 4
elements [ sum of a and b, difference of a and b, product of a and b, quotient on dividing a by
*/

package main

import "fmt"

func calculate(a int, b int) []float64 {
	sum := float64(a + b)
	difference := float64(a - b)
	product := float64(a * b)
	quotient := float64(a / b)

	var results []float64 = []float64{
		sum,
		difference,
		product,
		quotient,
	}
	return results
}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
