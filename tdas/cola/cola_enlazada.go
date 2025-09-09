package cola

type Nodo[T any] struct {
	dato      T
	siguiente *Nodo[T]
}

type ColaEnlazada[T any] struct {
	primero *Nodo[T]
	ultimo  *Nodo[T]
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
func (c ColaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

func (c ColaEnlazada[T]) VerPrimero() T

func (c *ColaEnlazada[T]) Encolar(T)

func (c *ColaEnlazada[T]) Desencolar() T

func CrearColaEnlazada[T any]() Cola[T] {
	nuevaCola := ColaEnlazada[T]{primero: nil, ultimo: nil}
	return &nuevaCola
}
