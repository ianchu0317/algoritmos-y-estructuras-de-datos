// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa
// Documentacion string utilizado: // https://pkg.go.dev/strings#FieldsSeq

package main

import (
	"bufio"
	AUX "dc/auxiliares"
	OP "dc/operaciones"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

// calcularOperacion calcula la operación pasada en formato polaco inverso
func procesarLinea(linea string) {
	operandos := TDAPila.CrearPilaDinamica[int64]()
	// Dividir espacios del string e ir por cada elemento
	for caracter := range strings.FieldsSeq(linea) {
		if AUX.EsOperacion(caracter) {
			// Si es operador, calcular operación con los ultimos elementos de la pila
			OP.CalcularOperacion(operandos, caracter)
		} else {
			// Si es numero, apilar
			num, _ := strconv.ParseInt(caracter, 10, 64) // Base 10, bitsize 64
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
