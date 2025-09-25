## Consigna
Dadas dos pilas de enteros positivos (con posibles valores repetidos) cuyos elementos fueron ingresados de menor a mayor, se pide implementar una función `func MergePilas(pila1, pila2 Pila[int]) []int` que devuelva un array ordenado de menor a mayor con todos los valores de ambas pilas sin repeticiones. Detallar y justificar la complejidad del algoritmo considerando que el tamaño de las pilas es `N` y `M` respectivamente.

## Puntos claves
- Pilas menor a mayor: `[1, 2, 3, 4, tope]`, `[2, 4, 6, 8, tope]`
- Array ordenado menor a mayor SIN REPETICIONES `[1, 2, 3, 4, 6, 8]`
- Justificar complejidad de algoritmo (Tamaño N y M)
- Función no primitiva: solo puedo utilizar Apilar, Desapilar, VerTope, EstaVacia

## Resolución
Lo que se me ocurre es: 
1. Desapilar de manera ordenada en un array (ordenado de mayor a menor, ya que pila se insertó de menor a mayor). Complejidad O(N + M).
2. Invertir array para que sea de menor a mayor. Complejidad `O(N + M)`

**Código**

```go
func MergePilas(pila1, pila2 Pila[int]) []int {
    array := make([]int, 0)
    var elementoActual int

    // Insertar en array ordenadamente
    for !pila1.EstaVacia() && !pila2.EstaVacia() {
        if pila1.VerTope() > pila2.VerTope() {
            elementoActual = pila1.Desapilar()
        } else if pila2.VerTope() > pila1.VerTope() {
            elementoActual = pila2.Desapilar()
        } else {
            elementoActual = pila1.Desapilar()
            _ := pila2.Desapilar()
        }
        array = append(array, elementoActual)
    }
    // Insertar los elementos restantes 
    for !pila1.EstaVacia() {
        elementoActual = pila1.Desapilar()
        array = append(array, elementoActual)
    }
    for !pila2.EstaVacia() {
        elementoActual = pila2.Desapilar()
        array = append(array, elementoActual)
    }
    // Invertir Array O((m+n)/2) = O(m+n)
    for i := 0; i < len(array)/2; i++ {
        array[i], array[len(array) - 1 - i] = array[len(array) - 1 - i], array[i]
    }
    // Devolver array ordenado menor a mayor
    return array
}
```
