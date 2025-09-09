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

func TestEncolar(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaFlotantes := TDACola.CrearColaEnlazada[float64]()
	colaCadenas := TDACola.CrearColaEnlazada[string]()

	colaEnteros.Encolar(1)
	colaFlotantes.Encolar(1.1)
	colaCadenas.Encolar("palabra1")

	require.False(t, colaEnteros.EstaVacia(), "Cola de enteros encolados no puede estar vacia")
	require.False(t, colaFlotantes.EstaVacia(), "Cola de flotantes encolados no puede estar vacia")
	require.False(t, colaCadenas.EstaVacia(), "Cola de cadenas encolados no puede que estar vacia")

	// Añadir otros elementos
	colaEnteros.Encolar(2)
	colaFlotantes.Encolar(2.2)
	colaCadenas.Encolar("palabra2")

	require.Equal(t, 1, colaEnteros.VerPrimero(), "Primero de cola de enteros tiene que coincidir con primero ingresado")
	require.Equal(t, 1.1, colaFlotantes.VerPrimero(), "Primero de cola de flotantes tiene que coincidir con primero ingresado")
	require.Equal(t, "palabra1", colaCadenas.VerPrimero(), "Primero de cola de cadenas tiene que coincidir con primero ingresado")
}

func TestDesencolar(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaFlotantes := TDACola.CrearColaEnlazada[float64]()
	colaCadenas := TDACola.CrearColaEnlazada[string]()

	// Encolar un elemento
	colaEnteros.Encolar(1)
	colaFlotantes.Encolar(1.1)
	colaCadenas.Encolar("palabra1")

	require.Equal(t, 1, colaEnteros.Desencolar(), "Desencolar cola de enteros tiene que coincidir con primero ingresado")
	require.Equal(t, 1.1, colaFlotantes.Desencolar(), "Desencolar cola de flotantes tiene que coincidir con primero ingresado")
	require.Equal(t, "palabra1", colaCadenas.Desencolar(), "Desencolar cola de cadenas tiene que coincidir con primero ingresado")

	require.True(t, colaEnteros.EstaVacia(), "Cola de enteros encolados tiene que estar vacia luego de encolar y desencolar un elemento")
	require.True(t, colaFlotantes.EstaVacia(), "Cola de flotantes encolados tiene que estar vacia luego de encolar y desencolar un elemento")
	require.True(t, colaCadenas.EstaVacia(), "Cola de cadenas encolados tiene que estar vacia luego de encolar y desencolar un elemento")

}
