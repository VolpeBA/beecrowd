package main

import (
	"fmt"
)

func main() {

	var id_funcionario int
	var horas_trabalhadas, valor_hora float64

	fmt.Scan(&id_funcionario)
	fmt.Scan(&horas_trabalhadas)
	fmt.Scan(&valor_hora)

	salario := horas_trabalhadas * valor_hora

	fmt.Printf("NUMBER = %d\n", id_funcionario)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}
