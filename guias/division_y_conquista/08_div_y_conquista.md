## Consigna
Implementar, por división y conquista, una función que determine el mínimo de un arreglo. Indicar y justificar el orden.

## Resolucion

Se me ocurre dividir el arreglo en 2. Buscar el mitad izquierdo y mitad derecho los minimos y luego compararlos. Realizar esto recursivamente. EL caso base es que si el array tiene un solo elemento, entonces ese es el minimo.

ˋˋˋgo
func minimo(arr []int) int {
    if len(arr) == 1 {
        return arr[0]
    }
    mitad := len(arr)/2
    min_izq := minimo(arr[0:mitad])
    min_der := minimo(arr[mitad:len(arr)])
    if min_izq min_der {
        return min_izq
    } else {
        return min_der
    }
}
ˋˋˋ

Con esta funcion si aplicamos teorema maestro podemos obtener que la complejidad es de O(n log n)