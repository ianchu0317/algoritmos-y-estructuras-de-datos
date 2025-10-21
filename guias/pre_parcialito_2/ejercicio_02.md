## Consigna
Implementar un algoritmo que reciba un arreglo desordenado de n enteros y un número K y determinar en O(n) si existe un par de elementos en el arreglo que sumen exactamente K.

## Resolucion 

- Recorrer cada numero en lista
- Para cada numero verificar si ya recorri el complemento (clave tiene que ser o(1))
    - Si ya lo recorri, return true
    - Si no lo recorri, seguir iterando pero agregar numero actual a mi lista de iterados

```go
func existeSumaK(k int, arr []int) bool {
    recorridos := CrearHash[int, int]()
    // Recorre n-veces siendo n la cantidad de elementos en lista
    for _, num := range arr {
        complemento := k - num
        if recorridos.Pertenece(complemento) {
            return true
        }
        recorridos.Guardar(num, num)
    }
    return false
}
```



