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
	fmt.Println("test borrar:", abb.Borrar(9))
	fmt.Println("Test pertenece", abb.Obtener(13))
}
