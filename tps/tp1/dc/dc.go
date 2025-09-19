// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// extraerSiguienteDigito
// func extraerSiguienteDigito
var OPERACIONES = []string{"+", "-", "*", "/", "sqrt", "log"}

func esOperacion(caracter string) bool {
	return slices.Contains(OPERACIONES, caracter)
}

// calcularOperacion calcula la operación en formato polaco inverso
func calcularOperacion(linea string) {
	// Dividir espacios del string y convertir en array
	miArrayDeLetras := strings.Fields(linea)
	fmt.Println(miArrayDeLetras)
	for _, caracter := range miArrayDeLetras {
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
		fmt.Println("")
		fmt.Println("\nLinea es:", lineaActual)
		calcularOperacion(lineaActual)
	}
}
