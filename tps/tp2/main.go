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
}
