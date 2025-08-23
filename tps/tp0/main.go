package main

import (
	"fmt"

	"./ejercicios"
)

func main() {
	a, b := 5, 6
	fmt.Println("Las variables son a, b=", a, b)
	ejercicios.Swap(&a, &b)
	fmt.Println("Las variables son a, b=", a, b)
}
