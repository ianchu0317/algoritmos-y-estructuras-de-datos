
## Consigna

Dada una lista enlazada implementada con las siguientes estructuras:
```go
     type nodoLista[T any] struct {
         prox *nodoLista[T]
         dato T
     }
     type ListaEnlazada[T any] struct {
         prim *nodoLista[T]
     }
```

Escribir una primitiva de lista que devuelva el elemento que esté a `k` posiciones del final (el ante-k-último), recorriendo la lista una sola vez y sin usar estructuras auxiliares. Considerar que k es siempre menor al largo de la lista. Por ejemplo, si se recibe la lista `[ 1, 5, 10, 3, 6, 8 ]`, y `k = 4`, debe devolver `10`. Indicar el orden de complejidad de la primitiva.

## Resolucion

La estructura que nos propone el ejercicio es una lista enlazada que sólo tiene en cuenta el primer elemento, por lo que no se puede acceder de último hacia adelante como lo haría en un arreglo. 

Una solucion que se me ocurre para este problema es recorrer la lista dos veces: la primera vez que se recorre, contar la cantidad de elementos en la lista. Con esa información sabríamos en qué posición está lo que busco. Y la segunda vez realizar para buscar el dato. Pensando así, cada recorrido sería O(n) por lo que sería 2*O(n) pero como se desprecia la constante entonces es O(n). 

Otra forma que había llegado a pensar es utilizar recursión pero estaba pensando como un array creo. La consigna menciona que la función tiene que ser una primitiva  

---

## Rehacer

Podria hacer dos punteros. Uno (puntero 1) que lleve la cuenta del nodo actual y el otro (puntero 2) que también esté llevando la cuenta pero que comienze desde inicio cuando hayan pasado k elementos desde que arracó el primer puntero. Luego si `puntero1.siguient == nil` entonces obtengo el objeto. Este algoritmo también es O(n) ya que depende del tamaño de la lista, pero lo hace recorriendo una sola vez 

```go
func (lista ListaEnlazada[T]) HallarK(k int) T {
    p1 := lista.prim
    p2 := lista.prim
    contadorP1 := 1 // primer nodo

    // avanzar punteros
    for p1.prox != nil {
        p1 = p1.prox
        contadorP1 += 1

        // no incluir contadorP1 == k porque ya empieza en 1
        if contadorP1 > k {
            p2 = p2.prox
        }
    }

    if contadorP1 >= k {
        return p2.dato
    } else {
        panic("K inválido")
    }

}
```