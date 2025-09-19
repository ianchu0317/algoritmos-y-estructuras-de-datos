package operaciones

import (
	AUX "dc/auxiliares"
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

// calcDosVariables calcula la operacion de dos variables
func calcDosVariables(operandos TDAPila.Pila[int64], operar func(int64, int64) int64) error {
	enteros, err := AUX.DesapilarCantidadN(operandos, 2)
	if err != nil {
		return err
	}
	resultado := operar(enteros[0], enteros[1])
	operandos.Apilar(resultado)
	return nil
}

// FUNCIONES A EXPORTAR

// calcularOperacion calcula el resultado aplicando la operación adecuada
func CalcularOperacion(operandos TDAPila.Pila[int64], operacion string) error {
	switch operacion {
	case "+":
		return calcDosVariables(operandos, sumar)
	case "-":
		return calcDosVariables(operandos, restar)
	case "*":
		return calcDosVariables(operandos, multiplicar)
	case "/":
		return calcDosVariables(operandos, dividir)
	case "^":
		return calcDosVariables(operandos, exp)
	case "log":
		return calcDosVariables(operandos, logBaseB)
	case "sqrt":
		enteros, err := AUX.DesapilarCantidadN(operandos, CANT_OPERANDO_SQRT)
		if err != nil {
			return err
		}
		operandos.Apilar(sqrt(enteros[0]))
		return nil
	case "?":
		enteros, err := AUX.DesapilarCantidadN(operandos, CANT_OPERANDO_TERN)
		if err != nil {
			return err
		}
		operandos.Apilar(tern(enteros[0], enteros[1], enteros[2]))
		return nil
	}
	return nil
}
