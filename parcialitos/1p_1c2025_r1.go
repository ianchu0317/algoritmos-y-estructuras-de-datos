package parcialitos

import TDAPila "tdas/pila"

/*
Ejercicio 1:

Se tiene una cadena que contiene () y ningún otro caracter (considerar que un único caracter es de tipo rune). Un
ejercicio típico es dada una cadena averiguar si está balanceada (es decir, todos los símbolos de apertura se cierran, y
además respetan el orden en el que se abrieron. Ejemplos balanceados: "()()()", o "(())()". No balanceados: "(()",
o ")(".
Teniendo en cuenta esto, se tiene una cadena que se asegura que en caso de borrar algunos paréntesis la cadena será
balanceada, se pide implementar una función func cantBorradosBalanceada(cadena string) int que dada una
cadena de este tipo, devuelva la cantidad mínima de paréntesis que se deben borrar para que la cadena esté balanceada.
Indicar y justificar la complejidad del algoritmo.
Ejemplos:
cadena: '()' -> 0
cadena: ')(' -> 2
cadena: '(()' -> 1
cadena: ')(()' -> 2
*/

func cantBorradosBalanceada(cadena string) int {
	// INicializar variables O(1)
	pila := TDAPila.CrearPilaDinamica[string]()
	// Visito cada parentesis y realizo operaciones O(1) -> O(n)
	for _, _letra := range cadena {
		letra := string(_letra)
		if string(letra) == "(" {
			pila.Apilar(letra)
		} else {
			if !pila.EstaVacia() && pila.VerTope() == "(" {
				pila.Desapilar()
			} else {
				pila.Apilar(letra)
			}
		}
	}
	// Contar elementos restantes de pila O(n) -> en el peor caso esta todo desbalanceado
	contador := 0
	for !pila.EstaVacia() {
		contador++
		pila.Desapilar()
	}
	return contador
	// Complejidad total O(n)
}

/*
Ejercicio 2:

Implementar una función func esCuadradoPerfecto(n int) bool que por División y Conquista determine si el
número n (un positivo entero) es un cuadrado perfecto. Un número es cuadrado perfecto si existe un número entero x
tal que x² = n.
Indicar y justificar la complejidad del algoritmo utilizando el Teorema Maestro.
*/

func esCuadradoPerfecto(n int) bool {
	if n < 0 {
		return false
	} else if n == 0 || n == 1 {
		return true
	}
	return _esCuadradoPerfecto(n, 1, n)
}

func _esCuadradoPerfecto(n, a, b int) bool {
	// Teorema maestro T(n) = T(n/2) + O(1)
	// A = 1, B = 2, C = 0 -> log2(1) == 0 -> O(n⁰log(n)) -> O(log(n))

	// Comparaciones y variables O(1)
	if a > b {
		return false
	}
	x := (a + b) / 2
	// LLamada recursiva con mitad de conjunto de enteros (matematicamente hablando)
	if x*x == n {
		return true
	} else if x*x > n {
		return _esCuadradoPerfecto(n, a, x-1)
	} else {
		return _esCuadradoPerfecto(n, x+1, b)
	}
}
