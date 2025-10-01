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
	// Actualizar ultimo nodo al insertar primer elemento
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
	// Actualizar ultimo puntero si lista está vacia
	if lista.EstaVacia() {
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

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	nuevoIterador := iteradorLista[T]{lista, lista.primero, nil}
	return &nuevoIterador
}

// *** Declarar primitivas y estructuras de iterador externo ***

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodo[T]
	anterior *nodo[T]
}

func (iter iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter iteradorLista[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(elemento T) {
	nuevoNodo := nodo[T]{elemento, iter.actual}
	// caso primer elemento
	if iter.anterior != nil {
		iter.anterior.siguiente = &nuevoNodo
	} else {
		iter.lista.primero = &nuevoNodo
	}
	// caso insertar al final, actualizar ultimo
	if iter.actual == nil {
		iter.lista.ultimo = &nuevoNodo
	}
	iter.actual = &nuevoNodo
	iter.lista.largo++
}

func (iter *iteradorLista[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	// caso primer elemento
	if iter.anterior != nil {
		iter.anterior.siguiente = iter.actual.siguiente
	} else {
		iter.lista.primero = iter.actual.siguiente
	}
	nodoBorrado := iter.actual
	iter.actual = iter.actual.siguiente
	nodoBorrado.siguiente = nil // Desvincular nodo a borrar para garbage collector
	iter.lista.largo--
	// actualizar ultimo nodo de la lista (caso ultimo elemento)
	if !iter.HaySiguiente() {
		iter.lista.ultimo = iter.anterior
	}
	return nodoBorrado.dato
}
