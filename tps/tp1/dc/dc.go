// Algoritmo en https://es.wikipedia.org/wiki/Notaci%C3%B3n_polaca_inversa

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// crear un scanner y leer stdin
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lineaActual := s.Text()
		fmt.Println("Linea es:", lineaActual)
	}
}
