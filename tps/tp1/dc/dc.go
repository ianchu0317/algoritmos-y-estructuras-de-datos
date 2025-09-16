package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// leer args
	args := os.Args
	fmt.Println(args)

	// crear un scanner y leer stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Linea es:", scanner.Text())
	}
}
