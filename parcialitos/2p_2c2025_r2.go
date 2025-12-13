package parcialitos

/*1. Implementar la función InvertirAgrupando con la siguiente firma: func InvertirAgrupando(dicc Diccionario[K, V],
cmpValores func(V, V) bool) Diccionario[V, Lista[K]]
La función debe devolver un nuevo diccionario, cumpliendo las siguientes condiciones:
(i) Las claves del nuevo diccionario serán los valores V del diccionario original.
(ii) El valor para cada clave será una lista de las claves K originales que estaban asociadas a ese valor V (no importa el orden
en el que queden).
Ejemplo: Si el diccionario original es: {'Ana': 10, 'Beto': 20, 'Carla': 10, 'Dani': 20, 'Facu': 40}
El nuevo diccionario tendría las siguientes claves asociadas a las correspondientes listas: {10: ['Ana', 'Carla'], 20: ['Beto',
'Dani'], 40: ['Facu']}
Indicar y justificar la complejidad temporal del algoritmo implementado.
*/

func InvertirAgrupando(dicc Diccionario[K, V], cmpValores func(V, V) bool) Diccionario[V, Lista[K]] {
	// Para cada elemento si no está en invertido, entonces crear una lista nueva
	// Complejidad O(n)
	invertido := CerearHash[V, Lista[K]](cmpValores)
	dicc.Iterar(func(clave K, dato V) bool {
		if !invertido.Pertenece(dato) {
			invertido.Guardar(dato, CrearListaEnlazada[K]())
		}
		invertido.Obtener(dato).InsertarUltimo(clave)
		return true
	})
	return invertido
}

/*
2. Implementar una primitiva para el Heap con la siguiente firma: func (heap *Heap[T]) Map(f func(T) T)
La primitiva Map debe modificar el Heap actual, reemplazando a cada uno de los elementos originales por el correspondiente resultado
de la función f (pasada por parámetro). Importante: la primitiva Map debe asegurarse de que el Heap quede en un estado válido
para seguir operando.
Indicar y justificar la complejidad temporal de la primitiva implementada.
*/
func (heap *Heap[T]) Map(f func(T) T) {
	// Itero el heap aplicando para cada elemnto -> O(n)
	// Heapify O(n) -> Total O(n)
	for i := 0; i < heap.cantidad; i++ {
		heap.arreglo[i] = f(heap.arreglo[i])
	}
	heapify(heap.arreglo, heap.cantidad, heap.cmpValores)
}

/*
Se tiene un AVL cuyas claves son números (que se comparan de la forma tradicional), y cuyo estado inicial corresponde a tener
el siguiente recorrido Preorder: 9 - 6 - 7 - 19 - 10 - 21 - 22. Mostrar el estado del árbol. Luego, realizar el seguimiento de
insertar los siguientes elementos (incluyendo rotaciones intermedias): 25 - 20 - 5 - 3 - 17 - 1.
*/
