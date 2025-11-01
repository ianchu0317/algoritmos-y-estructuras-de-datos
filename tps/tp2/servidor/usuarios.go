package servidor

import Heap "tdas/cola_prioridad"

type Usuario struct {
	nombre string
	feed   Heap.ColaPrioridad[*postEnFeed]
}

// CrearUsuario(nombre) toma un nombre y devuelve un puntero a usuario
func CrearUsuario(nombre string) *Usuario {
	return &Usuario{
		nombre: nombre,
		feed:   Heap.CrearHeap[*postEnFeed](compararPost),
	}
}
