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
	return p.datos[p.cantidad-1]
}

// Apilar agrega un nuevo elemento a la pila.
func (p *pilaDinamica[T]) Apilar(elemento T) {
	// Redimensionar slice si no alcanza
	if p.cantidad == len(p.datos) {
		redimensionarPila(p, p.cantidad*2)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad += 1
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	ultElemento := p.VerTope()
	p.cantidad = p.cantidad - 1
	// Redimensionar slice si solo uso 25%
	if 4*p.cantidad == len(p.datos) {
		redimensionarPila(p, len(p.datos)/2)
	}
	return ultElemento
}

// redimensionarPila cambia los tamanios de los slices pila
func redimensionarPila[T any](pila *pilaDinamica[T], largo int) {
	nuevosDatos := make([]T, largo)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}

// Devuelve una Pila
func CrearPilaDinamica[T any]() Pila[T] {
	nuevaPilaDinamica := pilaDinamica[T]{}
	redimensionarPila(&nuevaPilaDinamica, 1)
	return &nuevaPilaDinamica
}
