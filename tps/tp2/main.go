package main

import (
	"algogram/servidor"
	"bufio"
	"os"
	"strconv"
	"strings"
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
	textoSlice := strings.Split(texto, " ")
	return textoSlice[0], strings.Join(textoSlice[1:], " ")
}

// toma un parametro y devuelve en entero
func obtenerID(parametro string) int {
	num, _ := strconv.Atoi(parametro) // asumir que no hay error en conversion
	return num
}

// ejecutarServidor toma un servidor, un comando y los paramtros correspondientes al comando y lo ejecuta en el servidor.
func ejecutarComando(servidor servidor.Servidor, lineaActual string) {
	comando, parametro := procesarLinea(lineaActual)

	switch comando {
	case "login":
		servidor.Login(parametro)
	case "logout":
		servidor.Logout()
	case "publicar":
		servidor.Publicar(parametro)
	case "ver_siguiente_feed":
		servidor.VerProxFeed()
	case "likear_post":
		servidor.Likear(obtenerID(parametro))
	case "mostrar_likes":
		servidor.MostrarLikes(obtenerID(parametro))
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
		ejecutarComando(servidor, lineaActual)
	}
}
