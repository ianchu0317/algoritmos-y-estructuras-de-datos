package servidor

import (
	"fmt"
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
	// Complejidad O(1) -> asignación de variables y operaciones con hash de O(1)
	if servidor.sesion != nil {
		fmt.Println("Error: Ya habia un usuario loggeado")
		return
	}
	if !servidor.usuarios.Pertenece(nombre) {
		fmt.Println("Error: usuario no existente")
		return
	}
	servidor.sesion = servidor.usuarios.Obtener(nombre)
	fmt.Println("Hola", nombre)
}

func (servidor *AlgoGram) Logout() {
	// Complejidad O(1) -> asignacion de variables
	if servidor.sesion == nil {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	nombre := servidor.sesion.nombre
	servidor.sesion = nil
	fmt.Println("Adios", nombre)
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
