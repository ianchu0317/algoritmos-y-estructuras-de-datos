package servidor

import (
	"strings"
	Diccionario "tdas/diccionario"
)

type Post struct {
	id            int
	contenido     string
	creador       string
	cantidadLikes int
	likes         Diccionario.DiccionarioOrdenado[string, string]
}

// CrearPost (id, contenido, creador), toma nuevo ID de la publicacion, creador de publicacion,
// y contenido de la publicacion. Devuelve puntero al post creado
func CrearPost(nuevoId int, creadorPost, contenido string) *Post {
	return &Post{
		id:            nuevoId,
		creador:       creadorPost,
		contenido:     contenido,
		cantidadLikes: 0,
		likes:         Diccionario.CrearABB[string, string](strings.Compare),
	}
}

type postEnFeed struct {
	afinidad int
	post     *Post
}

// crearPostEnFeed toma afinidad del usuario creador con usuario del feed, y el puntero al post a guardar.
// Devuelve puntero del post a guardar en el feed.
func crearPostEnFeed(afinidad int, post *Post) *postEnFeed {
	return &postEnFeed{
		afinidad: afinidad,
		post:     post,
	}
}

// compararPost toma 2 posts del feed y devuevle 'num':
// num > 0 si post1 es menor a post2, 0 si son iguales, num < 0 si post1 es mayor a post2
func compararPost(post1, post2 *postEnFeed) int {
	return post2.afinidad - post1.afinidad
}
