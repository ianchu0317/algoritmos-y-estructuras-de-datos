package servidor

import (
	"fmt"
	Diccionario "tdas/diccionario"
)

type AlgoGram struct {
	sesion        Usuario
	usuarios      Diccionario.Diccionario[string, Usuario] // Hash de nombre -> usuario
	posts         Diccionario.Diccionario[int, Post]       // Hash de id -> post
	ordenRegistro Diccionario.Diccionario[string, int]     // Hash de nombre -> orden de la lista (para calc afinidad)
	proximoPostId int
}

// CrearServidor devuelve un servidor de Algogram
func CrearServidor(usuarios []string) Servidor {
	servidor := AlgoGram{
		sesion:        nil,
		usuarios:      Diccionario.CrearHash[string, Usuario](func(a, b string) bool { return a == b }),
		posts:         Diccionario.CrearHash[int, Post](func(a, b int) bool { return a == b }),
		ordenRegistro: Diccionario.CrearHash[string, int](func(a, b string) bool { return a == b }),
		proximoPostId: 0,
	}
	servidor.registrarUsuarios(usuarios)
	return &servidor
}

// *** Primitivas Algogram ***

func (servidor AlgoGram) HayLoggeado() bool {
	return servidor.sesion != nil
}

func (servidor *AlgoGram) Login(nombre string) {
	if servidor.HayLoggeado() {
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
	if !servidor.HayLoggeado() {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	servidor.sesion = nil
	fmt.Println("Adios")
}

func (servidor *AlgoGram) Publicar(contenido string) {
	if !servidor.HayLoggeado() {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	nombreLoggeado := servidor.sesion.ObtenerNombre()
	nuevoPost := crearPost(servidor.proximoPostId, nombreLoggeado, contenido)

	for iter := servidor.usuarios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombreActual, usuarioActual := iter.VerActual()
		if nombreActual != nombreLoggeado {
			afinidad := servidor.calcularAfinidad(nombreActual, nombreLoggeado)
			usuarioActual.ActualizarFeed(nuevoPost, afinidad)
		}
	}

	servidor.posts.Guardar(servidor.proximoPostId, nuevoPost)
	servidor.proximoPostId++
	fmt.Println("Post publicado")
}

func (servidor *AlgoGram) VerProxFeed() {
	if !servidor.HayLoggeado() || servidor.sesion.FeedEstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}
	proxPost := servidor.sesion.ObtenerProxPost()
	id, creador, contenido := proxPost.ObtenerInformacion()
	fmt.Println("Post ID", id)
	fmt.Println(creador, "dijo:", contenido)
	fmt.Println("Likes:", proxPost.ObtenerLikes())
}

func (servidor *AlgoGram) Likear(id int) {
	if !servidor.HayLoggeado() || !servidor.posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	// Complejidad O(log(likes))
	post := servidor.obtenerPostPorID(id)
	post.AgregarLike(servidor.sesion.ObtenerNombre())
	fmt.Println("Post likeado")
}

func (servidor AlgoGram) MostrarLikes(id int) {
	if !servidor.posts.Pertenece(id) {
		fmt.Println("Error: Post inexistente o sin likes")
		return
	}

	post := servidor.obtenerPostPorID(id)

	if !post.HayLikes() {
		fmt.Println("Error: Post inexistente o sin likes")
		return
	}

	// complejidad: O(likes)
	fmt.Println("El post tiene", post.ObtenerLikes(), "likes:")
	post.IterarLikes(func(nombre, _ string) bool {
		fmt.Printf("\t%s\n", nombre)
		return true
	})
}

// *** Funciones auxiliares ***

// registrarUsuarios toma una lista de nombres de usuarios y los registra en el servidor.
// Guarda en el servidor su orden de registro.
func (servidor *AlgoGram) registrarUsuarios(usuarios []string) {
	for i, nombre := range usuarios {
		usuario := crearUsuario(nombre)
		servidor.usuarios.Guardar(nombre, usuario)
		servidor.ordenRegistro.Guardar(nombre, i)
	}
}

// obtenerPostPorID devuelve el post que coincide con el ID pasado
func (servidor *AlgoGram) obtenerPostPorID(id int) Post {
	return servidor.posts.Obtener(id)
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
