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

// obtenerValoresEsperados toma punteros *int de desde, hasta y un arreglo ordenado.
// Devuelve un arreglo ordenado con elementos que cumplen desde <= num <= hasta, donde [num1, num2..numN].
// En caso de si alguno es nil entonces toma como si no hubiera rango.
func obtenerValoresEsperados(desde, hasta *int, arregloOrdenado []int) []int {
	valoresEsperados := []int{}
	for _, num := range arregloOrdenado {
		if (desde == nil || num >= *desde) && (hasta == nil || num <= *hasta) {
			valoresEsperados = append(valoresEsperados, num)
		}
	}
	return valoresEsperados
}

func testIteradorInternoRango(t *testing.T, desde, hasta *int, arregloOrdenado []int,
	abb TDADiccionario.DiccionarioOrdenado[int, int], msg string) {
	valoresEsperados := obtenerValoresEsperados(desde, hasta, arregloOrdenado)
	i := 0
	abb.IterarRango(desde, hasta, func(clave int, dato int) bool {
		require.Equal(t, clave, valoresEsperados[i], msg)
		i++
		return true
	})
}

func TestIteradorInterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararInt)
	rango := 100
	arregloDesordenado := crearArregloDesordenado(10, rango)
	arregloOrdenado := countingSort(arregloDesordenado, rango)

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
	// Test iterador interno con rango: los elementos deben coincidir con orden de arreglo ordenado
	// respetando rango: desde <= num <= hasta
	desde := arregloOrdenado[3]
	hasta := arregloOrdenado[7]

	// Test iterador interno con nil
	testIteradorInternoRango(t, nil, nil, arregloOrdenado, abb,
		"Iterar con rango desde-hasta las claves deben salir en in-orden ordenado")
	// Test iterador interno con nil en uno de los extremos: nil-hasta
	testIteradorInternoRango(t, nil, &hasta, arregloOrdenado, abb,
		"Iterar nil-hasta las claves deben salir en in-orden ordenado")
	// Test iterador interno con nil en uno de los extremos: desde-nil
	testIteradorInternoRango(t, &desde, nil, arregloOrdenado, abb,
		"Iterar desde-nil las claves deben salir en in-orden ordenado")
}

func testIteradorExternoRango(t *testing.T, desde, hasta *int,
	arregloOrdenado []int, abb TDADiccionario.DiccionarioOrdenado[int, int], msg string) {
	valoresEsperados := obtenerValoresEsperados(desde, hasta, arregloOrdenado)
	iter := abb.IteradorRango(desde, hasta)
	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.Equal(t, valoresEsperados[i], clave, msg)
		iter.Siguiente()
		i++
	}
}
func TestIteradorExterno(t *testing.T) {

	abb := TDADiccionario.CrearABB[int, int](compararInt)
	rango := 100
	arregloDesordenado := crearArregloDesordenado(10, rango)
	arregloOrdenado := countingSort(arregloDesordenado, rango)

	for _, num := range arregloDesordenado {
		abb.Guardar(num, num)
		require.True(t, abb.Pertenece(num), "Clave recien guardada debe pertenecer ABB")
		require.Equal(t, abb.Obtener(num), num, "Ver clave debe ser igual al dato guardado")
	}

	// Iterador sin rango debe salir en mismo orden que elementos en arregloOrdenado
	iter := abb.Iterador()
	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.Equal(t, arregloOrdenado[i], clave, "Tiene que coincidir orden de iterador externo con arreglo ordenado")
		iter.Siguiente()
		i++
	}

	// ** Test iterador externo con rango **
	desde := arregloOrdenado[3]
	hasta := arregloOrdenado[7]

	// Iterador externo con rango nil-nil
	testIteradorExternoRango(t, nil, nil, arregloOrdenado, abb,
		"En rango nil-nil tiene que coincidir orden de iterador externo con arreglo ordenado")
	// Iterador externo con rango nil-hasta
	testIteradorExternoRango(t, nil, &hasta, arregloOrdenado, abb,
		"En rango nil-hasta tiene que coincidir orden de iterador externo con arreglo ordenado")
	// Iterador externo con rango desde-nil
	testIteradorExternoRango(t, &desde, nil, arregloOrdenado, abb,
		"En rango desde-nil tiene que coincidir orden de iterador externo con arreglo ordenado")
}

func testVolumen(t *testing.T, volumen int) {
	// Crear arreglo
	rango := volumen + 1000
	arregloDesordenado := crearArregloDesordenado(volumen, rango)
	arregloOrdenado := countingSort(arregloDesordenado, rango)
	abb := TDADiccionario.CrearABB[int, int](compararInt)

	// Cargar ABB
	for _, num := range arregloDesordenado {
		abb.Guardar(num, num)
	}
	// Chequear con el arreglo ordenado si los elementos estan cargados
	for _, num := range arregloOrdenado {
		require.True(t, abb.Pertenece(num), "Numero guardado tiene que pertenecer a ABB")
		require.Equal(t, abb.Obtener(num), num, "Numero guardado tiene que ser igual al elemento guardado")
	}
	require.Equal(t, volumen, abb.Cantidad(), "Cantidad de elementos tiene que ser igual al voluemn cargado")
	// Borrar elementos desde el arreglo Desordenado
	for _, num := range arregloDesordenado {
		fmt.Println(abb.Obtener(num))
		require.Equal(t, num, abb.Borrar(num), "Elemento borrado tiene que coincidir con elemento cargado")
	}
	require.Equal(t, 0, abb.Cantidad(), "Luego de borrar los elementos no debe quedar nada en arbol")
}

func TestVolumen(t *testing.T) {
	// testear con 10000, 20000, 40000
	testVolumen(t, 10000)

}
