package main

import (
	"fmt"
	Pila "tdas/pila"
)

func main() {
	miPila := Pila.CrearPilaDinamica[int]()
	// probar apilar
	for i := 1; i <= 10; i++ {
		miPila.Apilar(i)
		fmt.Println(miPila)
	}
	// probar desapilar
	for i := 1; i <= 10; i++ {
		tope := miPila.Desapilar()
		fmt.Println(miPila, tope)
	}
}
