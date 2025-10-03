package diccionario

import "fmt"

// *** Estructura hash cerrado ***

const (
	VACIO            = 0
	OCUPADO          = 1
	BORRADO          = -1
	FACT_REDIMENSION = 0.7
)

type celdaHash[K any, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K any, V any] struct {
	tabla     []celdaHash[K, V]
	capacidad int
	cantidad  int
	borrados  int
	comparar  func(K, K) bool
}

// Funciones auxiliares

func CrearHash[K any, V any](comparar func(K, K) bool) Diccionario[K, V] {
	// Inicializar el hash
	nuevoDic := hashCerrado[K, V]{
		capacidad: 100,
		cantidad:  0,
		comparar:  comparar}
	nuevoDic.redimensionarTabla(nuevoDic.capacidad)
	return &nuevoDic
}

// convertirABytes toma la clave y la convierte a bytes
func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// djb2Hash toma una clave y el numero maximo que puede valer,
// y devuelve la clave en hash dentro del rango pasado.
// documentacion http://www.cse.yorku.ca/~oz/hash.html
func djb2Hash[K any](clave K, largo int) int {
	bytes := convertirABytes(clave)
	hash := 5381
	for _, b := range bytes {
		hash = ((hash << 5) + hash) + int(b) // hash * 33 + b
	}
	return hash % largo
}

// buscarCelda toma una clave y devuelve el pntero a celda correspondiente si existe.
// En caso de no existir devuelve el puntero a la celda donde deberia existir (la proxima celda VACIA)
func (hash hashCerrado[K, V]) buscarCelda(clave K) *celdaHash[K, V] {
	// Buscar en celda actual
	claveHash := djb2Hash(clave, hash.capacidad)
	celda := &hash.tabla[claveHash]
	// Buscar en celdas siguientes si no está en actual
	for celda.estado != VACIO && !hash.comparar(clave, celda.clave) {
		claveHash++
		// Si me pase indice, iniciar de nuevo en principio de tabla
		if claveHash == hash.capacidad {
			claveHash = 0
		}
		celda = &hash.tabla[claveHash]
	}
	// devolver celda encontrada
	return celda
}

func (hash *hashCerrado[K, V]) redimensionarTabla(largo int) {
	// crear nueva tabla
	nuevaTabla := make([]celdaHash[K, V], largo)
	// Para cada celda con posicion ocupada, reubicar en nueva tabla
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			claveHash := djb2Hash(celda.clave, largo)
			nuevaCelda := &nuevaTabla[claveHash]
			nuevaCelda.clave = celda.clave
			nuevaCelda.dato = celda.dato
			nuevaCelda.estado = OCUPADO
		}
	}
	// actualizar variables luego de redimensionar
	hash.tabla = nuevaTabla
	hash.capacidad = largo
}

// Primitivas hash cerrada

func (hash hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	// Hallar celda donde deberia estar la clave
	celda := hash.buscarCelda(clave)
	// Si celda esta ocupada y las claves coinciden entonces pertenece a diccionario
	return celda.estado == OCUPADO && hash.comparar(clave, celda.clave)
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	celda := hash.buscarCelda(clave)
	// Marcar ocupado y reservar celda si estaba vacia
	if celda.estado != OCUPADO {
		celda.clave = clave
		celda.estado = OCUPADO
		hash.cantidad++
	}
	celda.dato = dato
	// Chequeo redimension del hash
	if float64(hash.cantidad+hash.borrados)/float64(hash.capacidad) >= FACT_REDIMENSION {
		hash.redimensionarTabla(hash.capacidad * 2)
	}
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	celda := hash.buscarCelda(clave)
	return celda.dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	// Actualizar celda
	celda := hash.buscarCelda(clave)
	celda.estado = BORRADO
	// Actualizar contadores
	hash.cantidad--
	hash.borrados++
	return celda.dato
}

func (hash hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.dato) {
			return
		}
	}
}

func (hash hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIter := iteradorDiccionario[K, V]{hash.tabla, hash.capacidad, 0}
	return &nuevoIter
}

// *** Estructura iterador externo ***
type iteradorDiccionario[K any, V any] struct {
	tabla       []celdaHash[K, V]
	largoTabla  int
	celdaActual int
}

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.celdaActual < iter.largoTabla
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	celda := iter.tabla[iter.celdaActual]
	return celda.clave, celda.dato
}

func (iter iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.celdaActual++
	// Se detiene en el proximo casillero ocupado y dentro del rango
	if iter.HaySiguiente() && iter.tabla[iter.celdaActual].estado != OCUPADO {
		iter.celdaActual++
	}
}
