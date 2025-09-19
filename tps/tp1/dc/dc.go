// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa
// Documentacion string utilizado: // https://pkg.go.dev/strings#FieldsSeq

package main

import (
	"bufio"
	AUX "dc/auxiliares"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
)

// calcularOperacion calcula el resultado aplicando la operación adecuada
func calcularOperacion(operandos TDAPila.Pila[int], operacion string) {
	//desapilarCantidadEnSlice(operandos, 2)
	fmt.Println(AUX.DesapilarCantidadN(operandos, 2))
}

// calcularOperacion calcula la operación pasada en formato polaco inverso
func procesarLinea(linea string) {
	operandos := TDAPila.CrearPilaDinamica[int]()
	// Dividir espacios del string e ir por cada elemento
	for caracter := range strings.FieldsSeq(linea) {
		if AUX.EsOperacion(caracter) {
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
