package cola

// **** Declaracion de estructuras a utilizar ****

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primerNodo *nodo[T]
	ultimoNodo *nodo[T]
}

// **** Funciones primitivas ****

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primerNodo == nil && cola.ultimoNodo == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primerNodo.dato
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nuevonodo := nodo[T]{dato: elemento, siguiente: nil}
	if cola.EstaVacia() {
		cola.primerNodo = &nuevonodo
	} else {
		cola.ultimoNodo.siguiente = &nuevonodo
	}
	cola.ultimoNodo = &nuevonodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	datoPrimernodo := cola.VerPrimero()
	cola.primerNodo = cola.primerNodo.siguiente
	// Actualizar ultimo puntero si no hay mas elementos
	if cola.primerNodo == nil {
		cola.ultimoNodo = nil
	}
	return datoPrimernodo
}

// CrearColaEnlazada devuelve una instancia de Cola
func CrearColaEnlazada[T any]() Cola[T] {
	nuevaCola := colaEnlazada[T]{primerNodo: nil, ultimoNodo: nil}
	return &nuevaCola
}
