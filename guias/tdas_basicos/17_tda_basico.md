## Consigna
Escribir una primitiva para la pila (dinámica) cuya firma es `func (pila pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T]` que devuelva una nueva pila cuyos elementos sean los resultantes de aplicarle la función aplicar a cada elemento de la pila original. Los elementos en la nueva pila deben tener el orden que tenían en la pila original, y la pila original debe quedar en el mismo estado al inicial. Indicar y justificar la complejidad de la primitiva.

Por ejemplo, para la pila de enteros [ 1, 2, 3, 6, 2 ] (tope es el número 2), y la función sumarUno (que devuelve la suma entre el número 1 y el número recibido), la pila resultante debe ser [ 2, 3, 4, 7, 3 ] (el tope es el número 3).

## Resolución
### Puntos clave
- Función primitiva -> acceso al array interno -> Sin necesidad de desapilar y apilar de nuevo -> sin necesidad de estructuras auxiliares 

### Ideas
Lo que se me ocurre es acceder a la estructura dinámica (slice), e ir al último elemento hacia adelante con un contador. Apllicar para cada elemento la función y apilarlos a una pila.

```go
func (pila pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T] {
    nuevaPila := CrearPilaDinamica[T]()

    // Aplicar para cada elemento
    for i := len(pila.datos) - 1; i >= 0; i-- {
        nuevoDato := aplicar(pila.datos[i])
        nuevaPila.Apilar(nuevoDato)
    }

    return &nuevaPila
}
```

Este algoritmo tiene complejidad O(n) siendo N la cantidad de elementos de la pila en el momento que llama la primitiva. 