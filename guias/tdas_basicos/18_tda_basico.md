## Consigna

Implementar una función recursiva que reciba una pila y devuelva, sin utilizar estructuras auxiliares, la cantidad de elementos de la misma. Al terminar la ejecución de la función la pila debe quedar en el mismo estado al original.

## Resolucion 
### Puntos clave
- Función recursiva -> caso base y caso general -> NO es primitiva
- Sin estructuras auxiliares
- Pila intacta luego de ejecucion
- Devolver cantidad de elementos

### Idea
Llevar un contador que va incrementando por cada llamado de recursividad. Desapilar hasta la pila vacía (return 0). Cada largo de la pila es (largo de la pila-1) + 1. El algoritmo Es O(n) similar a cuando se llama con un bucle for ya que visita los N elementos una sola vez

```go
func LargoPila[T any](pila Pila[T]) int {
    // caso base si pila esta vacia
    if pila.EstaVacia() {
        return 0
    }
    dato := pila.Desapilar()
    contador := 1 + LargoPila(pila)
    pila.Apilar(dato)
    return contador
}
```