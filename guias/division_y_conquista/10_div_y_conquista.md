## Consigna
Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado (todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden.


## Puntos clave
- Sin elementos repetidos
- Arreglo casi ordenado menos un elemento
- Indicar y justificar orden

## Resolucion y planificacion
En un principio, no nos preocupemos si el largo de arreglo es 1 o 0 ya que es lo mas simple. El caso base en este caso sería si el arreglo posee dos elementos: [1, 2]. Para chequear que sea ordenado solo tengo que comparar esos dos elementos. En caso de que arr[0] > arr[1] devolver arr[0].

Entonces, podemos dividir el arreglo en dos partes: una mitad izquierda y otra mitad derecha, utilizamos la función anterior para chequear si está ordenado o no. Si una mitad está ordenada, entonces la descartamos y trabajamos con la otra mitad, así sucesivamente hasta que nos quede dos elementos.

```go
func elementoFueraDeOrden(arr []int) int {
    if len(arr) == 2 {
        arr[0] > arr[1]
    }
    
    mitad := len(arr) / 2
    izqOrdenada := estaOrdenado(arr[:mitad])
    derOrdenada := estaOrdenado(arr[mitad:])
    
    if izqOrdenada {
        return elementoFueraDeOrden(arr[mitad:])
    } else {
        return elementoFueraDeOrden(arr[:mitad])
    }
}
```