package parcialitos

/*
1. Implementar una primitiva eliminarColisiones(clave K) []K para el hash, que elimine del hash todas las
claves que colisionen con la clave pasada por parámetro en el estado actual (sin eliminar dicha clave del
diccionario, si se encuentra), y devuelva dichas claves. Implementar tanto para el hash abierto como para el hash
cerrado. Si no se implementa para alguno, el ejercicio no estará aprobable. Indicar y justificar la complejidad de
la primitiva para ambos casos.
*/
const (
	BORRADO = iota
	OCUPADO
	VACIO
)

type celda[K any, V any] struct {
	estado int
	clave  K
	valor  V
}

type hashCerrado[K any, V any] struct {
	tabla    []celda[K, V]
	cantidad int
}

type hashAbierto[K any, V any] struct {
	tabla    []Lista[celda]
	cantidad int
}

func (hash *hashCerrado[K, V]) eliminarColisiones(clave K) []K {
	claveHasheado := hash.Hashear(clave)
	borrados := make([]K, 0, hash.cantidad)

	for i := 0; i < len(hash.tabla); i++ {
		celda := &hash.tabla[i]
		if celda.clave != clave && celda.estado == OCUPADO && hash.Hashear(celda.clave) == claveHasheado {
			borrados = append(borrados, celda.clave)
			hash.cantidad--
			celda.estado = BORRADO
		}
	}
	return borrados
}

func (hash *hashAbierto[K, V]) eliminarColisiones(clave K) []K {
	claveHasheado := hash.Hashear(clave)
	borrados := make([]K, 0, hash.cantidad)

	for iter := hash.tabla[claveHasheado].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		celda := iter.VerActual()
		if celda.clave != clave {
			iter.Borrar()
			hash.cantidad--
			borrados = append(borrados, celda.clave)
		}
	}

	return borrados
}

// Complejidad O(n) -> itero n elementos del hash buscando colisiones -> Operaciones o(1) hashear, etc...

/*
2. Sobre un AVL cuyo estado inicial puede reconstruirse a partir del preorder: 40 - 10 - 3 - 17 - 15 - 64 -
47 - 74 - 92, realizar un seguimiento de insertar los siguientes elementos (incluyendo rotaciones intermedias):
20, 23, 13, 14, 16, 12.
*/

/*
3. Implementar una función mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] que, dado un arre-
glo de Alumnos y un valor entero k, nos devuelva una lista de los k alumnos de mayor promedio (ordenada de

mayor a menor). Indicar y justificar la complejidad del algoritmo implementado.
Considerar que la estructura del alumno es:
type Alumno struct {
nombre string
padron int
notas []int
}
*/

type Alumno struct {
	nombre string
	padron int
	notas  []int
}

func prom(notas []int) int {
	sum := 0
	for _, num := range notas {
		sum += num
	}
	return sum / len(notas)
}

func mejoresPromedios(alumnos []Alumno, k int) Lista[Alumno] {
	// Crear heap desde arreglo, se aplica heapify O(n)
	heap := CrearHeapArr(alumnos, func(a, b Alumno) int { return prom(a.notas) - prom(b.notas) })
	// Crear lista para resultados O(1)
	listaAlumnos := CrearListaEnlazada[Alumno]()
	// K veces -> desencolar heap -> K log(n). OJO K <= N
	for i := range k {
		listaAlumnos.InsertarUltimo(heap.Desencolar())
	}
	return listaAlumnos
}

// COmplejidad total: O(n) + O(k log(n)) = O(n + klog(n))

// Si la cantidad de notas es acotado O(c) entonces heapify es O(n*c), si c == 3, entonces heapify es O(n)
