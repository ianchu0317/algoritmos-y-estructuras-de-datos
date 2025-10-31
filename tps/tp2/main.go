package main

import (
	"algogram/servidor"
)

func main() {
	usuarios := []string{"papa", "manzana", "banana"}
	servidor := servidor.CrearServidor(usuarios)

	servidor.Login("manzana")
	servidor.Login("manzana")
}
