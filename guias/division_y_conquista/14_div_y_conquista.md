## Consigna

Se quiere implementar MergeSort pero, en vez de dividir en dos partes el arreglo, dividirlo en tres, llamando recursivamente al algoritmo para cada una de las partes y luego uniéndolas.
1. Suponiendo que el merge de las tres partes se realiza en tiempo lineal, calcular el orden del algoritmo.
2. Si en vez de dividir en tres partes, se dividiera el arreglo en n, siendo n la cantidad de elementos del arreglo ¿a qué otro algoritmo de ordenamiento se asemeja esta implementación? ¿Cuál es el orden de dicho algoritmo?
3. Dado lo obtenido en los puntos anteriores ¿tiene sentido implementar MergeSort con k separaciones, para k > 2?

## Resolucion
1. Para el primer punto directamente desarrollo un merge sort que me debería salir ya automaticamente

```go
func merge(arr1, arr2, arr3 []int) []int {
    var i, j, k int
    i = 0
    j = 0
    k = 0

    nuevoArray := make([]int, 0)
    for i < len(arr1) && j < len(arr2) && k < len(arr3) {
        if 
    } 
}

func mergeSort(arr []int) {
    // caso base, cuando hay un solo elemento
    if len(arr) == 1 {
        return arr[0]
    }
    // Definir los subarray O(1)
    pivot1 := len(arr)/3
    pivot2 := pivot1 * 2

    arr1 := arr[0:pivot1]
    arr2 := arr[pivot1:pivot2]
    arr3 := arr[pivot2:len(arr)]

    // Ordenar subarray 3T(n/3)
    mergeSort(arr1)
    mergeSort(arr2)
    mergeSort(arr3)
    // Insertar los subarray Ordenadamente O(n)
    arrayOrdenado := merge(arr1, arr2, arr3)
    return arrayOrdenado
}
```

1. Suponiendo que el merge se realiza de forma lineal, entonces el algoritmo tomaria ecuacion de recurrencia `T(n)=3T(n/3) + O(n)`. Para hallar la complejidad si aplicamos teorema maestro (A=3, B=3, C=1) nos quedaria log_3(3) = C = 1 por lo que la complejidad sigue siendo O(nlog(n)) como un merge con dos elementos

2. Si aplicamos división entre n elementos (es decir subarrays de un elementos), entonces lo que estamos haciendo es parecido a un bucket sort de un elemento en cada bucket. La ecuación de recurrencia en este caso sería de T(n)=n*t(n/n) + O(n), y consguimos lo mismo aplicando teorema maestro (A=n, B=n, C=1), log_n(n)=1, por lo que la complejidad sigue siendo O(nlogn) -> es parecido a un selection sort. chequea elemento en la que deberia estar en la posicion actual.

3. Dado con lo obtenido en los anteriores caso no tiene sentido hacer divisiones k>2 ya que solo se complica mas en la implementacion y el algoritmo mantiene la misma complejidad