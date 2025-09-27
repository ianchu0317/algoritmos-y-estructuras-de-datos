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
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := nodo[T]{dato: elemento, siguiente: nil}
	lista.ultimo.siguiente = &nuevoNodo
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

}

func (lista listaEnlazada[T]) VerPrimero() T {

}

func (lista listaEnlazada[T]) VerUltimo() T {

}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {

}

func (lista listaEnlazada[T]) Iterador() IteradorLista[T] {

}

// *** Declarar primitivas y estructuras de iterador externo ***
type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
}

func (iter iteradorLista[T]) HaySiguiente() bool {

}

func (iter iteradorLista[T]) Siguiente() {

}

func (iter iteradorLista[T]) Insertar(elemento T) {

}

func (iter iteradorLista[T]) Borrar() T {

}
