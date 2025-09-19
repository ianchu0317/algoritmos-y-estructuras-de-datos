package operaciones

import (
	AUX "dc/auxiliares"
	"fmt"
	TDAPila "tdas/pila"
)

// Variables globales
var OPERACIONES = []string{"+", "-", "*", "/", "sqrt", "^", "log", "?"}

const (
	CANT_OPERANDO_SUMA  = 2
	CANT_OPERANDO_RESTA = 2
	CANT_OPERANDO_MULT  = 2
	CANT_OPERANDO_DIV   = 2
	CANT_OPERANDO_SQRT  = 1
	CANT_OPERANDO_EXP   = 2
	CANT_OPERANDO_LOG   = 2
	CANT_OPERANDO_TERN  = 3
)

// Funciones Matemáticas (internas)
func Sumar(a, b int64) int64 {
	return a + b
}

func Restar(a, b int64) int64 {
	return a - b
}

func Multiplicar(a, b int64) int64 {
	return a * b
}

func Dividir(a, b int64) int64 {
	return a / b
}

// FUNCIONES A EXPORTAR

// calcularOperacion calcula el resultado aplicando la operación adecuada
func CalcularOperacion(operandos TDAPila.Pila[int], operacion string) {
	//desapilarCantidadEnSlice(operandos, 2)
	fmt.Println(AUX.DesapilarCantidadN(operandos, 2))
}
