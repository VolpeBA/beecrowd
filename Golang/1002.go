package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.14159
	var raio float64

	fmt.Scan(&raio)
	area := pi * math.Pow(raio, 2)
	fmt.Printf("A=%.4f\n", area)
}
