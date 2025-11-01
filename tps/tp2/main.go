package main

import (
	"algogram/servidor"
)

func main() {
	usuarios := []string{"papa", "manzana", "banana"}
	servidor := servidor.CrearServidor(usuarios)

	servidor.Logout()
	servidor.Login("manzana")
	servidor.Logout()
	servidor.Login("manzana")
	servidor.Logout()
	servidor.Logout()
	servidor.Publicar("hola hwyy")
	servidor.Login("manzana")
	servidor.Publicar("primera publicacion")
	servidor.Publicar("Segundo publicacion")
	servidor.VerProxFeed()
	servidor.Logout()
	servidor.Login("papa")
	servidor.Publicar("tercera publicacion")
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.Logout()
	servidor.Login("manzana")
	servidor.Publicar("Cuarta publicacion")
	servidor.Logout()
	servidor.Login("banana")
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.VerProxFeed()
	servidor.VerProxFeed()
}
