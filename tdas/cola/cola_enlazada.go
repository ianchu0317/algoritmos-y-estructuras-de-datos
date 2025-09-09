package cola

type Nodo[T any] struct {
	dato      T
	siguiente *Nodo[T]
}

type ColaEnlazada[T any] struct {
	primero *Nodo[T]
	ultimo  *Nodo[T]
}

func (cola ColaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil && cola.ultimo == nil
}

func (cola ColaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *ColaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := Nodo[T]{dato: elemento, siguiente: nil}
	if cola.EstaVacia() {
		cola.primero = &nuevoNodo
	} else {
		cola.ultimo.siguiente = &nuevoNodo
	}
	cola.ultimo = &nuevoNodo
}

func (cola *ColaEnlazada[T]) Desencolar() T {
	datoPrimerNodo := cola.VerPrimero()
	cola.primero = cola.primero.siguiente
	// actualizar ultimo puntero si esta vacio
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return datoPrimerNodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	nuevaCola := ColaEnlazada[T]{primero: nil, ultimo: nil}
	return &nuevaCola
}
