package parcialitos

/*
Implementar una primitiva Filter(f func(T) bool) para el Heap, que deje al heap únicamente con los elementos para
los que la función devuelva true. La primitiva debe funcionar en O(n), siendo n la cantidad de elementos inicialmente
en el heap. Por supuesto, luego de aplicar la operación, el heap debe quedar en un estado válido para poder seguir
operando. Justificar la complejidad de la primitiva implementada.*/

type Heap[T any] struct {
	arreglo  []T
	cantidad int
}

func (heap *Heap[T]) Filter(f func(T) bool) {
	// Inicializacion de variables O(1)
	aux := make([]T, len(heap.arreglo))
	cantidadAux := 0

	// para los n-elementos del heap, agregar a arreglo auxiliar si pasa filtro
	for i := 0; i < heap.cantidad; i++ {
		elem := heap.arreglo[i]
		if f(elem) {
			aux[cantidadAux] = elem
			cantidadAux++
		}
	}
	// Actualizar variabels O(1)
	heap.arreglo = aux
	heap.cantidad = cantidadAux

	// O(n) se considera ya implementada
	heapify(heap.arreglo, heap.cantidad, heap.cmp) // Le paso arreglo, cantidad de arreglo, funcion comparacion
	heap.chequearRedimension()                     // Chequear redimension del arreglo si hace falta / Sino se chequea en proximo encolar/desencolar
}

/*
Implementar una función eliminarRepetidos(arreglo []int) []int que dado un arreglo de números, nos devuelva
otro en el que estén los elementos del original sin repetidos. La primera aparición debe mantenerse, y las demás no ser
consideradas. Indicar y justificar la complejidad del algoritmo implementado.
*/

func eliminarRepetidos(arreglo []int) []int {
	// Crear Hash O(1)
	visitados := CrearHash[int, bool](func(a, b int) bool { return a == b })
	sinRepetidos := make([]int, 0, len(arreglo))
	// Calcular para elementos Guardar cada elemento como no visitado O(n) -> N veces operaciones O(1)
	for _, num := range arreglo {
		if !visitados.Pertenece(num) {
			sinRepetidos = append(sinRepetidos, num)
			visitados.Guardar(num, false)
		}
	}
	return sinRepetidos
}

// Complejiadad total: O(n) = O(N)

/*
Se tiene un árbol binario que representa la fase eliminatoria del mundial. En cada nodo guarda el nombre del país, así
como la cantidad de goles que convirtió en dicha fase (incluyendo la tanda de penales, si fuera necesario). El padre
del nodo debe si o si tener al hijo que ganó (tuvo mayor cantidad de goles). Implementar una primitiva para el árbol
donde solamente están los nombres de los equipos en las hojas (no en los internos), y deje el árbol completado con los
ganadores en cada fase. Se puede asumir que el árbol es o bien completo, o que al menos todos los nodos internos tienen
exactamente 2 hijos. La cantidad de goles en la raíz no es relevante. La estructura del árbol es:
type Arbol struct {
pais string
goles int
izq *Arbol
der *Arbol
}
Tomando el ejempo del dorso, si invocamos para el árbol de la izquierda, debe quedar como el de la derecha.
*/

type Arbol struct {
	pais  string
	goles int
	izq   *Arbol
	der   *Arbol
}

func (ab *Arbol) HallarGanador() {
	if ab.izq == nil && ab.der == nil {
		return
	}
	// Orden post-order
	ab.izq.HallarGanador()
	ab.der.HallarGanador()

	// Actualizar nodo actual
	if ab.izq.goles > ab.der.goles {
		ab.pais = ab.izq.pais
	} else {
		ab.pais = ab.der.pais
	}
}

// Complejidad por teorema maestro O(n) -> visito todos los nodos chequeando ganador
