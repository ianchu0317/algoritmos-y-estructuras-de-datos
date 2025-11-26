package hash

/*(★★) En un diccionario todas las claves tienen que ser diferentes, no así sus valores.
Escribir en Go una primitiva para el hash cerrado func (dicc *hashCerrado[K, V]) CantidadValoresDistintos(cmp func (V, V) bool) int
que, sin usar el iterador interno, dado un hash devuelva la cantidad de valores diferentes que almacena.
La función pasada por parámetro determina si dos valores son iguales, o no.
Indicar y justificar el orden del algoritmo.*/

// Iterar cada celda y si hay un valor q no esta en dict, guardar. Devlover cantidad total dict

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos(cmp func(V, V) bool) int {
	// Iterar para cada celda y contar la frecuencia de los valores
	hashValores := CrearHash[V, int](cmp)
	for _, celda := range dicc.tabla {
		if celda.estado == OCUPADO {
			if !hashValores.Pertenece(celda.valor) {
				hashValores.Guardar(celda.valor, 0)
			}
		}
	}
	return hashValores.Cantidad()
}
