package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaFlotantes := TDAPila.CrearPilaDinamica[float32]()
	pilaCadenas := TDAPila.CrearPilaDinamica[string]()

	require.True(t, pilaEnteros.EstaVacia(), "Pila de numeros enteros recien creada tiene que estar vacia")
	require.True(t, pilaFlotantes.EstaVacia(), "Pila de numeros flotantes recien creada tiene que estar vacia")
	require.True(t, pilaCadenas.EstaVacia(), "Pila de cadenas recien creada tiene que estar vacia")
	require.Panics(t, func() { pilaEnteros.VerTope() }, "Pila de numeros enteros recien creada no puede ver tope")
	require.Panics(t, func() { pilaFlotantes.VerTope() }, "Pila de numeros flotantes recien creada no puede ver tope")
	require.Panics(t, func() { pilaCadenas.VerTope() }, "Pila de cadenas recien creada no puede ver tope")
	require.Panics(t, func() { pilaEnteros.Desapilar() }, "Pila de numeros enteros recien creada no puede desapilar")
	require.Panics(t, func() { pilaFlotantes.Desapilar() }, "Pila de numeros flotantes recien creada no puede desapilar")
	require.Panics(t, func() { pilaCadenas.Desapilar() }, "Pila de cadenas recien creada no puede desapilar")
}
