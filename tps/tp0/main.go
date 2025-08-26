package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Cargar vector desde archivo.in que contiene enteros por línea
func CargarVectorDeArchivo(archivo *os.File) []int {
	vector := []int{}
	contenido := bufio.NewScanner(archivo)
	var numero int
	for contenido.Scan() {
		numero, _ = strconv.Atoi(contenido.Text())
		vector = append(vector, numero)
	}
	return vector
}

func main() {
	// Leer archivos
	const ruta1 = "archivo1.in"
	const ruta2 = "archivo2.in"

	archivo1, err1 := os.Open(ruta1)
	if err1 != nil {
		fmt.Printf("Error %s abriendo archivo '%s'", err1, ruta1)
		return
	}
	defer archivo1.Close()

	archivo2, err2 := os.Open(ruta2)
	if err2 != nil {
		fmt.Printf("Error '%s' abriendo archivo '%s'", err2, ruta2)
		return
	}
	defer archivo2.Close()

	// Leer archivos
	//vector1 := []int{}
	//vector2 := []int{}

	vector1 := CargarVectorDeArchivo(archivo1)
	vector2 := CargarVectorDeArchivo(archivo2)

	fmt.Println(vector1)
	fmt.Println(vector2)
}
