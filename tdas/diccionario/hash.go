package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

// *** Estructura de hash abierto ***

type celdaHash[K any, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K any, V any] struct {
	tabla     []TDALista.Lista[celdaHash[K, V]]
	cantidad  int
	capacidad int
}

// Funciones auxiliares

// CrearHash devuelve una instancia de hashAbierto
func CrearHash[K any, V any]() Diccionario[K, V] {
	nuevoDiccionario := hashAbierto[K, V]{}
	return &nuevoDiccionario
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

// buscarCelda toma una clave del hash y devuelve la celda correspondiente si existe.
// En caso de no existir devuelve nil
func (hash hashAbierto[K, V]) buscarCelda(clave K) *celdaHash[K, V] {
	// Obtener la lista de celdas
	claveHash := djb2Hash(clave, hash.capacidad)
	listaHash := hash.tabla[claveHash]
	// Crear iterador externo de lista e iterar
	iter := listaHash.Iterador()
	for iter.HaySiguiente() {
		celdaActual := iter.VerActual()
		if celdaActual.clave == clave {
			return &celdaActual
		}
		iter.Siguiente()
	}
	return nil
}

// crearCelda devuelve puntero a una celda creada
func crearCelda[K any, V any](clave K, dato V) *celdaHash[K, V] {
	nuevaCelda := celdaHash[K, V]{clave, dato}
	return &nuevaCelda
}

// Primitivas de hash abierto

func (hash hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash hashAbierto[K, V]) Pertenece(clave K) bool {
	if hash.buscarCelda(clave) != nil {
		return true
	}
	return false
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	celda := hash.buscarCelda(clave)
	// Creaer nueva celda si no existe, sino actualizar
	if celda == nil {
		celda = crearCelda(clave, dato)
		claveHash := djb2Hash(clave, hash.capacidad)
		hash.tabla[claveHash].InsertarUltimo(*celda)
	} else {
		celda.dato = dato
	}
	hash.cantidad++
}
