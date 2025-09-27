## Consigna
Implementar, por división y conquista, una función que dado un arreglo y su largo, determine si el mismo se encuentra ordenado. Indicar y justificar el orden.



## Resolucion

Podemos pensar que un arreglo está ordenado si la mitad izquierda del arreglo está ordenado y lo mismo con la mitad derecha. Así recursivamente hasta que arreglo tenga dos elementos. 

Caso de que arreglo tiene 1 elemento o 0 entonces está ordenado. 

Si tiene 2 elementos entonces está ordenado si arr[0] < arr[1] (depenede de si está ordenado ascendentemente o descendentemente).


La implementacion en Go sería así
```go
func estaOrdenado[T any](arr []T) bool {
    // si arreglo tiene menos de un elemento
    if len(arr) <= 1 {
        return true
    } 
    // si arreglo tiene dos elementos
    if len(arr) == 2 {
        return arr[0] <= arr[1]
    }
    // caso general
    mitad := len(arr) / 2
    izqOrdenado := estaOrdenado(arr[:mitad])
    derOrdenado := estaOrdenado(arr[mitad:])
    return izqOrdenado && derOrdenado
}
```
De este algoritmo a continuación calculamos complejidad, pero primero hallamos la ecuación de recurrencia `T(n)= 2T(n/2) + O(1)`, por lo que A=2, B=2, C=0 y como log_b(A)=1 > C=0 entonces el algoritmo tiene complejidad O(N) y tiene sentido ya que es lo mismo que un algoritmo que avanza linealmente y chequea si el elemento actual es mayor que al anterior, por lo que al menos debe haber recorrido una vez todos los elementos N, siendo N el tamaño del arreglo.

--- 

El problema con la primera resolución es que el arreglo izquierdo puede estar ordenado y el derecho también pero no se conectan. Por ejemplo [4, 5, 6, 1, 2, 3]. No se conecta. 

Releyendo la consigna descubrí que tenemos el largo del arreglo pasado por parámetro así que la firma sería `func estaOrdenado(arr []int, largo)`

```go
func estaOrdenado[T any](arr []T, largo int) bool {
    // si arreglo tiene menos de un elemento
    if largo <= 1 {
        return true
    } 
    // si arreglo tiene dos elementos
    if largo == 2 {
        return arr[0] <= arr[1]
    }
    // caso general
    mitad := largo / 2
    izqOrdenado := estaOrdenado(arr[:mitad], mitad)
    derOrdenado := estaOrdenado(arr[mitad:], largo-mitad)
    // verificar el elemento del medio
    conexion := arr[mitad-1] <= arr[mitad]
    return izqOrdenado && derOrdenado && conexion
}
```
La complejidad sigue siendo la misma