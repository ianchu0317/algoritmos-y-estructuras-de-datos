package main

import (
	"fmt"
	Pila "tdas/pila"
)

func main() {
	miPila := Pila.CrearPilaDinamica[int]()
	for i := 1; i <= 10; i++ {
		miPila.Apilar(i)
		fmt.Println(miPila)
	}
}
