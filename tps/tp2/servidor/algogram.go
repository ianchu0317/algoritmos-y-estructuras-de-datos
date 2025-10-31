package servidor

import (
	Diccionario "tdas/diccionario"
)

type AlgoGram struct {
	sesion        *Usuario
	usuarios      Diccionario.Diccionario[string, *Usuario] // Hash de nombre -> usuario
	posts         Diccionario.Diccionario[int, *Post]       // Hash de id -> post
	proximoPostId int
}

// CrearServidor devuelve un servidor de Algogram
func CrearServidor(usuarios []string) Servidor {
	servidor := AlgoGram{
		sesion:        nil,
		usuarios:      Diccionario.CrearHash[string, *Usuario](func(a, b string) bool { return a == b }),
		posts:         Diccionario.CrearHash[int, *Post](func(a, b int) bool { return a == b }),
		proximoPostId: 0,
	}
	return &servidor
}

// Primitivas Algogram

func (servidor *AlgoGram) Login(nombre string) {

}

func (servidor *AlgoGram) Logout() {

}

func (servidor *AlgoGram) Publicar() {

}

func (servidor *AlgoGram) VerProxFeed() {

}

func (servidor *AlgoGram) Likear(id int) {

}

func (servidor AlgoGram) MostrarLikes(id int) {

}
