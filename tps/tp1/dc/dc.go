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
	var opErr error = nil
	var huboOp bool = false

	// Dividir espacios del string e ir por cada elemento
	// Si es operador, calcular operación con los ultimos elementos de la pila
	// Si es numero, apilar a operandos
	for caracter := range strings.FieldsSeq(linea) {
		if AUX.EsOperacion(caracter) {
			opErr = OP.CalcularOperacion(operandos, caracter)
			huboOp = true
			if opErr != nil {
				break // Salir de operacion si hay error
			}
		} else {
			num, _ := strconv.ParseInt(caracter, 10, 64) // Base 10, bitsize 64
			operandos.Apilar(num)
		}
	}

	// Imprimir resultado cuando hubo operación y sin errores
	if opErr == nil && huboOp {
		fmt.Println(operandos.Desapilar())
	} else {
		fmt.Println("ERROR")
	}
}

func main() {
	// crear un scanner y leer STDIN
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lineaActual := s.Text()
		//fmt.Println("\nLinea es:", lineaActual)
		procesarLinea(lineaActual)
	}
}
