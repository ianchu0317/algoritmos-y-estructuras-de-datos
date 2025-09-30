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
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	listaCadenas := TDALista.CrearListaEnlazada[string]()

	// Insertar una lista en primero y otra en ultimo,
	// como tienen un elemento deberia ser lo mismo
	listaEnteros.InsertarPrimero(1)
	listaCadenas.InsertarUltimo("a")

	// lista.EstaVacia() debe ser false
	require.Equal(t, false, listaEnteros.EstaVacia(), "Lista de enteros con un elemento no debe estar vacia (EstaVacia()-> true)")
	require.Equal(t, false, listaCadenas.EstaVacia(), "Lista de cadenas con un elemento no debe estar vacia (EstaVacia()-> true)")

	// lista.Largo debe ser 1
	require.Equal(t, 1, listaEnteros.Largo(), "Lista de enteros con un elemento debe tener largo 1")
	require.Equal(t, 1, listaCadenas.Largo(), "Lista de cadenas con un elemento debe tener largo 1")

	// lista.VerPrimero == listaVerUltimo
	require.Equal(t, listaEnteros.VerPrimero(), listaEnteros.VerUltimo(), "Lista de enteros de un elemento coincide primer y ultimo elemento")
	require.Equal(t, listaCadenas.VerPrimero(), listaCadenas.VerUltimo(), "Lista de cadenas de un elemento coincide primer y ultimo elemento")

	// lista.BorrarPrimero() debe coincidir con elemento ingresado
	require.Equal(t, 1, listaEnteros.BorrarPrimero(), "Lista de enteros de un elemento debe devolver unico elemento al borrar primero")
	require.Equal(t, "a", listaCadenas.BorrarPrimero(), "Lista de cadenas de un elemento debe devolver unico elemento al borrar primero")

}

// Test de caso especifico
func TestOrdenLista(t *testing.T) {
	ordenEnteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	listaEnteros := TDALista.CrearListaEnlazada[int]()

	// ORDEN INSERTAR PRIMERO
	// Llenar y vaciar la lista en orden
	for _, num := range ordenEnteros {
		listaEnteros.InsertarPrimero(num)
		require.Equal(t, num, listaEnteros.VerPrimero(), "El valor de ingresado tiene que coincidir con primer elemento")
	}
	for i := len(ordenEnteros) - 1; i >= 0; i-- {
		require.Equal(t, ordenEnteros[i], listaEnteros.BorrarPrimero(), "El valor borrado tiene que salir en orden como la inversa de arreglo")
	}
	// Lista luego de llenar y vaciar se comporta como lista vacia
	require.Equal(t, true, listaEnteros.EstaVacia(), "Lista de enteros recien creada debe estar vacia (EstaVacia()-> true)")

	// ORDEN INSERTAR ULTIMO
	for _, num := range ordenEnteros {
		listaEnteros.InsertarUltimo(num)
		require.Equal(t, num, listaEnteros.VerUltimo(), "El valor de ingresado tiene que coincidir con ultimo elemento")
	}
	for _, num := range ordenEnteros {
		require.Equal(t, num, listaEnteros.BorrarPrimero(), "El valor borrado tiene que salir en orden del arreglo")
	}
	// Lista luego de llenar y vaciar se comporta como lista vacia
	require.Equal(t, true, listaEnteros.EstaVacia(), "Lista de enteros recien vaciada debe estar vacia (EstaVacia()-> true)")
}

func TestVolumen(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	volumen := 10000

	// Llenar con 10000 elementos y vaciar 10000 elementos
	for i := 0; i <= volumen; i++ {
		listaEnteros.InsertarPrimero(i)
		require.Equal(t, i, listaEnteros.VerPrimero(), "El valor de ingresado tiene que coincidir con ultimo ingresado test volumen")
	}
	for i := volumen; i >= 0; i-- {
		require.Equal(t, i, listaEnteros.BorrarPrimero(), "El valor borrado tiene que salir en orden del ingresado volumen")
	}

	require.Equal(t, true, listaEnteros.EstaVacia(), "Lista de enteros vaciada debe estar vacia (EstaVacia()-> true)")
}

// *** Test de iteradores ***
// Test iterador interno
func TestIteradorInterno(t *testing.T) {
	listaEnteros := TDALista.CrearListaEnlazada[int]()

	// Usar iterador interno para sumar todos los elementos de la lista
	arregloEnteros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // Expectativa 550
	sumatoriaIterar := 0

	for _, num := range arregloEnteros {
		listaEnteros.InsertarPrimero(num)
	}
	listaEnteros.Iterar(func(elemento int) bool {
		sumatoriaIterar += elemento
		return true
	})
	require.Equal(t, 550, sumatoriaIterar, "La sumatoria de los elementos usando iterador interno no coincide con expectativa")

	// Test con false y corte en el medio al encontrar numero par
	nuevaListaEnteros := TDALista.CrearListaEnlazada[int]()
	arregloUnSoloPar := []int{1, 3, 5, 7, 9, 2, 13, 11}
	primerElementoPar := 2
	var primerParEncontrado int

	for _, num := range arregloUnSoloPar {
		nuevaListaEnteros.InsertarPrimero(num)
	}

	nuevaListaEnteros.Iterar(func(numero int) bool {
		if numero%2 == 0 {
			primerParEncontrado = 2
			return false
		}
		return true
	})
	require.Equal(t, primerElementoPar, primerParEncontrado, "El primer par encontrado debe ser numero 2")
}

// Tests iterador externo
func TestIteradorExterno(t *testing.T) {
	// Test insertar con iterador recien creado
	// Debe comportarse como insertar un elemento, reutilizar tests anteriores
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	iter := listaEnteros.Iterador()
	iter.Insertar(10)

	require.Equal(t, 10, listaEnteros.VerPrimero(), "Insertar con iterador recién creado debe ser primer elemento")
	require.Equal(t, 10, listaEnteros.VerUltimo(), "En lista de un elemento, primero y ultimo deben coincidir")
	require.Equal(t, 1, listaEnteros.Largo(), "Lista debe tener largo 1 después de insertar")

	// Test insertar al final
	arregloEnteros := []int{20, 30, 40}
	for _, num := range arregloEnteros {
		listaEnteros.InsertarUltimo(num)
	}

	iterFinal := listaEnteros.Iterador()
	for iterFinal.HaySiguiente() {
		iterFinal.Siguiente()
	}
	iterFinal.Insertar(50)
	require.Equal(t, 50, listaEnteros.VerUltimo(), "Insertar al final con iterador debe cambiar ultimo elemento")

	// Test insertar en el medio
	nuevoArregloEnteros := []int{10, 30, 40}
	nuevaLista := TDALista.CrearListaEnlazada[int]()
	for _, num := range nuevoArregloEnteros {
		nuevaLista.InsertarUltimo(num)
	}

	iterMedio := nuevaLista.Iterador()
	iterMedio.Siguiente()
	iterMedio.Insertar(20) // resultado deberia ser...10, 20, 30, 40
	require.Equal(t, 10, nuevaLista.BorrarPrimero(), "Primer elemento debe ser 10")
	require.Equal(t, 20, nuevaLista.BorrarPrimero(), "Segundo elemento debe ser 20 (insertado)")

	// Test borrar primer elemento
	nuevoArregloEnteros2 := []int{100, 200, 300}
	listaBorrar := TDALista.CrearListaEnlazada[int]()
	for _, num := range nuevoArregloEnteros2 {
		listaBorrar.InsertarUltimo(num)
	}

	iterBorrar := listaBorrar.Iterador()
	elementoBorrado := iterBorrar.Borrar()
	require.Equal(t, 100, elementoBorrado, "Elemento borrado debe ser 100")
	require.Equal(t, 200, listaBorrar.VerPrimero(), "Nuevo primer elemento debe ser 200")

	// Test borrar último elemento
	iterBorrar.Siguiente() // ir a 300
	elementoUltimoBorrado := iterBorrar.Borrar()
	require.Equal(t, 300, elementoUltimoBorrado, "Elemento borrado debe ser 300")
	require.Equal(t, 200, listaBorrar.VerUltimo(), "Nuevo último elemento debe ser 200")
}
