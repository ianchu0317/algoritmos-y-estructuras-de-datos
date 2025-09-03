package pila

// IMPLEMENTACIÓN DE PILA CON SLICES (DINÁMICOS)
// Definición del struct pila proporcionado por la cátedra.
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	// Devolver valor tope si no esta vacia
	return p.datos[p.cantidad]
}

// Apilar agrega un nuevo elemento a la pila.
func (p *pilaDinamica[T]) Apilar(elemento T) {

}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	ultElemento := p.datos[p.cantidad]
	p.cantidad = p.cantidad - 1
	return ultElemento
}

// Devuelve una Pila
func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{}
}
