package servidor

import diccionario "tdas/diccionario"

type Post struct {
	id            int
	contenido     string
	creador       string
	afinidad      int
	cantidadLikes int
	likes         diccionario.DiccionarioOrdenado[string, string]
}

// CrearPost (id, contenido, creador), toma nuevo ID de la publicacion, creador de publicacion,
// y contenido de la publicacion. Devuelve puntero al post creado
func CrearPost(nuevoId, afinidad int, creadorPost, contenido string) *Post {
	return &Post{
		id:        nuevoId,
		creador:   creadorPost,
		contenido: contenido,
		afinidad:  afinidad,
	}
}

// compararPost toma 2 posts y devuevle:
// num > 0 si post1 es menor a post2, 0 si son iguales, num < 0 si post1 es mayor a post2
func compararPost(post1, post2 *Post) int {
	return post2.afinidad - post1.afinidad
}
