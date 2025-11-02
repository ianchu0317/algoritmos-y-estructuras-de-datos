package servidor

import (
	"strings"
	Diccionario "tdas/diccionario"
)

type Post struct {
	id        int
	contenido string
	creador   string
	likes     Diccionario.DiccionarioOrdenado[string, string]
}

type postEnFeed struct {
	afinidad int
	post     *Post
}

// CrearPost (id, contenido, creador), toma nuevo ID de la publicacion, creador de publicacion,
// y contenido de la publicacion. Devuelve puntero al post creado
func CrearPost(nuevoId int, creadorPost, contenido string) *Post {
	return &Post{
		id:        nuevoId,
		creador:   creadorPost,
		contenido: contenido,
		likes:     Diccionario.CrearABB[string, string](strings.Compare),
	}
}

// crearPostEnFeed toma afinidad del usuario creador con usuario del feed, y el puntero al post a guardar.
// Devuelve puntero del post a guardar en el feed.
func crearPostEnFeed(afinidad int, post *Post) *postEnFeed {
	return &postEnFeed{
		afinidad: afinidad,
		post:     post,
	}
}
