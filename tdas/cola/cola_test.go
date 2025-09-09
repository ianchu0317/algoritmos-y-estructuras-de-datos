package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaFlotantes := TDACola.CrearColaEnlazada[float64]()
	colaCadenas := TDACola.CrearColaEnlazada[string]()

	require.True(t, colaEnteros.EstaVacia(), "Cola de enteros recien creada deberia estar vacia")
	require.True(t, colaFlotantes.EstaVacia(), "Cola de floantes recien creada deberia estar vacia")
	require.True(t, colaCadenas.EstaVacia(), "Cola de cadenas recien creada deberia estar vacia")
	require.Panics(t, func() { colaEnteros.VerPrimero() }, "Cola de numeros enteros recien creada no puede ver primero")
	require.Panics(t, func() { colaFlotantes.VerPrimero() }, "Cola de numeros flotantes recien creada no puede ver primero")
	require.Panics(t, func() { colaCadenas.VerPrimero() }, "Cola de cadenas recien creada no puede ver primero")
	require.Panics(t, func() { colaEnteros.Desencolar() }, "Cola de numeros enteros recien creada no puede desencolar")
	require.Panics(t, func() { colaFlotantes.Desencolar() }, "Cola de numeros flotantes recien creada no puede desencolar")
	require.Panics(t, func() { colaCadenas.Desencolar() }, "Cola de cadenas recien creada no puede desencolar")
}
