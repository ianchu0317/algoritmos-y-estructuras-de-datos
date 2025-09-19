// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// VARIABLES GLOBALES
var OPERACIONES = []string{"+", "-", "*", "/", "sqrt", "log"}
var ERROR = "ERROR"

// esOperacion devuelve si el caracter es una operación o no
func esOperacion(caracter string) bool {
	return slices.Contains(OPERACIONES, caracter)
}

// calcularOperacion calcula la operación pasada en formato polaco inverso
func calcularOperacion(linea string) {
	// Dividir espacios del string y convertir en array
	//miArrayDeLetras := strings.Fields(linea)

	// fmt.Println(miArrayDeLetras)
	for caracter := range strings.FieldsSeq(linea) {
		if esOperacion(caracter) {
			fmt.Println(caracter)
		}
	}
}

func main() {
	// crear un scanner y leer stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lineaActual := s.Text()
		fmt.Println("\nLinea es:", lineaActual)
		calcularOperacion(lineaActual)
	}
}
