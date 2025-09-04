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

func TestApilar(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaFlotantes := TDAPila.CrearPilaDinamica[float32]()
	pilaCadenas := TDAPila.CrearPilaDinamica[string]()

	pilaEnteros.Apilar(1)
	pilaFlotantes.Apilar(1.1)
	pilaCadenas.Apilar("palabra")

	require.False(t, pilaEnteros.EstaVacia(), "Pila de enteros apilados no puede estar vacia")
	require.False(t, pilaFlotantes.EstaVacia(), "Pila de flotantes apilados no puede estar vacia")
	require.False(t, pilaCadenas.EstaVacia(), "Pila de cadenas apilaods no puede que estar vacia")
	require.Equal(t, 1, pilaEnteros.VerTope(), "Tope de pila de enteros tiene que coincidir con ultimo ingresado")
	require.Equal(t, 1.1, pilaFlotantes.VerTope(), "Tope de pila de flotantes tiene que coincidir con ultimo ingresado")
	require.Equal(t, "palabra", pilaCadenas.VerTope(), "Tope de pila de cadenas tiene que coincidir con ultimo ingresado")
}

func TestDesapilar(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaFlotantes := TDAPila.CrearPilaDinamica[float32]()
	pilaCadenas := TDAPila.CrearPilaDinamica[string]()

	pilaEnteros.Apilar(1)
	pilaFlotantes.Apilar(1.1)
	pilaCadenas.Apilar("palabra")

	require.Equal(t, 1, pilaEnteros.Desapilar(), "Desapilar pila de enteros tiene que devolver elemento desapilado")
	require.Equal(t, 1.1, pilaFlotantes.Desapilar(), "Desapilar pila de flotantes tiene que devolver elemento desapilado")
	require.Equal(t, "palabra", pilaCadenas.Desapilar(), "Desapilar pila de cadenas tiene que devolver elemento desapilado")
	require.True(t, pilaEnteros.EstaVacia(), "Pila de numeros enteros apilados y desapilados tiene que estar vacia")
	require.True(t, pilaFlotantes.EstaVacia(), "Pila de numeros flotantes apilados y desapilados tiene que estar vacia")
	require.True(t, pilaCadenas.EstaVacia(), "Pila de cadenas apilados y desapilados tiene que estar vacia")
}
