package servidor

import Heap "tdas/cola_prioridad"

type usuario struct {
	nombre string
	feed   Heap.ColaPrioridad[*postEnFeed]
}

// CrearUsuario(nombre) toma un nombre y devuelve un puntero a usuario
func crearUsuario(nombre string) *usuario {
	return &usuario{
		nombre: nombre,
		feed:   Heap.CrearHeap[*postEnFeed](compararPost),
	}
}

// compararPost toma 2 posts del feed y devuevle 'num':
// num > 0 si post1 es menor a post2, num < 0 si post1 es mayor a post2.
// Si tienen misma afinidad, compara por id del post:
// num > 0 si post1 tiene menor id a post2, num < 0 si post1 tiene mayor id a post2.
func compararPost(post1, post2 *postEnFeed) int {
	if post2.afinidad == post1.afinidad {
		id2, _, _ := post2.post.ObtenerInformacion()
		id1, _, _ := post1.post.ObtenerInformacion()
		return id2 - id1
	}
	return post2.afinidad - post1.afinidad
}
