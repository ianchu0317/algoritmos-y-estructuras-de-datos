package ejercicios

/*
Swap intercambia dos valores enteros.
*/
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

/*
Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0.
Si el máximo elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
*/
func Maximo(vector []int) int {
	// Caso base sin arreglo es largo 0
	if len(vector) == 0 {
		return -1
	}
	// Resolver con un for lineal
	max_num := vector[0]
	pos_max_num := 0
	for pos, valor := range vector {
		if valor > max_num {
			max_num = valor
			pos_max_num = pos
		}
	}
	return pos_max_num
}

/*
Comparar compara dos arreglos de longitud especificada.

	Devuelve -1 si el primer arreglo es menor que el segundo;
	0 si son iguales;
	o 1 si el primero es el mayor.

Un arreglo es menor a otro cuando al compararlos elemento a elemento,
el primer elemento en el que difieren no existe o es menor.
*/
func Comparar(vector1 []int, vector2 []int) int {
	// Hallar largo del vector mas corto
	var len_vector_corto int
	if len(vector1) < len(vector2) {
		len_vector_corto = len(vector1)
	} else {
		len_vector_corto = len(vector2)
	}
	// Comparar elementos de los vectores
	for i := range len_vector_corto {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	// Si sale ciclo for (elemntos iguales), comparar por largo
	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	}
	// Si llega hasta acá los vectores son iguales
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector)-1; i++ {
		pivot := i
		for j := i + 1; j < len(vector); j++ {
			if vector[j] < vector[pivot] {
				pivot = j
			}
		}
		if pivot != i {
			Swap(&vector[i], &vector[pivot])
		}
	}
}

/*
Suma devuelve la suma de los elementos de un arreglo.
En caso de no tener elementos, debe devolver 0.
Esta función debe implementarse de forma RECURSIVA.
Se puede usar una función auxiliar (que sea la recursiva).
*/
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo.
// Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA.
// Se puede usar una función auxiliar (que sea la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) <= 1 {
		return true
	}

	if cadena[0] != cadena[len(cadena)-1] {
		return false
	}

	return EsCadenaCapicua(cadena[1 : len(cadena)-1])
}
