/*
Implementar el TDA Fracción.
Dicho TDA debe tener las siguientes primitivas:

CrearFraccion(numerador, denominador int) Fraccion
Sumar(otra Fraccion) Fraccion
Multiplicar(otra Fraccion) Fraccion
ParteEntera() int
Representacion() string
*/

package tda_01

import "fmt"

type Fraccion struct {
	numerador, denominador int
}

// Sumar recibe una fraccion y devuelve la suma de la pasada con la actual
// Complejidad O(1)
func (fraccion Fraccion) Sumar(otra Fraccion) Fraccion {
	nuevoNumerador := fraccion.numerador*otra.denominador + otra.numerador*fraccion.denominador
	nuevoDenominador := fraccion.denominador * otra.denominador
	nuevaFraccion := Fraccion{numerador: nuevoNumerador, denominador: nuevoDenominador}
	return nuevaFraccion
}

// Multiplicar recibe una fraccion y devuelve la multiplicacion de ambas
func (fraccion Fraccion) Multiplicar(otra Fraccion) Fraccion {
	nuevoNumerador := fraccion.numerador * otra.numerador
	nuevoDenominador := fraccion.denominador * otra.denominador
	nuevaFraccion := Fraccion{numerador: nuevoNumerador, denominador: nuevoDenominador}
	return nuevaFraccion
}

// ParteEntera devuelve la parte entera de la fraccion
func (fraccion Fraccion) ParteEntera() int {
	return int(fraccion.numerador / fraccion.denominador)
}

// Representacion devuelve la un string con 'numerador/denominador'
func (fraccion Fraccion) Representacion() string {
	nuevoString := fmt.Sprintf("%d/%d", fraccion.numerador, fraccion.denominador)
	return nuevoString
}

// Crear una instancia de fraccion
func CrearFraccion(numerador, denominador int) Fraccion {
	return Fraccion{numerador: numerador, denominador: denominador}
}
