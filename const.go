package main

import (
	"fmt"
	"math"
)

func main() {

	const pi = math.Pi
	const e = math.E

	radius := 5.0
	circleArea := pi * math.Pow(radius, 2)
	exponential := math.Pow(e, 2)

	fmt.Printf("Значение π: %.5f\n", pi)
	fmt.Printf("Значение e: %.5f\n", e)
	fmt.Printf("Площадь круга с радиусом %.1f: %.2f\n", radius, circleArea)
	fmt.Printf("e^2: %.2f\n", exponential)
}
