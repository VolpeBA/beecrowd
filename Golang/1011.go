package main

import (
	"fmt"
	"math"
)

func main() {
	//volume é: (4/3) * pi * R3
	var raio float64

	fmt.Scan(&raio)

	pi := 3.14159
	volume := (4.0 / 3.0) * pi * math.Pow(raio, 3)
	fmt.Printf("VOLUME = %.3f\n", volume)
}
