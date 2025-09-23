package operaciones

import (
	AUX "dc/auxiliares"
	"errors"
	"math"
	TDAPila "tdas/pila"
)

// Variables globales
const (
	CANT_OPERANDO_SQRT = 1
	CANT_OPERANDO_TERN = 3
)

// Funciones Matemáticas (internas)

// sumar toma dos eneteros (a,b) y devuelve a+b.
// error siempre es nil
func sumar(a, b int64) (int64, error) {
	return a + b, nil
}

// restar toma dos enteros (a,b) y devuelve a-b.
// error siempre es nil
func restar(a, b int64) (int64, error) {
	return a - b, nil
}

// multiplicar toma dos enteros (a,b) y devuelve a*b.
// error siempre es nil
func multiplicar(a, b int64) (int64, error) {
	return a * b, nil
}

// dividir toma dos enteros (a,b) y devuelve a/b.
// devuelve -1 y 'division por cero' si hay división por cero.
// devuelve nil cuando operación es válida
func dividir(a, b int64) (int64, error) {
	if b == 0 {
		return -1, errors.New("error division por cero")
	}
	return (a / b), nil
}

// exp toma dos enteros y devuelve la exponencial a^b.
// devuelve nil cuando operación es válida
func exp(a, b int64) (int64, error) {
	// a^b
	return int64(math.Pow(float64(a), float64(b))), nil
}

// logBaseB toma dos enteros (a,b) y calcula logaritmo de A en base B.
// devuelve -1 con error 'base inválida' si la base b es inválida
// error es nil si operación es válida
func logBaseB(a, b int64) (int64, error) {
	// LogB(A) = log(a)/log(b)
	arg := float64(a)
	base := float64(b)
	if base <= 1 {
		return -1, errors.New("base logaritmo inválido")
	}
	return int64(math.Log(arg) / math.Log(base)), nil
}

func sqrt(n int64) int64 {
	return int64(math.Sqrt(float64(n)))
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
func calcDosVariables(operandos TDAPila.Pila[int64], operar func(int64, int64) (int64, error)) error {
	// Intentar desapilar la cantidad de variables a utilizar
	// devolver error si no se ingresaron necesario
	enteros, err := AUX.DesapilarCantidadN(operandos, 2)
	if err != nil {
		return err
	}
	// Intentar realizar cálculo
	// devolver error si la operación es inválida
	resultado, err := operar(enteros[0], enteros[1])
	if err != nil {
		return err
	}
	// Volver a apilar el resultado y devolver nil
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
