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

func (c *ColaEnlazada[T]) Encolar(T) {

}

func (c *ColaEnlazada[T]) Desencolar() T {
	return c.primero.dato
}

func CrearColaEnlazada[T any]() Cola[T] {
	nuevaCola := ColaEnlazada[T]{primero: nil, ultimo: nil}
	return &nuevaCola
}
