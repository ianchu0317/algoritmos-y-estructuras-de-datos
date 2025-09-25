## Consigna
Implementar la función `func Multiprimeros[T any](cola Cola[T], k int) []T` con el mismo comportamiento de la primitiva anterior.


## Resolucion
Pide realizar lo mismo pero como función, por lo que no hay acceso para la estructuras internas pero sí podemos utilizar las primitivas de la cola. 
Lo que se me ocurre es desencolar k elementos utilizando la primitiva `cola.Desencolar()` e ir agregando a un slice. Hay que ir desencolando si contador < k y cola no está vacía `!cola.EstaVacia()`. No hace falta luego reestructurar la cola como el ejercicio anterior ya que lo hace internamente el tda. 

```go
func Multiprimeros[T any](cola Cola[T], k int) []T {
    // declarar variables a utilizar
    contador := 0
    listaElementos := make([]T, 0)
    var elementoActual T

    // desencolar k elementos si no está vacía
    for !cola.EstaVacia() && contador < k {
        elementoActual = cola.Desencolar()
        listaElementos = append(listaElementos, elementoActual)
        contador += 1
    }

    return listaElementos
}
```
Esta operación cuesta lo mismo que el anterior. Cada primitiva del TDA cola es O(1) repetidos k veces, que depende de la cantidad de elementos que se utilza para llamar la función, y no depende de la cantidad de elemtnos en la cola. Las declaraciones son O(1). Finalmente entonces la operación es O(K) como la primitiva del ejercicio anterior