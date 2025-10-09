package diccionario

import "fmt"

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
	return &nodo[K, V]{clave: clave, dato: dato, izq: nil, der: nil}
}

// CrearABB crea una instancia de diccionario ordenado
func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &arbolBinario[K, V]{nil, funcionCmp, 0}
}

// Funciones auxiliares

// buscarNodo toma una clave y devuelve la posicion del nodo correspondiente.
// En caso de existir devuelve el puntero al puntero del nodo
// En caso de no existir devuelve nil
func (abb arbolBinario[K, V]) buscarNodo(clave K, nodoActual **nodo[K, V]) **nodo[K, V] {
	// Ir por cada nodo hasta encontrar clave igual o si nodo actual es nil
	if *nodoActual == nil || abb.comparar(clave, (**nodoActual).clave) == 0 {
		return nodoActual
	}
	// Si clave a buscar es menor a clave del nodo actual ir a izquierda
	// Sino ir a nodo derecha
	if abb.comparar(clave, (**nodoActual).clave) < 0 {
		return abb.buscarNodo(clave, &((**nodoActual).izq))
	}
	return abb.buscarNodo(clave, &((**nodoActual).der))
}

// esHoja devuelve toma un nodo y devuelve True si es hoja del arbol (si no tiene hijos),
// de lo contrario devuelve false
func esHoja[K any, V any](nodoActual *nodo[K, V]) bool {
	return nodoActual.izq == nil && nodoActual.der == nil
}

// buscarNodoReemplazo toma el nodo actual y busca el candidato para su sucesion.
// Devuelve puntero al puntero del nodo encontrado
func (abb *arbolBinario[K, V]) buscarNodoReemplazo(nodoActual **nodo[K, V]) **nodo[K, V] {
	// Mover hacia derecha o izquierda dependiendo si existe o no
	if *nodoActual == nil {
		return nodoActual
	}
	// Mover hasta nodo final muy derecho
	return abb.buscarNodoReemplazo(&((**nodoActual).der))
}

// Primitivas de ABB

func (abb *arbolBinario[K, V]) Guardar(clave K, dato V) {
	nodo := abb.buscarNodo(clave, &abb.raiz)
	// Si nodo esta vacio, crear uno nuevo e insertar en su posicion
	if *nodo == nil {
		nuevoNodo := crearNodo(clave, dato)
		*nodo = nuevoNodo // que el puntero apunte a ese nodo
		nodo = &nuevoNodo // avanzar el punter a ese nodo para editar dato
		fmt.Println("clave insertado es:", (**nodo).clave)
		abb.cantidad++
	} else {
		(**nodo).dato = dato
	}
}

func (abb arbolBinario[K, V]) Pertenece(clave K) bool {
	nodo := abb.buscarNodo(clave, &abb.raiz)
	return *nodo != nil
}

func (abb arbolBinario[K, V]) Obtener(clave K) V {
	nodo := abb.buscarNodo(clave, &abb.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (**nodo).dato
}

func (abb *arbolBinario[K, V]) Borrar(clave K) V {
	// Si no tiene hijos (es hoja), entonces solo borrar
	// Si tiene hijos (ver cual hijo tiene) -> buscar reemplazo
	nodo := abb.buscarNodo(clave, &abb.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (**nodo).dato
	if !esHoja(*nodo) {
		// mover uno izquierda
		nodoReemplazo := &((**nodo).izq)
		nodoReemplazo = abb.buscarNodoReemplazo(nodoReemplazo)
		// Cambiar y copiar clave
		(**nodo).clave, (**nodo).dato = (**nodoReemplazo).clave, (**nodoReemplazo).dato
		// borrar nodo reemplazo
		*nodoReemplazo = nil
	} else {
		*nodo = nil
	}
	return dato
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
