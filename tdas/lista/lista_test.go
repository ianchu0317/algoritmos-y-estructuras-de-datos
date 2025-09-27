package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// *** Test de TDA Lista sin iterador ***

// Test de la lista vacia
func TestListaVacia(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	listaCadenas := TDALista.CrearListaEnlazada[string]()

	// lista.EstaVacia debe devolver 'true'
	require.Equal(t, true, listaEnteros.EstaVacia(), "Lista de enteros recien creada debe estar vacia (EstaVacia()-> true)")
	require.Equal(t, true, listaCadenas.EstaVacia(), "Lista de cadenas recien creada debe estar vacia (EstaVacia()-> true)")
	// lista.Largo debe ser 0
	require.Equal(t, 0, listaEnteros.Largo(), "Lista de enteros recien creada debe ser largo 0")
	require.Equal(t, 0, listaCadenas.Largo(), "Lista de cadenas recien creada debe ser largo 0")
	// lista.VerPrimero() y listaVerUltimo() debe llamar panic
	require.Panics(t, func() { listaEnteros.VerPrimero() }, "Lista de enteros recien creada no puede ver primero")
	require.Panics(t, func() { listaCadenas.VerPrimero() }, "Lista de cadenas recien creada no puede ver primero")
	require.Panics(t, func() { listaEnteros.VerUltimo() }, "Lista de enteros recien creada no puede ver ultimo")
	require.Panics(t, func() { listaCadenas.VerUltimo() }, "Lista de cadenas recien creada no puede ver ultimo")
	// lista.BorrarPrimero() debe ser panic
	require.Panics(t, func() { listaEnteros.BorrarPrimero() }, "Lista de enteros recien creada no puede Borrar primero")
	require.Panics(t, func() { listaCadenas.BorrarPrimero() }, "Lista de cadenas recien creada no puede Borrar primero")
}

// Test caso borde de un solo elemento
func TestUnSoloElemento(t *testing.T) {
	// lista.EstaVacia() debe ser false
	// lista.Largo debe ser 1
	// lista.VerPrimero == listaVerUltimo
	// lista.BorrarPrimero() debe coincidir con elemento ingresado
}

// Test de caso especifico
func TestEspecificos(t *testing.T) {
	// Llenar y vaciar la lista en orden

	// Lista luego de llenar y vaciar se comporta como lista vacia
}

func TestVolumen(t *testing.T) {
	// Llenar con 10000 elementos y vaciar 10000 elementos
}

// *** Test de iteradores ***
// Test iterador interno
func TestIteradorInterno(t *testing.T) {
	// Test con true en todos los casos (sumatoria en una lista de enteros)
	// Test con false y corte en el medio al encontrar numero par
	// Caso borde solo un elemento en lista
}

// Test iterador externo
func TestIteradorExterno(t *testing.T) {
	// Insertar un elemento en lista con iterador recien creado
	// debe comportarse a lista primer elemento
}

// - Insertar primer elemento de iterador
// - Insertar elemento e interador al final
// - Insertar elemento en medio de iterador, en su posicion correcta
// - Remover primer elemento, cambia lista original
// - Verificar que remover elemento del medio no esta
// - Casos borde espeficiso
