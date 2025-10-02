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
	comparar  func(K, K) bool
}

// Funciones auxiliares

// CrearHash devuelve una instancia de hashAbierto
func CrearHash[K any, V any](comparar func(K, K) bool) Diccionario[K, V] {
	nuevoDiccionario := hashAbierto[K, V]{
		make([]TDALista.Lista[celdaHash[K, V]], 100),
		0,
		100,
		comparar}
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
		if hash.comparar(celdaActual.clave, clave) {
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
	return hash.buscarCelda(clave) != nil
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

func (hash hashAbierto[K, V]) Obtener(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	celda := hash.buscarCelda(clave)
	return celda.dato
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	claveHash := djb2Hash(clave, hash.capacidad)
	iter := hash.tabla[claveHash].Iterador()
	var dato V
	for iter.HaySiguiente() {
		celdaActual := iter.VerActual()
		if hash.comparar(celdaActual.clave, clave) {
			dato = iter.Borrar().dato
		}
		iter.Siguiente()
	}
	return dato
}

func (hash hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	var celdaActual celdaHash[K, V]
	for _, lista := range hash.tabla {
		for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			celdaActual = iter.VerActual()
			// Si es false salir de iteracion
			if !visitar(celdaActual.clave, celdaActual.dato) {
				return
			}
		}
	}
}

func (hash hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIterador := iteradorDiccionario[K, V]{
		hash.tabla,
		hash.capacidad,
		hash.tabla[0].Iterador(),
		0}
	return &nuevoIterador
}

// *** Estructura Iterador Externo ***

type iteradorDiccionario[K any, V any] struct {
	tabla           []TDALista.Lista[celdaHash[K, V]]
	largoTabla      int
	iterListaActual TDALista.IteradorLista[celdaHash[K, V]]
	numeroLista     int
}

// Primitivas iterador externo de diccionario

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.numeroLista < iter.largoTabla && iter.iterListaActual.HaySiguiente()
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	celdaActual := iter.iterListaActual.VerActual()
	return celdaActual.clave, celdaActual.dato
}

func (iter iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	// Si hay elementos en lista actual, avanzar elemento,
	// sino avanzar a una lista que haya elemnetos
	if iter.iterListaActual.HaySiguiente() {
		iter.iterListaActual.Siguiente()
	} else {
		// Crear nuevo iterador para siguiente lista en tabla
		for !iter.iterListaActual.HaySiguiente() && iter.numeroLista < iter.largoTabla {
			iter.numeroLista++
			iter.iterListaActual = iter.tabla[iter.numeroLista].Iterador()
		}
	}
}
