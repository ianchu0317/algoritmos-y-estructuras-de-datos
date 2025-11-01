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
	if servidor.hayLoggeado() {
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
	if !servidor.hayLoggeado() {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	nombre := servidor.sesion.nombre
	servidor.sesion = nil
	fmt.Println("Adios", nombre)
}

func (servidor *AlgoGram) Publicar(contenido string) {
	if !servidor.hayLoggeado() {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	// Complejidad: O(1) + O(u * log(posts)) + O(1)
	nuevoPost := CrearPost(servidor.proximoPostId, servidor.sesion.nombre, contenido)

	for iter := servidor.usuarios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombreActual, usuarioActual := iter.VerActual()
		if nombreActual != servidor.sesion.nombre {
			afinidad := servidor.calcularAfinidad(nombreActual, servidor.sesion.nombre)
			postFeed := crearPostEnFeed(afinidad, nuevoPost)
			usuarioActual.feed.Encolar(postFeed)
		}
	}

	servidor.posts.Guardar(servidor.proximoPostId, nuevoPost)
	servidor.proximoPostId++
	fmt.Println("Post publicado")
}

func (servidor *AlgoGram) VerProxFeed() {
	if !servidor.hayLoggeado() || servidor.sesion.feed.EstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}
	// Complejidad O(log(posts))
	siguientePost := servidor.sesion.feed.Desencolar().post
	fmt.Println("Post ID", siguientePost.id)
	fmt.Println(siguientePost.creador, "dijo:", siguientePost.contenido)
}

func (servidor *AlgoGram) Likear(id int) {
	if !servidor.hayLoggeado() || !servidor.posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	// Chequear errores
	// Obtener el post indicado -> actualizar su valor de like
	// Agregar a ABB -> actualizar cantidad de likes
	post := servidor.posts.Obtener(id)
	post.likes.Guardar(servidor.sesion.nombre, "")
	post.cantidadLikes++
	fmt.Println("Post likeado")
}

func (servidor AlgoGram) MostrarLikes(id int) {

}

// *** Funciones auxiliares ***

// hayLoggeado devuelve True si hay algun usuario en la sesion actual, false en caso contrario
func (servidor AlgoGram) hayLoggeado() bool {
	return servidor.sesion != nil
}

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

// calcularAfinidad(u1, u2) Toma dos nombres de usuarios y devuelve su afinidad según orden de registro.
func (servidor AlgoGram) calcularAfinidad(usuario1, usuario2 string) int {
	ordenUsuario1 := servidor.ordenRegistro.Obtener(usuario1)
	ordenUsuario2 := servidor.ordenRegistro.Obtener(usuario2)
	return abs(ordenUsuario1 - ordenUsuario2)
}

// abs(num) toma un numero y devuelve el valor absoluto (modulo)
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
