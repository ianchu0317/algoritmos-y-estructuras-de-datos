package servidor

import Heap "tdas/cola_prioridad"

type Usuario struct {
	nombre string
	feed   Heap.ColaPrioridad[*Post]
}

// CrearUsuario(nombre) toma un nombre y devuelve un puntero a usuario
func CrearUsuario(nombre string) *Usuario {
	return &Usuario{
		nombre:nombre,
		feed: Heap.CrearHeap[*Post]()
	}
}

