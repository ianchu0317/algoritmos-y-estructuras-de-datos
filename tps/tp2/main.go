package main

import (
	"algogram/servidor"
	"bufio"
	"os"

	"golang.org/x/text/cases"
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

// procesarLinea toma una linea de string y devuelve el comando y el parametro de la linea
func procesarLinea(texto string) (string, string) {
	comando := texto
	parametro := texto
	return comando, parametro
}

// ejecutarServidor toma un servidor, un comando y los paramtros correspondientes al comando y lo ejecuta en el servidor.
func ejecutarComando(servidor servidor.Servidor, comando, parametro string) {
	switch comando 
	}
}

func main() {
	archivo := os.Args[1]
	usuarios := cargarUsuarios(archivo)
	servidor := servidor.CrearServidor(usuarios)
	// Leer STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineaActual := scanner.Text()
		cmd, parametro := procesarLinea(lineaActual)
		ejecutarComando(servidor, cmd, parametro)
	}
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
