package auxiliares

import (
	"math/rand"
)

// Crear un array de numeros random entre 0 y 100 dado cantidad
func CrearArrayRandom(cantidad int) []int {
	nuevoArr := make([]int, 0)
	for range cantidad {
		nuevoArr = append(nuevoArr, rand.Intn(100))
	}
	return nuevoArr
}
