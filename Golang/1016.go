package main

import (
	"fmt"
)

func main() {

	var km float64

	fmt.Scan(&km)

	km *= 2.0

	fmt.Printf("%.0f minutos\n", km)
}
