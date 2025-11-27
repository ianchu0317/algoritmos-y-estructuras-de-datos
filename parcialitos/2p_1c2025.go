package parcialitos

/*
Ejercicio 1
Implementar una primitiva para el ABB func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K que reciba
2 claves y devuelva el último ancestro en común entre ambas claves. Dicho ancestro en común podría ser incluso alguna
de estas claves. Si alguna clave pasada no se encuentra en el árbol, finalizar con panic. Indicar y justificar la complejidad
de la primitiva implementada.
Mostramos ejemplos de resultados esperados de invocar la primitiva al árbol del dorso:
arbol.AncestroComun(1, 4) --> 2
arbol.AncestroComun(2, 4) --> 2
arbol.AncestroComun(9, 1) --> 5
*/

type abb[K any, V any] struct {
	clave K
	valor V
	izq   *abb[K, V]
	der   *abb[K, V]
	cmp   func(K, K) int
}

func (arbol *abb[K, V]) AncestroComun(clave1, clave2 K) K {
	if !arbol.Pertenece(clave1) || !arbol.Pertenece(clave2) {
		panic("No hay clave en el arbol")
	}
	var _max K
	var _min K
	if arbol.cmp(clave1, clave2) > 0 {
		_max, _min = clave1, clave2
	} else {
		_max, _min = clave2, clave1
	}
	return arbol._ancestroComun(_min, _max)
}

func (arbol *abb[K, V]) _ancestroComun(min, max K) K {
	if arbol.cmp(arbol.clave, min) == 0 {
		return min
	} else if arbol.cmp(arbol.clave, max) == 0 {
		return max
	}
	cmp1 := arbol.cmp(arbol.clave, min)
	cmp2 := arbol.cmp(arbol.clave, max)

	// Si es el nodo actual es mayor a izquierdo y menor a derecho
	// devolver nodo actual
	if cmp1 > 0 && cmp2 < 0 {
		return arbol.clave
	}
	// Si nodo actual es mayor a izq y derecho -> ir izquierda
	// Si nodo actual es menor a izq y derecha -> ir derecha
	if cmp1 > 0 && cmp2 > 0 {
		return arbol.izq._ancestroComun(min, max)
	} else {
		return arbol.der._ancestroComun(min, max)
	}
}
