package parcialitos

import Pila "tdas/pila"

/*Se tiene un arreglo de n cadenas y una buena función de hashing. Se quiere ordenar dichas cadenas por su resultado
en la función de hashing, habiéndole hecho previamente % K a cada uno de los resultados (donde K es un valor enorme,
muchísimo más grande que n). Implementar un algoritmo que ordene las cadenas por dicho criterio en O(n). La firma de
la función debe ser: func Ordenar(cadenas []string, valoresHash []int64, K int64) []string. valoresHash
ya tiene cada valor posterior a haber hecho % K. Recomendamos recordar las propiedades de las funciones de hashing.
Si necesitás un algoritmo de ordenamiento auxiliar al que estés implementando, podés considerarlo ya implementado.
Justificar brevemente por qué es correcta la aplicación del algoritmo que implementaste. Justificar la complejidad del
algoritmo implementado.
Por ejemplo, si queremos ordenar las cadenas: gato, perro, elefante, comadreja y los resultados de aplicarles la
función de hashing (y % K) son 19, 703, 9872, 37, respectivamente, las cadenas deben de quedar: gato, comadreja,
perro, elefante.*/

type cadenaHash struct {
	cadena string
	hash   int64
}

func Ordenar(cadenas []string, valoresHash []int64, K int64) []string {
	// Variables auxiliares a utilizar	O(1)
	numBaldes := len(cadenas)            // k = n
	elemPorBalde := K / int64(numBaldes) // b = K/n
	baldes := make([][]cadenaHash, numBaldes)

	// Insertar cada cadena en su balde. Guardar el struct para ordenar después
	// O(n*b)
	for i := range cadenas {
		cadena := cadenas[i]
		hash := valoresHash[i]
		indexBalde := hash / elemPorBalde
		baldes[indexBalde] = append(baldes[indexBalde], cadenaHash{cadena, hash})
	}

	// Ordenar cada balde O(n log (n/b))
	for i := range baldes {
		baldes[i] = mergeSort(baldes[i], func(a, b cadenaHash) int { return int(a.hash - b.hash) })
	}

	// Iterar cada balde y poner ordenado O(n)
	ordenado := make([]string, len(cadenas))
	i := 0
	for _, balde := range baldes {
		for _, _cadenaHash := range balde {
			ordenado[i] = _cadenaHash.cadena
			i++
		}
	}

	return ordenado
}

/*
Implementar un algoritmo que reciba un arreglo de enteros desordenado y un número elem que, por división y
conquista determine si elem se encuentra en el arreglo. Indicar y justificar adecuadamente la complejidad del algoritmo
implementado.
*/

func elemEnArreglo(elem int, arr []int) bool {
	// Complejidad por teorema maestro: T(n)=2T(n/2) + O(1) -> O(n)
	if len(arr) == 0 {
		return false
	} else if len(arr) == 1 {
		return elem == arr[0]
	}
	mitad := len(arr) / 2
	return elemEnArreglo(elem, arr[:mitad]) || elemEnArreglo(elem, arr[mitad:])
}

/*
Implementar una función que reciba una pila de enteros y devuelva la suma de todos los elementos. Al finalizar la
ejecución de la función, la pila debe quedar en el mismo estado que tenía antes de ejecutar la misma. La función no puede
utilizar estructuras auxiliares (incluyendo arreglos). La firma de la función debe ser pilaSumar(pila Pila[int]) int.
Indicar y justificar la complejidad de la función implementada.
*/

func pilaSumar(pila Pila.Pila[int]) int {
	if pila.EstaVacia() {
		return 0
	}
	elem := pila.Desapilar()
	suma := elem + pilaSumar(pila)
	pila.Apilar(elem)
	return suma
}

// Complejidad O(n): POrque practicamente recorro toda la pila
