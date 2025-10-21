## Consigna
Implementar en lenguaje Go una función recursiva con la firma `func esHeap(arr []int)`. Esta función debe devolver true o false de acuerdo a si el arreglo que recibe como parámetro cumple la propiedad de heap (de mínimos).


## Resolucion
Un arreglo es heap si cumple la propiedad de que para cada elemento, sus hijos (si existen) son menores o iguales al elemento actual, así lo mismo recursivamente con el hijo izquierdo y con el hijo derecho.
```
i -> elemento actual
i*2 + 1 -> hijo izquierdo
i*2 + 2 -> hijo derecho
```

Teniendo esto en cuenta, nos pide que sea función.

```go
func esHeap(arr []int) bool {
    return esHeapRecursivo(0, arr)
}   

func esHeapRecursivo(actual int, arr []int) bool {
    if actual >= len(arr) {
        return true
    }
    indxIzq := actual*2 + 1
    indxDer := actual*2 + 2

    // Si no cumple condicion es porque es false
    if indxIzq < len(arr) && arr[indxIzq] > arr[actual] {
        return false
    }

    if indxDer < len(arr) && arr[indxDer] > arr[actual] {
        return false
    }
    // Si los hijos cumplen condicion heap minimo recursivamente entonces es Heap
    return esHeapRecursivo(indxIzq, arr) && esHeapRecursivo(indxDer, arr)
}
```

La complejidad es O(n) ya que tengo que chequear para los n-nodos del arbol aplicando operaciones de O(1): comparar, asignar variables, para ver si cumple propiedad de heap