## Consigna
Implementar en Go una primitiva `func (lista *ListaEnlazada) Invertir()` que invierta la lista, sin utilizar estructuras auxiliares. Indicar y justificar el orden de la primitiva.

**Condiciones clave**
- Invertir lista
- No utilizar estructuras aux
- Indicar complejidad algoritmica

## Resolucion
Podemos suponer que existe una declaracion así como estructura de Lista
```go
type nodoLista[T any] struct{
    dato T
    siguiente *nodoLista
}

type ListaEnlazada[T any] struct {
    primero *nodoLista[T]
    ultimo *nodoLista[T]
}
```
Teniendo eso en cuenta, se tiene que trabajar con la lógica interna sin utilizar auxiliares.

Lo que se me ocurre es iterar por la lista hasta el final. En cada paso reubicar los nodos tal que el elemento siguiente pase apuntando al actual así sucesivamente hasta que el actual sea nil o que el proximo sea nil. Se me ocurre utilizar tres punteros. Uno para actual, otro para el anterior y el otro para siguiente

```go 
func (lista *ListaEnlazada[T]) Invertir() {
    actual := lista.primero
    var anterior *nodoLista = nil

    for actual != nil {
        siguiente := actual.siguiente
        actual.siguiente = anterior
        anterior = actual
        actual = siguiente
    }

    // actualizar lista
    lista.ultimo = lista.primero
    lista.primero = anterior
}
```

El algoritmo es de complejidad O(n) ya que recorre toda la lista de tamaño N y las operaciones de inversión son O(1)

