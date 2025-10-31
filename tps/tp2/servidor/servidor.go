package servidor

// Interfaz de servidor

type Servidor interface {
	// Login(nombre) toma un nombre de usuario (cadena) e inicia sesión en servidor.
	// Si el usuario está en el servidor, se imprime 'Hola <nombre>'.
	// Si el usuario no existe en servidor, se imprime 'Error: usuario no existente'.
	// Si ya hay usuario en la sesión actual, se imrpime 'Error: Ya haboa un usuario loggeado'.
	Login(nombre string)

	// Logout() cierra sesión del usuario actual imprimiendo 'Adios <usuario>'.
	// Si no hay usuario en la sesión actual imprime 'Error: no habia usuario loggeado'.
	Logout()

	// Publicar(contenido) toma un contenido en formato cadena y crea un contenido en el servidor.
	// DEBE Haber usuario en la sesión actual para publicar.
	// Imprime 'Post publicado' si publica exitosamente.
	// En caso que no hubiera usuario loggeado, se imprime 'Error: no habia usuario logeado'.
	Publicar(contenido string)

	// VerProxFeed() devuelve la proxima publicacion del feed del usuario actual.
	// Si usuario no tiene mas publicaciones o no esta loggeado se imprime 'Usuario no loggeado o no hay mas posts para ver'.
	VerProxFeed()

	// Likear(id) toma el id de una publicación y le da like.
	// Si hay usuario loggeado y post existe entonces imprime 'Post likeado'.
	// De lo contrario 'Error: Usuario no loggeado o Post inexistente'.
	Likear(id int)

	// MostrarLikes(id) toma el id de una publicación existente y devuelve la cantidad de likes con los usuarios que likearon.
	// En caso de no existir la publicacion o no tiene likes imprime 'Error: Post inexistente o sin likes'.
	MostrarLikes(id int)
}
