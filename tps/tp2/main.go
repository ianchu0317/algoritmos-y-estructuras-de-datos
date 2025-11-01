package main

import (
	"algogram/servidor"
)

// cargarUsuarios toma un path de archivo y devuelve los usuarios contenidos en el mismo
func cargarUsuarios(archivo string) []string {
	usuarios := []string{}
	return usuarios
}

func main() {
	usuarios := []string{"chorch", "cacatua2030", "mondi", "chicho1994", "eldiego"}
	servidor := servidor.CrearServidor(usuarios)

	servidor.Login("chicho1994")
	servidor.Publicar("TIene todo el dinero del mundo, pero hay algo que no puede comprar... un dinosaurio")
	servidor.Logout()
	servidor.Login("chorch")
	servidor.Publicar("te corto internert")
	servidor.Publicar("es por el teorema de chuck norris")
	servidor.Logout()
	servidor.Login("cacatua2030")
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.Logout()
	servidor.Login("mondi")
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.Logout()
	servidor.Login("chicho1994")
	servidor.VerProxFeed()
}
