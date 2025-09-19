// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa

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
var OPERACIONES = []string{"+", "-", "*", "/", "sqrt", "log"}
var ERROR = "ERROR"

// esOperacion devuelve si el caracter es una operación o no
func esOperacion(caracter string) bool {
	return slices.Contains(OPERACIONES, caracter)
}

// calcularOperacion calcula el resultado aplicando la operación adecuada
func calcularOperacion(operandos TDAPila.Pila[int], operacion string) {

}

// calcularOperacion calcula la operación pasada en formato polaco inverso
func procesarLinea(linea string) {
	operandos := TDAPila.CrearPilaDinamica[int]()
	// Dividir espacios del string y convertir en array
	// https://pkg.go.dev/strings#FieldsSeq
	for caracter := range strings.FieldsSeq(linea) {
		if esOperacion(caracter) {
			calcularOperacion(operandos, caracter)
		} else {
			num, _ := strconv.Atoi(caracter)
			operandos.Apilar(num)
			fmt.Println(operandos)
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
