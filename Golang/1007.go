package main

import (
	"fmt"
)

func main() {
//DIFERENCA = (A * B - C * D).

  var a,b,c,d int

  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Scan(&c)
  fmt.Scan(&d)
  
  diferenca := ( a * b - c * d )

  fmt.Printf("DIFERENCA = %d\n", diferenca)
}
