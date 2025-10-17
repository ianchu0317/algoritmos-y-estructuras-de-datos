### Consigna
Se tienen k arreglos de enteros previamente ordenados y se quiere obtener
un arreglo ordenado que contenga a todos los elementos de los k arreglos. Sabiendo que cada arreglo
tiene tamaño h, definimos como n a la sumatoria de la cantidad de elementos de todos los arreglos,
es decir, n = k \times h.

Escribir en Go una función func KMerge(arr [][]int) que reciba los k arreglos y
devuelva uno nuevo con los n elementos ordenados entre sí. La función debe ser de orden
O(n log k). Justificar el orden del algoritmo propuesto.

```go
package kmerge

// Para [0, 1, 2...h1]
//      [2, 3, 5...h2]
// Para quiero recorrer los h*K elementos de la matriz
// Recorrer la matriz por columna. 
// Para cada columna crear un HEAP O(k)
// Luego ir desencolando heap para el resultado

func compararMin(a, b int) int {
    return b-a
}

func KMerge(arr [][]int) []int {
    resultado := make([]int, 0)
    // Recorre H veces, siendo H la cantidad de columnas
    for columna := 0; columna < len(arr); columna++ {
        heap := CrearHeap[int](compararMin)
        // Meter todos los elementos de columna a heap K*log(K)
        for fila := range arr[0] {
            heap.Encolar(arr[fila][columna])
        }
        // Sacar todos los elementos de heap e insertar a resultado K*log(K)
        for fila := range arr {
            resultado = append(resultado, heap.Desencolar())
        }
    }
    return resultado
}
```