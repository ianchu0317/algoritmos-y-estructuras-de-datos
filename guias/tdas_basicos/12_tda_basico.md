## Consigna
Implementar una función `func FiltrarCola[K any](cola Cola[K], filtro func(K) bool)` , que elimine los elementos encolados para los cuales
la función filtro devuelve false. Aquellos elementos que no son eliminados
deben permanecer en el mismo orden en el que estaban antes de invocar a la
función. No es necesario destruir los elementos que sí fueron eliminados. Se
pueden utilizar las estructuras auxiliares que se consideren necesarias y no está permitido acceder a la estructura interna de la cola (es una función). ¿Cuál es el orden del algoritmo implementado?


## Resolucion

Lo que se puede hacer es ir desencolando y encolando en otra cola auxiliar para los elementos que pasen el filtro y finalmente encolarlos en la cola dada

```go
func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) {
    colaAux := Cola[K]{}
    var elementoActual K

    // Filtrar a cola auxiliar los elementos
    for !cola.EstaVacia() {
        elementoActual = cola.Desencolar()
        if filtro(elementoActual) {
            colaAux.Encolar(elementoActual)
        }    
    }
    // Pasar de cola auxiliar a cola de parametro
    for !colaAux.EstaVacia() {
        elementoActual = colaAux.Desencolar()
        cola.Encolar(elementoActual)
    }
}
```