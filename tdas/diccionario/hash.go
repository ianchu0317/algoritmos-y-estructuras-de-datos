package diccionario

import "fmt"

// *** Estructura hash cerrado ***

const (
	VACIO            = 0
	OCUPADO          = 1
	BORRADO          = -1
	FACT_REDIMENSION = 0.7
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
	// comparar con factor redimension
	// crear nueva tabla
	// pasar claves actuales a claves nuevas (copiar)
}

// Primitivas hash cerrada

func (hash hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

}

func (hashCerrado[K, V]) Pertenece(clave K) bool {

}

func (hashCerrado[K, V]) Obtener(clave K) V {

}

func (hash *hashCerrado[K, V]) Borrar(clave K) {

}

func (hash hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {

}

func (hash hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

}

// *** Estructura iterador externo ***
type iteradorDiccionario[K any, V any] struct {
	tabla       []celdaHash[K, V]
	largoTabla  int
	celdaActual int
}

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {

}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {

}

func (iter iteradorDiccionario[K, V]) Siguiente() {

}
