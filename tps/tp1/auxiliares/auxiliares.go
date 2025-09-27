package auxiliares

import (
	"errors"
	"slices"
	TDAPila "tdas/pila"
)

// VARIABLES GLOBALES
var OPERACIONES = []string{"+", "-", "*", "/", "sqrt", "^", "log", "?"}

// FUNCIONES AUXILIARES
// esOperacion devuelve si el caracter es una operación o no
func EsOperacion(caracter string) bool {
	return slices.Contains(OPERACIONES, caracter)
}

// invertirSlice invierte un slice [a, b] -> [b, a]
func invertirSlice(arr []int64) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// desapilarCantidadN desapila 'n' elementos y los devuelve en Slice.
// Si la pila es [a, b, c, tope] y n=2 -> [b, c]
// Devuelve error si no hay cantidad solicitada en la pila
func DesapilarCantidadN(pila TDAPila.Pila[int64], n int) ([]int64, error) {
	resultado := make([]int64, 0)
	for range n {
		if !pila.EstaVacia() {
			num := (pila).Desapilar()
			resultado = append(resultado, num)
		} else {
			invertirSlice(resultado)
			return resultado, errors.New("elementos insuficientes en la pila")
		}
	}
	invertirSlice(resultado)
	return resultado, nil
}
