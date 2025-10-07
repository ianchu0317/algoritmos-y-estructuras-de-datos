package diccionario

// *** Estructura de ABB ***

type nodo[K any, V any] struct {
	clave K
	dato  V
	izq   *nodo[K, V]
	der   *nodo[K, V]
}

type arbolBinario[K any, V any] struct {
	raiz     *nodo[K, V]
	comparar func(K, K) int
	cantidad int
}

// crearNodo crea un nodo para el abb
func crearNodo[K any, V any](clave K, dato V) *nodo[K, V] {
	return &nodo[K, V]{clave: clave, izq: nil, der: nil}
}

// CrearABB crea una instancia de diccionario ordenado
func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &arbolBinario[K, V]{nil, funcionCmp, 0}
}

// Funciones auxiliares

// buscarNodo toma una clave y devuelve la posicion del nodo correspondiente.
// En caso de existir devuelve la posicion
// En caso de no existir devuelve nil
func (abb arbolBinario[K, V]) buscarNodo(clave K) **nodo[K, V] {
	nodoActual := &abb.raiz
	// Ir por cada nodo hasta encontrar clave igual o si nodo actual es nil
	for *nodoActual != nil && abb.comparar(clave, (**nodoActual).clave) == 0 {
		// Si clave a buscar es menor a clave del nodo actual ir a izquierda
		// Sino ir a nodo derecha
		if abb.comparar(clave, (**nodoActual).clave) < 0 {
			nodoActual = &((**nodoActual).izq)
		} else {
			nodoActual = &((**nodoActual).der)
		}
	}
	return nodoActual
}

// Primitivas de ABB
func (abb arbolBinario[K, V]) Guardar(clave K, dato V) {
	nodo := abb.buscarNodo(clave)
	// Si nodo esta vacio, crear uno nuevo e insertar en su posicion
	if *nodo == nil {
		nuevoNodo := crearNodo(clave, dato)
		*nodo = nuevoNodo // que el puntero apunte a ese nodo
		nodo = &nuevoNodo // avanzar el punter a ese nodo para editar dato
		abb.cantidad++
	}
	(**nodo).dato = dato
}

func (abb arbolBinario[K, V]) Pertenece(clave K) bool {
	nodo := abb.buscarNodo(clave)
	return nodo != nil
}

func (abb arbolBinario[K, V]) Obtener(clave K) V {
	var data V
	return data
}

func (abb arbolBinario[K, V]) Borrar(clave K) V {
	var data V
	return data
}

func (abb arbolBinario[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb arbolBinario[K, V]) Iterar(visitar func(clave K, dato V) bool) {
}

func (abb arbolBinario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (abb arbolBinario[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorABB[K, V]{nil, nil}
}

func (abb arbolBinario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return &iteradorABB[K, V]{desde, hasta}
}

// *** Estructura Iterador Externo ABB ***
type iteradorABB[K any, V any] struct {
	desde *K
	hasta *K
}

func (iter iteradorABB[K, V]) HaySiguiente() bool {
	return false
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	var clave K
	var dato V
	return clave, dato
}

func (iter *iteradorABB[K, V]) Siguiente() {

}
