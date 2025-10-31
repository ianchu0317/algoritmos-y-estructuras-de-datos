package servidor

import diccionario "tdas/diccionario"

type Post struct {
	id            int
	contenido     string
	creador       string
	cantidadLikes int
	likes         diccionario.DiccionarioOrdenado[string, string]
}

// CrearPost (id, contenido, creador), toma nuevo ID de la publicacion, creador de publicacion,
// y contenido de la publicacion. Devuelve puntero al post creado
func CrearPost(nuevoId int, creadorPost, contenido string) *Post {
	return &Post{
		id:        nuevoId,
		creador:   creadorPost,
		contenido: contenido,
	}
}
