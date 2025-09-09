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

}
