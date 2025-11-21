package grafo

import Diccionario "tdas/diccionario"

type grafoListaAdyacencia[T any, K Numero] struct {
	vertices   Diccionario.Diccionario[T, K]
	esDirigido bool
	esPesado   bool
}

// CrearGrafo(esDirigido, esPesado) devuelve una instancia de Grafo

// Primitivas de grafo
