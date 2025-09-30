package lista

// *** Declarar estructuras para Lista Enlazada ***

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	nuevaLista := listaEnlazada[T]{nil, nil, 0}
	return &nuevaLista
}

// Primitivas Lista Enlazada

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := nodo[T]{dato: elemento, siguiente: lista.primero}
	lista.primero = &nuevoNodo
	lista.largo++
	if lista.largo == 1 {
		lista.ultimo = &nuevoNodo
	}
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := nodo[T]{dato: elemento, siguiente: nil}
	if lista.ultimo != nil {
		lista.ultimo.siguiente = &nuevoNodo
	}
	lista.ultimo = &nuevoNodo
	lista.largo++
	if lista.largo == 1 {
		lista.primero = lista.ultimo
	}
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	primerElemento := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	// Actualizar ultimo puntero si hay un solo elemento
	if lista.primero == nil {
		lista.ultimo = nil
	}
	return primerElemento
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	nodoActual := lista.primero
	for nodoActual != nil && visitar(nodoActual.dato) {
		nodoActual = nodoActual.siguiente
	}
}

func (lista listaEnlazada[T]) Iterador() IteradorLista[T] {
	nuevoIterador := iteradorLista[T]{lista.primero, nil}
	return &nuevoIterador
}

// *** Declarar primitivas y estructuras de iterador externo ***

type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
}

func (iter iteradorLista[T]) VerActual() T {
	return iter.actual.dato
}

func (iter iteradorLista[T]) HaySiguiente() bool {
	return iter.actual == nil
}

func (iter *iteradorLista[T]) Siguiente() {
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	nuevoNodo := nodo[T]{elemento, iter.actual}
	iter.anterior.siguiente = &nuevoNodo
	iter.actual = &nuevoNodo
}

func (iter *iteradorLista[T]) Borrar() T {
	iter.anterior.siguiente = iter.actual.siguiente
	nodoViejo := iter.actual
	iter.actual = iter.actual.siguiente
	nodoViejo.siguiente = nil // Desvincular nodo a borrar para garbage collector
	return nodoViejo.dato
}
