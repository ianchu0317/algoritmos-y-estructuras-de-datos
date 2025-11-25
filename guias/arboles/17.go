package arboles

/*
Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio y una clave final fin,
se devuelva una lista con todos los datos cuyas claves estén entre inicio y fin,
que estén dentro de los primeros M niveles del árbol (considerando a la raíz en nivel 1).
Indicar y justificar la complejidad temporal.
*/
func (raiz *abb[int]) obtenerValores(m, inicio, fin int) Lista[int] {
	resultado := CrearListaEnlazada[int]()
	raiz._obtenerValores(m, inicio, fin, 1, resultado)
	return resultado
}

func (nodo *abb[T]) _obtenerValores(m, ini, fin, nivel int, res Lista[int]) {
	if nodo == nil {
		return
	}
	if nivel > m {
		return
	}
	// Si nodo actual mayor al inicio puedo ir izquierda
	// Si nodo actual menor al final puedo ir derecha
	if nodo.valor > ini {
		nodo.izq._obtenerValores(m, ini, fin, nivel+1, res)
	}
	if nodo.valor >= ini && nodo.valor <= fin {
		res.AgregarUltimo(nodo.valor)
	}
	if nodo.valor < fin {
		nodo.der._obtenerValores(m, ini, fin, nivel+1, res)
	}
}
