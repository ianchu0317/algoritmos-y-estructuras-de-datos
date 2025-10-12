package diccionario_test

import (
	"fmt"
	"math/rand"
	"slices"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

// Funciones auxiliares de comparacion

func compararInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func compararStr(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

type cmpStruct struct {
	num int
}

func compararStruct(a, b cmpStruct) int {
	return compararInt(a.num, b.num)
}

func TestAbbVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, bool](compararStr)
	require.Equal(t, 0, abb.Cantidad(), "Un ABB recien creado deberia tener cantidad 0")
	require.False(t, abb.Pertenece("a"), "ABB recien creado no deberia tener ninguna clave que pertence")
	require.Panics(t, func() { abb.Obtener("a") }, "ABB recien creado no se puede obtener clave")
	require.Panics(t, func() { abb.Borrar("a") }, "ABB recien creado no se puede borrar clave")
}

func TestUnSoloElemento(t *testing.T) {
	// Probar primitivas de ABB para un solo elemento
	abb := TDADiccionario.CrearABB[string, string](compararStr)
	clave, dato := "a", "palabra"
	abb.Guardar(clave, dato)
	// Test de un elemento guardado
	require.Equal(t, 1, abb.Cantidad(), "Un ABB con un elemento guardado deberia tener cantidad 1")
	require.True(t, abb.Pertenece(clave), "Clave guardada debe pertenecer a ABB")
	require.Equal(t, abb.Obtener(clave), dato, "ABB con clave guardada deberia devolver bien dato")
	// Test de borrado de ese elemento
	datoBorrado := abb.Borrar(clave)
	require.Equal(t, datoBorrado, dato, "ABB recien creado no se puede borrar clave")
	require.Equal(t, 0, abb.Cantidad(), "ABB sin elementos deberia ser cantidad 0")
	require.False(t, abb.Pertenece(clave), "ABB sin elementos no deberia tener ninguna clave que pertence")
	require.Panics(t, func() { abb.Obtener(clave) }, "ABB sin elementos no se puede obtener clave")
	require.Panics(t, func() { abb.Borrar(clave) }, "ABB sin elementos no se puede borrar clave")
}

func probarPrimitivas(t *testing.T, abb TDADiccionario.DiccionarioOrdenado[cmpStruct, int], clave cmpStruct, dato int) {
	require.True(t, abb.Pertenece(clave), "Clave compleja debe pertenecer a diccionario")
	require.Equal(t, abb.Obtener(clave), dato, "Clave compleja debe coincidir con dato ingresado")
	require.Equal(t, abb.Borrar(clave), dato, "Dato borrado debe coincidir con dato guardado")
}

func TestClaveCompleja(t *testing.T) {
	clave1 := cmpStruct{3}
	clave2 := cmpStruct{6}
	clave3 := cmpStruct{9}
	clave4 := cmpStruct{1}
	clave5 := cmpStruct{2}
	clave6 := cmpStruct{0}
	abb := TDADiccionario.CrearABB[cmpStruct, int](compararStruct)
	abb.Guardar(clave1, 1)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)
	abb.Guardar(clave6, 6)
	require.Equal(t, abb.Cantidad(), 6, "Elementos de ABB debe coincidir con la cantidad guardada")
	// Probar que cumplen las primitivas basicas para las 6 claves anteriores, luego borrarlas del ABB
	probarPrimitivas(t, abb, clave1, 1)
	probarPrimitivas(t, abb, clave2, 2)
	probarPrimitivas(t, abb, clave3, 3)
	probarPrimitivas(t, abb, clave4, 4)
	probarPrimitivas(t, abb, clave5, 5)
	probarPrimitivas(t, abb, clave6, 6)
	require.Equal(t, abb.Cantidad(), 0, "Elementos de ABB debe ser 0 luego de borrar todos")
}

func TestReemplazarDato(t *testing.T) {
	// Claves y datos a guardar
	clave1, clave2, clave3 := 75, 339, 23
	dato1, dato2, dato3 := "palabra1", "palabra2", "palabra3"
	nuevoDat1, nuevoDat2, nuevoDat3 := "dato1", "dato2", "dato3"
	claves := []int{clave1, clave2, clave3}
	datos := []string{dato1, dato2, dato3}
	nuevosDat := []string{nuevoDat1, nuevoDat2, nuevoDat3}
	abb := TDADiccionario.CrearABB[int, string](compararInt)
	// Guardar los datos
	for i := range 3 {
		abb.Guardar(claves[i], datos[i])
	}
	// Verificar que pertenecen
	for i := range 3 {
		require.True(t, abb.Pertenece(claves[i]), "Clave guardada tiene que pertenecer a diccionario")
		require.Equal(t, abb.Obtener(claves[i]), datos[i], "Obtneer clave guardada tiene que coincidir con dato guardado")
	}
	// Reemplazar los datos y verificar cambio
	for i := range 3 {
		abb.Guardar(claves[i], nuevosDat[i])
		require.True(t, abb.Pertenece(claves[i]), "Clave guardada tiene que pertenecer a diccionario")
		require.Equal(t, abb.Obtener(claves[i]), nuevosDat[i], "Obtneer clave guardada tiene que coincidir con dato guardado")
	}

}

// countinSort ordena un arreglo de numeros enteros dentro del rango
func countingSort(arr []int, rango int) []int {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[num]++
	}
	inicios := make([]int, rango)
	for i := 1; i < rango; i++ {
		inicios[i] = inicios[i-1] + frecuencias[i-1]
	}
	ordenado := make([]int, len(arr))
	for _, num := range arr {
		indxInicio := inicios[num]
		ordenado[indxInicio] = num
		inicios[num]++
	}
	return ordenado
}

// crearArregloDesordenado toma cantidad y rango en enteros.
// Devuelve un arreglo de largo 'cantidad' desordenado de numeros aleatorios dentro del rango.
func crearArregloDesordenado(cantidad, rango int) []int {
	nuevoArr := make([]int, cantidad)
	for i := range cantidad {
		numRandom := rand.Int() % rango
		// Crear arreglo desordenado de numeros no repetidos
		for slices.Contains(nuevoArr, numRandom) {
			numRandom = rand.Int() % rango
		}
		nuevoArr[i] = numRandom
	}
	return nuevoArr
}

func TestIteradorInterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	rango := 100
	arregloDesordenado := crearArregloDesordenado(10, rango)
	arregloOrdenado := countingSort(arregloDesordenado, rango)
	fmt.Println(arregloDesordenado)
	fmt.Println(arregloOrdenado)

	for _, num := range arregloDesordenado {
		abb.Guardar(num, num)
		require.True(t, abb.Pertenece(num), "Clave recien guardada debe pertenecer ABB")
		require.Equal(t, abb.Obtener(num), num, "Ver clave debe ser igual al dato guardado")
	}

	// Test iterador interno sin rango: los elementos deben coincidir con orden de arreglo Ordenado
	indx := 0
	abb.Iterar(func(clave int, dato int) bool {
		require.Equal(t, clave, arregloOrdenado[indx], "Iterar sin rango las claves deben salir en in-orden ordenado")
		require.Equal(t, dato, arregloOrdenado[indx], "Iterar sin rango los datos deben coincidir con clave")
		indx++
		return true
	})

}

func TestIteradorExterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	abb.Guardar(7, 7)
	abb.Guardar(6, 6)
	abb.Guardar(9, 9)
	abb.Guardar(10, 10)
	abb.Guardar(3, 3)
	abb.Guardar(5, 5)
	abb.Guardar(8, 8)

	iter := abb.Iterador()

	for iter.HaySiguiente() {
		clave, datos := iter.VerActual()
		fmt.Println("Iterador (clave, datos):", clave, datos)
		iter.Siguiente()
	}

	a, b := 6, 9
	iter = abb.IteradorRango(&a, &b)
	for iter.HaySiguiente() {
		clave, datos := iter.VerActual()
		fmt.Println("Iterador (clave, datos):", clave, datos)
		iter.Siguiente()
	}
}
