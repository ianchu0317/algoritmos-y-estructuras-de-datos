package parcialitos

/*
Dados dos arreglos ordenados A y B, donde B tiene “un elemento menos que A”, implementar un algoritmo de división y
conquista que permita obtener el valor faltante de A en B. Ejemplo, si A = {2, 4, 6, 8, 9, 10, 12} y B = {2, 4,
6, 8, 10, 12}, entonces la salida del algoritmo debe ser o bien la posición 4, o el valor 9 (lo que decidan que devuelva).
Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

// ElemFaltante toma dos arreglos ordenados y devuelve indice del elemento faltante del otro
func elemFaltante(A, B []int) int {
	return _elemFaltante(A, B, 0, max(len(A), len(B))-1)
}

// Complejidad O(log(n)), T(n)= T(n/2) + O(1), como log2(1)=0==0, entonces O(n⁰*log(n))= O(log(n))
func _elemFaltante(A, B []int, ini, fin int) int {
	// Devuelve el inidce del elemento faltante
	if ini > fin {
		return -1
	} else if ini == fin {
		return ini
	}

	mitad := (ini + fin) / 2

	if A[mitad] == B[mitad] {
		return _elemFaltante(A, B, mitad+1, fin)
	} else {
		return _elemFaltante(A, B, ini, mitad)
	}
}

/*
Realizar el seguimiento de aplicar BucketSort al siguiente conjunto de equipos de fútbol, ordenando por la cantidad
de descensos que tienen (entre paréntesis se indica la cantidad en cada caso). Para este caso aplicar la versión de
BucketSort que trabaja con valores discretos con baldes unitarios (no por rangos). Implementar dicho algoritmo, e
indicar y justificar la complejidad del mismo.
Olimpo (4) - Boca (0) - Almagro (3) - Rosario Central (4) - Banfield (8) -
Sarmiento (2) - Defensa y Justicia (0) - Plantense (3) - River (1) - Independiente (1) -
Estudiantes LP (2) - Racing (1) - Tigre (8) - Velez (1) - Atlanta (4) - Gimnasia LP (5)
*/

/*
4. Dado un arreglo de enteros ordenado de n elementos en el cual sus elementos van de 0 a M, con M  n, implementar
una función que determine en O(log n) si hay algún elemento que aparezca más de la mitad de la veces en el arreglo.
Justificar la complejidad del algoritmo implementado.
*/

/*
5. Implementar un algoritmo que permita ordenar cronológicamente un arreglo de cadenas que representan horarios en
formato HH:MM:SS. Indicar y justificar la complejidad del algoritmo implementado.
*/

/*
Tenemos un arreglo de números de 1 a n, ordenado. A dicho arreglo se le quita un elemento. Implementar un algoritmo
que determine qué elemento falta en el arreglo. Indicar y justificar la complejidad del algoritmo implementado.
*/

/*
Implementar para la cola enlazada la primitiva Consumir(accion func (T)) que aplique la función accion a todos
los elementos de la cola. Al terminar la ejecución, la cola debe quedar vacía. Se espera que se implemente sin utilizar
otras primitivas, para demostrar el conocimiento sobre estructuras enlazadas. Indicar y justificar la complejidad de la
primitiva.
*/

/*
1. Implementar una función func minimoExcluido(arr []int) int que dado un arreglo de valores enteros (mayores o iguales a 0),
obtenga el mínimo valor que no se encuentre en el arreglo. Indicar y justificar la complejidad del algoritmo (explicar en detalle este
paso, porque es fácil que se te puedan pasar detalles importantes a explicar). ¿Es el mismo ejercicio del parcialito? Si.
Por ejemplo:
minimoExcluido([]int{0, 5, 1}) --> 2
minimoExcluido([]int{3, 5, 1}) --> 0
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2}) --> 6
minimoExcluido([]int{0, 5, 1, 3, 4, 1, 2, 12345675433221345}) --> 6
*/

/*
Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede
formar con dichos dígitos. Indicar y justificar la complejidad del algoritmo implementado.
*/

/*
Implementar un algoritmo que dado un arreglo de números, determine si hay un elemento dentro del mismo que
aparece al menos la mitad veces. La complejidad del algoritmo debe ser lineal. Justificar la complejidad del algoritmo
implementado.
*/
