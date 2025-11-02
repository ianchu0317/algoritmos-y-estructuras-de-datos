package servidor

import (
	"strings"
	Diccionario "tdas/diccionario"
)

type post struct {
	id        int
	contenido string
	creador   string
	likes     Diccionario.DiccionarioOrdenado[string, string]
}

type postEnFeed struct {
	afinidad int
	post     Post
}

// CrearPost (id, contenido, creador), toma nuevo ID de la publicacion, creador de publicacion,
// y contenido de la publicacion. Devuelve puntero al post creado
func crearPost(nuevoId int, creadorPost, contenido string) Post {
	return &post{
		id:        nuevoId,
		creador:   creadorPost,
		contenido: contenido,
		likes:     Diccionario.CrearABB[string, string](strings.Compare),
	}
}

// crearPostEnFeed toma afinidad del usuario creador con usuario del feed, y el puntero al post a guardar.
// Devuelve puntero del post a guardar en el feed.
func crearPostEnFeed(afinidad int, post Post) *postEnFeed {
	return &postEnFeed{
		afinidad: afinidad,
		post:     post,
	}
}

// Primitivas de Posts

func (post *post) AgregarLike(nombre string) {
	post.likes.Guardar(nombre, "")
}

func (post *post) ObtenerLikes() int {
	return post.likes.Cantidad()
}

func (post *post) IterarLikes(visitar func(nombre, _ string) bool) {
	post.likes.Iterar(visitar)
}
