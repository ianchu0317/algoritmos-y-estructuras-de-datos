## Consigna

Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio y una clave final fin, se devuelva una lista con todos los datos cuyas claves estén entre inicio y fin, que estén dentro de los primeros M niveles del árbol (considerando a la raíz en nivel 1). Indicar y justificar la complejidad temporal.


## Resolucion
Podemos pensar como es una primitiva, podemos pensar algo así `func (arbol *abb[K, V]) primerosDatosEntre(m, inicio, fin int) Lista[V]`. Donde tiene que devolver entre los niveles m y los valores que cumplen`inicio<=valor<=fin`.
No nos especifica con que orden resolver así que podemos hacer como querramos. En mi caso lo voy a hacer in-order.

```go
type nodo[K comparable, V any] struct {
    izq *nodo[K, V]
    der *nodo[K, V]
    clave K
    dato V
}
type abb[K comparable, V any] struct {
    raiz *nodo[K, V]
}

func (arbol *abb[K, V]) primerosDatosEntre(m int, inicio, fin K) Lista[V] {
    res := CrearListaEnlazada[V]()
    arbol.raiz.iterarNodosEntre(1, m, inicio, fin, res)
    return res
}

func (nodo *nodo[K, V]) iterarNodosEntre(nivel, m int, inicio, fin K, lista Lista[V]) {
    if nodo == nil || nivel > m {
        return 
    }
    // Iterar in-order
    // Si es mayor al inicio entonces ir a la izquierda (puede haber elementos mayores a inicio)
    if nodo.clave > inicio {
        nodo.izq.iterarNodosEntre(nivel+1, m, inicio, fin, lista)
    }
    // Si nodo actual cumple condicion entonces agregar a lista
    if nodo.clave >= inicio && nodo.clave <= fin {
        lista.InsertarUltimo(nodo.dato)
    }
    // Si es menor a fin entonces ir derecha, puede haber elementos menores a fin
    if nodo.clave < fin {
        nodo.der.iterarNodosEntre(nivel+1, m, inicio, fin, lista)
    }
}
```
La complejidad es O(n), siendo n la cantidad de nodos que tienen el arbol. En el peor de los casos visito todos los nodos del arbol aplicando operaciones de O(1) (comparacion, insercion a listas); por lo que es N*O(1)=O(n). tener en cuenta que M también limita los niveles. Si M=Log(n) y todos los nodos estan dentro del rango entonces la complejidad es O(n). 
