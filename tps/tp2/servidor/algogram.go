package servidor

import (
	Diccionario "tdas/diccionario"
)

type AlgoGram struct {
	sesion        *Usuario
	usuarios      Diccionario.Diccionario[string, *Usuario] // Hash de nombre -> usuario
	posts         Diccionario.Diccionario[int, *Post]       // Hash de id -> post
	ordenRegistro Diccionario.Diccionario[string, int]      // Hash de nombre -> orden de la lista (para calc afinidad)
	proximoPostId int
}

// CrearServidor devuelve un servidor de Algogram
func CrearServidor(usuarios []string) Servidor {
	servidor := AlgoGram{
		sesion:        nil,
		usuarios:      Diccionario.CrearHash[string, *Usuario](func(a, b string) bool { return a == b }),
		posts:         Diccionario.CrearHash[int, *Post](func(a, b int) bool { return a == b }),
		ordenRegistro: Diccionario.CrearHash[string, int](func(a, b string) bool { return a == b }),
		proximoPostId: 0,
	}
	servidor.registrarUsuarios(usuarios)
	return &servidor
}

// *** Primitivas Algogram ***

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

// *** Funciones auxiliares ***

func (servidor *AlgoGram) registrarUsuarios(usuarios []string) {
	// Para cada usuario en la lista:
	// - Crear usuario y guardar en hash al puntero
	// - Guardar orden de registro de la lista
	for i, nombre := range usuarios {
		usuario := CrearUsuario(nombre)
		servidor.usuarios.Guardar(nombre, usuario)
		servidor.ordenRegistro.Guardar(nombre, i)
	}
}
