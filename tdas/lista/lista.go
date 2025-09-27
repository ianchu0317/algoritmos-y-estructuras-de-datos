package lista

// Declaraciones de interfaces

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un elemento en la primera posicion de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta un elemento en la ultima posicion de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra y devuelve el elemento que se encuentra en primera posicion de la lista.
	// Si está vacia, entra en pánico con mensaje "La lista esta vacia"
	BorrarPrimero() T

	// VerPrimero devuelve el elemento que se encuentra en primera posicion de la lista.
	// La lista queda intacta. Si está vacia, entra en pánico con mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el elemento que se encuentra en ultima posicion de la lista.
	// La lista queda intacta. Si esta vacia, entra en pánico con mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve un entero con la cantidad de elementos que hay en la lista.
	Largo() int

	// Iterar recorre la lista aplicando la funcion 'visitar func(T) bool'.
	// Termina la iteracion si no quedan elementos de lista por iterar,
	// o si la funcion visitar(T) devuelve false para un elemento de la lista.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo IteradorLista[T] asociada a la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual devuelve el elemento actual de la lista.
	// La lista queda intacta.
	VerActual() T

	// HaySiguiente devuelve verdadero si hay siguiente elemento para iterar, false en caso contrario.
	HaySiguiente() bool

	// Siguiente pasa al siguiente elemento en la lista. No devuelve nada.
	Siguiente()

	// Insertar agrega un elemento en la posicion actual de iterador en la lista asociada.
	Insertar(T)

	// Borrar elimina y devuelve el elemento actual del iterador en la lista asociada.
	Borrar() T
}
