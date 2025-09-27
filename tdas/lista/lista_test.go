package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test de TDA Lista sin iterador

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

// - Caso lista vacia
// - Caso lista con un solo elemento
// - Casos especificos, entrada y salida

// Test iterador interno
// - Test por false (idea numeros pares)
// - Test por true

// Test iterador externo
// - Insertar primer elemento de iterador
// - Insertar elemento e interador al final
// - Insertar elemento en medio de iterador, en su posicion correcta
// - Remover primer elemento, cambia lista original
// - Verificar que remover elemento del medio no esta
// - Casos borde espeficiso
