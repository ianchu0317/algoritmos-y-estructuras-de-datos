package main

// Pila "tdas/pila"
import (
	"fmt"
	Dict "tdas/diccionario"
)

func igualdadStrings(a, b string) bool {
	return a == b
}

type basico struct {
	a string
	b int
}
type avanzado struct {
	w int
	x basico
	y basico
	z string
}

func main() {
	/*
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
	*/

	// TEST DICCIONARIO
	//dic := Dict.CrearHash[string, string](igualdadStrings)
	//fmt.Println(dic.Pertenece("A"))
	//dic.Guardar("A", "Autito")
	//fmt.Println("Luego de guardar:", dic.Obtener("A"))

	dic := Dict.CrearHash[avanzado, int](func(a, b avanzado) bool {
		return a.w == b.w && a.z == b.z && a.x.a == b.x.a && a.x.b == b.x.b && a.y.a == b.y.a && a.y.b == b.y.b
	})

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)
	fmt.Println(dic.Pertenece(a1))
}
