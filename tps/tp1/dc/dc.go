// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// extraerSiguienteDigito
// func extraerSiguienteDigito

// calcularOperacion calcula la operación en formato polaco inverso
func calcularOperacion(linea string) {
	miArrayDeLetras := strings.Fields(linea)
	fmt.Println(miArrayDeLetras)
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
