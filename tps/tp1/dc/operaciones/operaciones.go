package operaciones

import (
	AUX "dc/auxiliares"
	"fmt"
	"math"
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

func sumar(a, b int64) int64 {
	return a + b
}

func restar(a, b int64) int64 {
	return a - b
}

func multiplicar(a, b int64) int64 {
	return a * b
}

func dividir(a, b int64) int64 {
	return a / b
}

func sqrt(n int64) int64 {
	return int64(math.Sqrt(float64(n)))
}

func exp(a, b int64) int64 {
	// a^b
	return int64(math.Pow(float64(a), float64(b)))
}

func logBaseB(a, b int64) int64 {
	// LogB(A) = log(a)/log(b)
	return int64(math.Log(float64(a)) / math.Log(float64(b)))
}

func tern(a, b, c int64) int64 {
	var resultado int64
	if a != 0 {
		resultado = b
	} else {
		resultado = c
	}
	return resultado
}

// FUNCIONES A EXPORTAR

// calcularOperacion calcula el resultado aplicando la operación adecuada
func CalcularOperacion(operandos TDAPila.Pila[int], operacion string) {
	//desapilarCantidadEnSlice(operandos, 2)
	fmt.Println(AUX.DesapilarCantidadN(operandos, 2))
}
