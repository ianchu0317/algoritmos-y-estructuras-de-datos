package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"
)

func compararInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}

	return 0
}

func TestGuardar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	abb.Guardar(5, 3)
	abb.Guardar(3, 4)
	abb.Guardar(8, 8)
	abb.Guardar(9, 8)
	abb.Guardar(13, 6)
	abb.Guardar(7, 7)
	fmt.Println("test1:", abb.Obtener(3))
	abb.Guardar(3, 2)
	//require.True(t, abb.Pertenece(13), "Elemento guardado tiene que pertenecer a ABB")
	//require.True(t, abb.Pertenece(9), "Elemento 2 guardado tiene que pertenecer a ABB")
	fmt.Println("test2:", abb.Obtener(3))
	fmt.Println("Cantidad antes de borrar:", abb.Cantidad())
	fmt.Println("test borrar:", abb.Borrar(9))
	fmt.Println("Test pertenece", abb.Obtener(13))
	fmt.Println("Cantidad despues de borrar:", abb.Cantidad())
}

func TestIteradorInterno(t *testing.T) {
	// Test iterador interno sin rango
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	abb.Guardar(7, 7)
	abb.Guardar(6, 6)
	abb.Guardar(9, 9)
	abb.Guardar(10, 10)
	abb.Guardar(3, 3)
	abb.Guardar(5, 5)
	abb.Guardar(8, 8)

	abb.Iterar(func(clave int, dato int) bool {
		fmt.Println("Clave, dato:", clave, dato)
		return true
	})

	fmt.Println("Debug Iterador interno rangos:")
	a, b := 6, 9
	abb.IterarRango(&a, &b, func(clave int, dato int) bool {
		fmt.Println("Clave, dato:", clave, dato)
		return true
	})
}

func TestIteradorExterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	abb.Guardar(7, 7)
	abb.Guardar(6, 6)
	abb.Guardar(9, 9)
	abb.Guardar(10, 10)
	abb.Guardar(3, 3)
	abb.Guardar(5, 5)
	abb.Guardar(8, 8)

	iter := abb.Iterador()

	for iter.HaySiguiente() {
		clave, datos := iter.VerActual()
		fmt.Println("Iterador (clave, datos):", clave, datos)
		iter.Siguiente()
	}

	a, b := 6, 9
	iter = abb.IteradorRango(&a, &b)
	for iter.HaySiguiente() {
		clave, datos := iter.VerActual()
		fmt.Println("Iterador (clave, datos):", clave, datos)
		iter.Siguiente()
	}
}
