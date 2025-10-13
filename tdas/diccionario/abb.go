package diccionario

import (
	TDAPila "tdas/pila"
)

// Constantes
const (
	BUSQUEDA_IZQ = -1
	BUSQUEDA_DER = 1
)

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

// buscarNodoReemplazo toma el nodo actual y busca el candidato para su sucesion (dependiendo de direccion).
// Devuelve puntero al puntero del nodo encontrado.
func (abb arbolBinario[K, V]) buscarNodoReemplazo(nodoActual **nodo[K, V], direccion int) **nodo[K, V] {
	// Mover hacia derecha o izquierda dependiendo si existe o no
	// Mover hacia el final segun direccion

	// Busqueda en lado derecho
	if direccion == BUSQUEDA_DER {
		if (**nodoActual).der == nil {
			return nodoActual
		}
		return abb.buscarNodoReemplazo(&((**nodoActual).der), direccion)
	}
	// Busqueda en lado izquierdo
	if (**nodoActual).izq == nil {
		return nodoActual
	}
	return abb.buscarNodoReemplazo(&((**nodoActual).izq), direccion)
}

// Primitivas de ABB

func (abb *arbolBinario[K, V]) Guardar(clave K, dato V) {
	nodo := abb.buscarNodo(clave, &abb.raiz)
	// Si nodo esta vacio, crear uno nuevo e insertar en su posicion
	if *nodo == nil {
		nuevoNodo := crearNodo(clave, dato)
		*nodo = nuevoNodo // que el puntero apunte a ese nodo
		nodo = &nuevoNodo // avanzar el punter a ese nodo para editar dato
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
	nodo := abb.buscarNodo(clave, &abb.raiz)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := (**nodo).dato
	// Si no tiene hijos entonces, solo borrar el nodo (desenlazar de estructura)
	// Si tiene hijos (ver cual hijo tiene) -> buscar reemplazo
	if (**nodo).izq == nil && (**nodo).der == nil {
		*nodo = nil
	} else {
		nodoReemplazo := nodo
		if (**nodo).izq != nil {
			// Si tiene hijo izquierdo entonces aplicar "busqueda maximo de minimos"
			// Mover uno izquierda y luego full derecha
			nodoReemplazo = &((**nodoReemplazo).izq)
			nodoReemplazo = abb.buscarNodoReemplazo(nodoReemplazo, BUSQUEDA_DER)
		} else {
			// Si no tiene hijo izquierdo pero sí derecho entonces aplicar "busqueda minimo de maximos"
			// Mover uno derecha y luego full izquierda
			nodoReemplazo = &((**nodoReemplazo).der)
			nodoReemplazo = abb.buscarNodoReemplazo(nodoReemplazo, BUSQUEDA_IZQ)
		}
		// Copiar datos de heredero a nodo a borrar
		(**nodo).clave, (**nodo).dato = (**nodoReemplazo).clave, (**nodoReemplazo).dato
		// Borrar el nodo de reemplazo dependiendo si tiene hijos o no
		if (**nodoReemplazo).izq != nil || (**nodoReemplazo).der != nil {
			// Si tiene hijo izquierdo entonces apuntar a hijo izquierdo, sino al derecho
			if (**nodoReemplazo).izq != nil {
				*nodoReemplazo = (**nodoReemplazo).izq
			} else {
				*nodoReemplazo = (**nodoReemplazo).der
			}
		} else {
			*nodoReemplazo = nil
		}
	}
	abb.cantidad--
	return dato
}

func (abb arbolBinario[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb arbolBinario[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	// Iterar in-Order
	abb.raiz.iterarRango(nil, nil, abb.comparar, visitar)
}

// nodo.iterar() itera los nodos del subarbol del nodo que se utiliza para llamar el metodo.
// Devuelve true si se quiere seguir iterando, devuelve false si se quiere dejar de iterar.
func (nodo *nodo[K, V]) iterarRango(desde, hasta *K,
	comparar func(K, K) int, visitar func(clave K, dato V) bool) bool {
	// Caso base si es final de arbol, volver al anterior nodo pero seguir iterando
	if nodo == nil {
		return true
	}
	// Iterar in-order: izq, nodo, der si cumple condicion.
	// Si cumple condicion:  desde <= nodo.clave <= hasta entonces hacer inorder
	// Si nodo.clave < desde, ir hacia derecha
	// Si nodo.clave > final, ir hacia izquierda
	// La unica forma de devolver false es con salida de funcion visitar.
	if (desde == nil || comparar(nodo.clave, *desde) >= 0) && (hasta == nil || comparar(nodo.clave, *hasta) <= 0) {
		if !nodo.izq.iterarRango(desde, hasta, comparar, visitar) {
			return false
		}
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
		return nodo.der.iterarRango(desde, hasta, comparar, visitar)
	} else {
		if desde != nil && comparar(nodo.clave, *desde) < 0 {
			return nodo.der.iterarRango(desde, hasta, comparar, visitar)
		}
		if hasta != nil && comparar(nodo.clave, *hasta) > 0 {
			return nodo.izq.iterarRango(desde, hasta, comparar, visitar)
		}
	}
	return true
}

func (abb arbolBinario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(desde, hasta, abb.comparar, visitar)
}

func (abb arbolBinario[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIter := iteradorABB[K, V]{
		pilaNodos: TDAPila.CrearPilaDinamica[*nodo[K, V]](),
		desde:     nil,
		hasta:     nil,
		comparar:  abb.comparar}
	nuevoIter.apilarMinimos(abb.raiz)
	return &nuevoIter
}

func (abb arbolBinario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	nuevoIter := iteradorABB[K, V]{
		pilaNodos: TDAPila.CrearPilaDinamica[*nodo[K, V]](),
		desde:     desde,
		hasta:     hasta,
		comparar:  abb.comparar}
	nuevoIter.apilarMinimos(abb.raiz)
	return &nuevoIter
}

// *** Estructura Iterador Externo ABB ***
type iteradorABB[K any, V any] struct {
	pilaNodos TDAPila.Pila[*nodo[K, V]]
	desde     *K
	hasta     *K
	comparar  func(K, K) int
}

// Funciones auxiliares Iter Externo

// apilarIzquierda() toma una un nodo.
// Apila a la pila interna del iterador todos los nodos izquierdos
// al nodo actual inclusive (si no es nil)
func (iter *iteradorABB[K, V]) apilarMinimos(nodo *nodo[K, V]) {
	if nodo == nil {
		return
	}
	// Si el nodo actual cumple:
	// - desde <= nodo.clave <= hasta
	// - O iter== nil y hasta == nil
	// Entonces apilar normal todo izquierdo
	if (iter.desde == nil || iter.comparar(nodo.clave, *iter.desde) >= 0) &&
		(iter.hasta == nil || iter.comparar(nodo.clave, *iter.hasta) <= 0) {
		iter.pilaNodos.Apilar(nodo)
		iter.apilarMinimos(nodo.izq)
	}
	// Si nodo actual cumple nodo.clave < desde entonces considerar hijos derechos
	if iter.desde != nil && iter.comparar(nodo.clave, *iter.desde) < 0 {
		iter.apilarMinimos(nodo.der)
	}
	// Si nodo actual cumple nodo.clave > hasta entonces considerar ijos izquierdo
	if iter.hasta != nil && iter.comparar(nodo.clave, *iter.hasta) > 0 {
		iter.apilarMinimos(nodo.izq)
	}
}

// Primitivas Iterador Externo

func (iter iteradorABB[K, V]) HaySiguiente() bool {
	// Chequear que la pila no esté vacia
	return !iter.pilaNodos.EstaVacia()
}

func (iter iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pilaNodos.VerTope().clave, iter.pilaNodos.VerTope().dato
}

func (iter iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodoTope := iter.pilaNodos.Desapilar()
	iter.apilarMinimos(nodoTope.der)
}
