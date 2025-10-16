## Consigna

Escribir una función en Go que, dado un arreglo de `n` cadenas y un entero positivo `k` cadenas más largas. Se 
espera que el orden del algoritmo sea O(n+klogn). Justificar el orden.

## Resolucion

```go
package cadenaslargas

func CadenasLargas(arr []string, k int) []string {
    // Crear heap desde arreglo O(n)
    heapPalabras := CrearHeapArr[string](arr, func(a, b string) int {
        if len(a) > len(b) {
            return 1 
        } 
        if len(b)> len(a) {
            return -1
        }
        return 0
    })

    kMasLargas := make([]string, 0)
    contador := 0

    // Ir al si no llegue K veces o si no esta vacia
    for contador < k && !heapPalabras.EstaVacia() {
        kMasLargas = append(kMasLargas, heapPalabras.Desencolar())
        contador++
    }
    return kMasLargas
}

// Cuando creo desde array, realizo una vez heapify: O(N)
// Luego cuando desencolo K veces es K*log(n) ya que cada accion de desencolar es O(log n) porque se realiza una vez downheap para el nuevo primer elemento

```