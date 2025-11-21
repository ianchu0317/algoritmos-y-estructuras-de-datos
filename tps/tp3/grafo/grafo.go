package grafo

type Grafo[T any] interface {
	// AgregarVertice(v) toma un vertice 'v' y lo agrega al grafo
	AgregarVertice(v T)

	// AgregarArista(v, w) toma vertices 'v' y 'w' y los une.
	// Si es dirigido, une v -> w
	// Si no es dirigido, une v -> w, w -> v
	AgregarArista(v, w T)

	// Adyacentes(v) toma un vertice 'v'
	// y devuelve un arreglo con sus aristas
	Adyacentes(v T) []T

	// PesoArista(v, w) toma dos vertices 'v' y 'w'
	// Y devuelve el peso de la arista que los une
	// En caso de no existir arista devuelve panic "No existe arista"
	PesoArista(v, w T) int

	// HayArista(v, w) toma dos vertices y devuelve:
	// True en caso de existir arista entre 'v' y 'w', False en caso contrario
	HayArista(v, w T) bool

	// HayVertice(v) toma un vertice y devuelve True si vértice es parte del grafo, False en caso contrario
	HayVertice(v T) bool

	// BorrarVertice(v) toma un vertice y lo borra del grafo
	// En caso de no existir devuelve panic "No existe vertice"
	BorrarVertice(v T)

	// BorrarArista(v, w) toma dos vértices y borra el arista que los conecta.
	// En caso de no existir arista devuelve panic "NO existe arista"
	BorrarArista(v, w T)

	// ObtenerVertices devuelve un arreglo con todos los vertices del grafo
	ObtenerVertices() []T

	// VerticeAleatorio() devuelve un vértice cualquiera del grafo
	VerticeAleatorio() T

	// Cantidad() devuelve la cantidad de vértices en el grafo
	Cantidad() int
}
