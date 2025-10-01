package diccionario

import TDALista "tdas/lista"

// *** Estructura de hash abierto ***

type celda[K any, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K any, V any] struct {
	arreglo  []TDALista.Lista[celda[K, V]]
	cantidad int
}
