package diccionario

import "fmt"

// *** Estructura de ABB ***

type nodo[K any, V any] struct {
	clave K
	dato  V
	izq   *nodo[K, V]
	der   *nodo[K, V]
}

type arbolBinario[K any, V any] struct {
	raiz     *nodo[K, V]
	comparar func(K, K) int
	cantidad int
}

// crearNodo crea un nodo para el abb
func crearNodo[K any, V any](clave K, dato V) *nodo[K, V] {
	return &nodo[K, V]{clave: clave, dato: dato, izq: nil, der: nil}
}

// CrearABB crea una instancia de diccionario ordenado
func CrearABB[K any, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &arbolBinario[K, V]{nil, funcionCmp, 0}
}

// Funciones auxiliares

// convertirBytes toma la clave y la convierte a bytes
func convertirBytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2HashInt toma una clave y devuelve un entero unico con respecto a la clave.
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func djb2HashInt[K any](clave K) int {
	bytes := convertirBytes(clave)
	hash := uint32(5381)
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + uint32(b) // hash*33 + b
	}
	return int(hash)
}

// Primitivas de ABB
func (abb arbolBinario[K, V]) Guardar(clave K, dato V) {

}

func (abb arbolBinario[K, V]) Pertenece(clave K) bool {
	return false
}

func (abb arbolBinario[K, V]) Obtener(clave K) V {
	var data V
	return data
}

func (abb arbolBinario[K, V]) Borrar(clave K) V {
	var data V
	return data
}

func (abb arbolBinario[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb arbolBinario[K, V]) Iterar(visitar func(clave K, dato V) bool) {
}

func (abb arbolBinario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

}

func (abb arbolBinario[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorABB[K, V]{nil, nil}
}

func (abb arbolBinario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return &iteradorABB[K, V]{desde, hasta}
}

// *** Estructura Iterador Externo ABB ***
type iteradorABB[K any, V any] struct {
	desde *K
	hasta *K
}

func (iter iteradorABB[K, V]) HaySiguiente() bool {
	return false
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {

}

func (iter *iteradorABB[K, V]) Siguiente() {

}
