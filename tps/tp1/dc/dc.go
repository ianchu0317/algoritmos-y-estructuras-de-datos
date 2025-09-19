// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa
// Documentacion string utilizado: // https://pkg.go.dev/strings#FieldsSeq

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

// VARIABLES GLOBALES
var OPERACIONES = []string{"+", "-", "*", "/", "?", "sqrt", "log"}
var ERROR = "ERROR"

// FUNCIONES AUXILIARES
// esOperacion devuelve si el caracter es una operación o no
func esOperacion(caracter string) bool {
	return slices.Contains(OPERACIONES, caracter)
}

// invertirSlice invierte un slice [a, b] -> [b, a]
func invertirSlice(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

// desapilarCantidadEnArray desapila 'n' elementos y los devuelve en Slice.
// Si la pila es [a, b, c, tope] y n=2 -> [b, c]
func desapilarCantidadEnSlice(pila TDAPila.Pila[int], n int) []int {
	resultado := make([]int, 0)
	for range n {
		if !pila.EstaVacia() {
			num := (pila).Desapilar()
			resultado = append(resultado, num)
		} else {
			fmt.Fprintln(os.Stderr, ERROR)
			invertirSlice(resultado)
			return resultado
		}
	}
	invertirSlice(resultado)
	return resultado
}

// calcularOperacion calcula el resultado aplicando la operación adecuada
func calcularOperacion(operandos TDAPila.Pila[int], operacion string) {
	desapilarCantidadEnSlice(operandos, 2)
	//fmt.Println(desapilarCantidadEnSlice(operandos, 2))
}

// calcularOperacion calcula la operación pasada en formato polaco inverso
func procesarLinea(linea string) {
	operandos := TDAPila.CrearPilaDinamica[int]()
	// Dividir espacios del string e ir por cada elemento
	for caracter := range strings.FieldsSeq(linea) {
		if esOperacion(caracter) {
			calcularOperacion(operandos, caracter)
		} else {
			num, _ := strconv.Atoi(caracter)
			operandos.Apilar(num)
		}
	}
}

func main() {
	// crear un scanner y leer STDIN
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lineaActual := s.Text()
		fmt.Println("\nLinea es:", lineaActual)
		procesarLinea(lineaActual)
	}
}
