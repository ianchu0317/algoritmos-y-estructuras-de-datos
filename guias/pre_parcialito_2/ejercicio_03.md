## Consigna
Escribir en Go una función k-merge que reciba k arreglos ordenados de h elementos cada uno (en total tenemos n = k*h elementos) y devuelva uno nuevo con todos los n elementos ordenados. La complejidad tiene que ser O(nlog(k))

## Resolucion

- Cargar heap con primera columna de la matriz (primer elemento de cada arreglo)
- Mientras heap no este vacia
    - Ir Desencolar el heap y agregar al arreglo a devolver (elemento minimo de todos) y encolar el elemento proximo de ese elemento de su arreglo (si tiene) **O(log(k))** -> eventualmente el heap se va a quedar cada vez con menos elementos


- Alternativa: ordenar el arreglo O(n log n). Osea pero si K < H entonces es mucho mas eficiente el otro metodo


```go

package kmerge

type elemento struct {
    dato int
    arreglo int
    posicion int
}

func cmpMin(a, b elemento) int {
    return b.dato - a.dato
}

func KMerge(arr [][]int) []int {
    // Declaraciones de variables O(1)
    k := len(arr)
    h := len(arr[0])
    resultado := make([]int, 0, k*h)
    heap := CrearHeap[elemento](cmpMin) 
    // Cargar el heap con la primera columna O(log K)
    for i := 0; i < k; i++ {
        e := elemento{arr[i][0], i, 0}
        heap.Encolar(e)
    }
    // Ir desencolando y agregar a resultado. N * log (k)
    for !heap.EstaVacia() {
        minimo := heap.Desencolar() // O(log K)
        resultado = append(resultado, minimo.dato)  // O(1)
        // Encolar el proximo elemento del arreglo si hay  
        // O(log K) por encolar
        if minimo.posicion < h - 1{
            proxIndx := minimo.posicion + 1
            proximo := elemento{arr[minimo.arreglo][proxIndx], minimo.arreglo, proxIndx}
            heap.Encolar(proximo)
        }
    }
    // devolver resultado una vez que el heap quedo vacio
    return resultado
}

// Complejidad total: O(1) + O(log k) + O(n log K) + O(1) = O(n log K)

```