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

// compararPost toma 2 posts del feed y devuevle 'num':
// num > 0 si post1 es menor a post2, 0 si son iguales, num < 0 si post1 es mayor a post2
func compararPost(post1, post2 *postEnFeed) int {
	if post2.afinidad == post1.afinidad {
		return post2.post.id - post1.post.id
	}
	return post2.afinidad - post1.afinidad
}
