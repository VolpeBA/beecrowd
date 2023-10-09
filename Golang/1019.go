package main

import (
	"fmt"
)

func main() {

	var n, minutos, horas int

	fmt.Scan(&n)

	horas = n / 3600
	n %= 3600
	minutos = n / 60
	n %= 60
	segundos := n

	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)
}
