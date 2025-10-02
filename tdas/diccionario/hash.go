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
	tabla     []celdaHash[K, V]
	capacidad int
	cantidad  int
	comparar  func(K, K) bool
}

// Funciones auxiliares

func CrearHash[K any, V any]() Diccionario[K, V] {
	nuevoDic := hashCerrado[K, V]{}
	return &nuevoDic
}

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

// buscarCelda toma una clave y devuelve la posicion de celda correspondiente si existe.
// En caso de no existir devuelve la posicion donde deberia existir (la proxima posicion VACIA)
func (hash hashCerrado[K, V]) buscarCelda(clave K) int {
	// Buscar en celda actual
	claveHash := djb2Hash(clave, hash.capacidad)
	celda := hash.tabla[claveHash]
	// Buscar en celdas siguientes si no está en actual
	for celda.estado != VACIO && !hash.comparar(clave, celda.clave) {
		claveHash++
		// Si me pase indice, iniciar de nuevo en principio de tabla
		if claveHash == hash.capacidad {
			claveHash = 0
		}
		celda = hash.tabla[claveHash]
	}
	// devolver celda encontrada
	return claveHash
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

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	// Hallar celda donde deberia estar la clave
	posCelda := hash.buscarCelda(clave)
	celda := hash.tabla[posCelda]
	// Si celda esta ocupada y las claves coinciden entonces pertenece a diccionario
	return celda.estado == OCUPADO && hash.comparar(clave, celda.clave)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	posCelda := hash.buscarCelda(clave)
	celda := hash.tabla[posCelda]
	// Guardar informacion a la celda
	celda.clave = clave
	celda.dato = dato
	celda.estado = OCUPADO
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	posCelda := hash.buscarCelda(clave)
	celda := hash.tabla[posCelda]
	return celda.dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	posCelda := hash.buscarCelda(clave)
	celda := hash.tabla[posCelda]
	celda.estado = BORRADO
	return celda.dato
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
