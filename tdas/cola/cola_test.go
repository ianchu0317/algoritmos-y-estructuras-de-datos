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

// TEST CON VECTORES
// Funciones auxiliares para testear encolar y desencolar desde vector

func encolarDesdeVector[T any](cola TDACola.Cola[T], vec []T) {
	for _, elemento := range vec {
		cola.Encolar(elemento)
	}
}

func desencolarComparandoVector[T any](t *testing.T, cola TDACola.Cola[T], vec []T) {
	for _, elemento := range vec {
		require.Equal(t, elemento, cola.Desencolar(), "El orden de desencolado tiene que coincidir con orden de vector")
	}
}

// TestColaConVectores chequea orden de encolamiento y desencolamiento sea correcto
func TestColaConVectores(t *testing.T) {
	vec1 := []int{5, -6, 7, 2, 0}
	vec2 := []string{"perro", "gato", "pollito", "elefante", "jirafa"}
	vec3 := []bool{true, false, false, false, true, true}
	vec4 := []float64{1.1, 2.2, 3.3, 4.4, -1.1, -3.4}

	cola1 := TDACola.CrearColaEnlazada[int]()
	cola2 := TDACola.CrearColaEnlazada[string]()
	cola3 := TDACola.CrearColaEnlazada[bool]()
	cola4 := TDACola.CrearColaEnlazada[float64]()

	// encolar valores de vector
	encolarDesdeVector(cola1, vec1)
	encolarDesdeVector(cola2, vec2)
	encolarDesdeVector(cola3, vec3)
	encolarDesdeVector(cola4, vec4)

	// desencolar valores y comparar con orden de vector
	desencolarComparandoVector(t, cola1, vec1)
	desencolarComparandoVector(t, cola2, vec2)
	desencolarComparandoVector(t, cola3, vec3)
	desencolarComparandoVector(t, cola4, vec4)
}

// TEST DE VOLUMEN
// Testear que se pueda encolar y desencolar gran cantidad de elemento respetando orden

func testearConRango(t *testing.T, cola TDACola.Cola[int], volumen int) {
	for i := 1; i <= volumen; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= volumen; i++ {
		require.Equal(t, i, cola.Desencolar(), "El valor desencolado tiene que coincidir con orden ingresado")
	}
	require.True(t, cola.EstaVacia(), "Luego de encolar y desencolar misma cantidad la cola tiene que estar vacia")
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	volumen1 := 1000
	volumen2 := 10000

	testearConRango(t, cola, volumen1)
	testearConRango(t, cola, volumen2)
}
