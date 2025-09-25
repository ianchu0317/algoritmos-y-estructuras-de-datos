## Consigna
Se tiene un arreglo de N >= 3 elementos en forma de pico, esto es:
estrictamente creciente hasta una determinada posición p, y estrictamente decreciente a partir de ella (con 0 < p < N − 1). Por ejemplo, en el arreglo `[1, 2, 3, 1, 0, -2]` la posición del pico es p = 2. Se pide:

1. Implementar un algoritmo de división y conquista de orden O(log n) que encuentre la posición p del pico: `func PosicionPico(v []int, ini, fin int) int`. La función será invocada inicialmente como:  `PosicionPico(v, 0, len(v)-1)`, y tiene como pre-condición que el arreglo tenga forma de
pico.
2. Justificar el orden del algoritmo mediante el teorema maestro.


## Puntos clave
- El arreglo siempre tiene mas de 3 elementos

## Resolucion
Para este ejercicio podemos aplicar algo parecido a una busqueda binaria. El elemento que estamos buscando tiene que cumplir con la condición de que el el elemento anterior es menor y que el que sigue es también menor. 
Teniendo eso en cuenta, vamos reduciendo el arreglo teniendo en cuenta: si el elemento anterior es mayor al elemento del medio, entonces reducimos hacia la izquierda, en cambio si el emento posterior es mayor al actual entonces buscamos en subarreglo derecha. En codigo seria así

```go
func PosicionPico(v[]int, ini, fin int) int {
    mitad := (ini + fin) / 2
    // caso base
    if v[mitad-1] < v[mitad] && v[mitad] > v[mitad+1] {
        return mitad
    } else {
        if v[mitad - 1] > v[mitad] {
            return PosicionPico(v, ini, mitad-1)
        } else if v[mitad + 1] > v[mitad] {
            return PosicionPico(v, mitad+1, fin)
        } 
    }
}
```

La ecuacion de recurrencia del algoritmo es: T(n)=T(n/2)+O(1), por lo que si aplicamos teorema maestro entonces nos da que A=1, B=2, C=0 y log_2(1)=C=0, entonces la complejidad es O(log(n)), que tiene sentido ya que estamos aplicando algo similar a la busqueda binaria