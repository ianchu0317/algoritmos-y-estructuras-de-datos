package diccionario

import "fmt"

// *** Estructura hash cerrado ***

const (
	VACIO   = 0
	BORRADO = -1
	OCUPADO = 1
)

type celdaHash[K any, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K any, V any] struct {
	tabla     []hashCerrado[K, V]
	capacidad int
	cantidad  int
}

// Funciones auxiliares

// convertirABytes toma la clave y la convierte a bytes
func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2Hash toma una clave y el numero maximo que puede valer,
// y devuelve la clave en hash dentro del rango pasado.
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func djb2Hash[K any](clave K, largo int) int {
	bytes := convertirABytes(clave)
	hash := 5381
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + int(b) // hash * 33 + b
	}
	return hash % largo
}

// buscarCelda toma una clave del hash y devuelve la celda correspondiente si existe.
// En caso de no existir devuelve nil
func (hash hashCerrado[K, V]) buscarCelda(clave K) {

}

func (hash *hashCerrado[K, V]) redimensionarTabla() {

}

// Primitiva hash cerrad

// *** Estructura iterador externo ***
