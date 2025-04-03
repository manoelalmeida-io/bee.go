package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a+b > c && b+c > a && a+c > b {
		fmt.Printf("Perimetro = %.1f\n", (a + b + c))
	} else {
		fmt.Printf("Area = %.1f\n", (0.5 * (a + b) * c))
	}
}
