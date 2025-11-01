package main

import (
	"bufio"
	"fmt"
	"os"
)

// cargarUsuarios toma un path de archivo y devuelve los usuarios contenidos en el mismo
func cargarUsuarios(pathArchivo string) []string {
	archivo, _ := os.Open(pathArchivo) // asumiendo que no hay error en abrir archivo
	defer archivo.Close()

	usuarios := []string{}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		usuarios = append(usuarios, scanner.Text())
	}
	return usuarios
}

func main() {
	archivo := os.Args[1]
	usuarios1 := cargarUsuarios(archivo)
	fmt.Println(usuarios1)
	// Debug
	/*
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
	*/
}
