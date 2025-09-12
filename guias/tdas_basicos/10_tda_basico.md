## Consigna

Se quiere implementar un TDA ColaAcotada sobre un arreglo. Dicho TDA
tiene un espacio para k elementos (que se recibe por parámetro al crear la
estructura). Explicar cómo deberían implementarse las primitivas encolar y
desencolar de tal manera que siempre sean operaciones de tiempo constante.

## Resolucion

Estaríamos hablando de algo así, donde hay un espacio para datos, y llevo registro del primer elemento y ultimo elemento. 
```go
type ColaAcotada[T any] struct{
    datos []T
    primero int
    final int
}
```
[1, 2, 3, 4, 5, , , , ], k espacio

Si quiero desencolar entonces simplemente muevo el indice del primer elemento +1 si quiero encolar entonces aumento final en 1. Si al inicio del arreglo hay espacio y el indice del final está al final puedo reutilizar y hacer que utilice el espacio necesario. De esta manera no hay que agrandar ni achicar el arreglo haciendo el algoritmo O(n), sino que cada operacion de asignacion de variable en cada posicion de arreglo es una operacion de O(1). El problema con esto es que no puedo encolar más de k elementos así que hay que tener cuiadado con esto y capaz llevar un contador externo cuando se utilice la estructura.