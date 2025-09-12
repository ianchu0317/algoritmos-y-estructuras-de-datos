
## Consigna

Implementar la primitiva `func (cola *colaEnlazada[T]) Multiprimeros(k int) []T` que dada una cola y un número k, devuelva los primeros k elementos de la cola, en el mismo orden en el que habrían salido de la cola. 
En caso que la cola tenga menos de k elementos. Si hay menos elementos que k en la cola, devolver un slice del tamaño de la cola. Indicar y justificar el orden de ejecución del algoritmo.

## Resolución

Como es una primitiva, tenemos acceso directamente a los componentes internos de la estructura. De hecho, evitar utilizar otras primitivas. Podría llevar un contador y un slice para ir desencolando los primeros elementos de la cola



```go

// cola: 1, 2, 3, 4
// cont: 0, 1, 2, 3, K=2


func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {
    // variables para iterar
    actual := cola.primero
    var anterior *nodoCola
    contador := 0
    listaElementos := make([]T, 0)
    
    // iterar la cola, siempre cuando sea menor a k elementos y no llegado al final
    for contador < k && actual != nil {
        // guardar el dato del nodo actual
        listaElementos = append(listaElementos, actual.dato)
        // pasar al siguiente nodo
        anterior = actual
        actual = actual.siguiente
        contador += 1
    }

    // actualizar el primer puntero
    cola.primero = actual
    // actualizar ultimo puntero si esta vacio
    if actual == nil {
        cola.ultimo = nil
    } 

    return listaElementos
}
```
Este algoritmo depende de k así que es O(k). No depende del tamaño de la cola así que no es O(n), siendo n el tamaño de la cola.