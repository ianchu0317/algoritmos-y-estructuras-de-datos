## Consigna

Dada una pila de enteros, escribir una función que determine si sus elementos están ordenados de manera ascendente. Una pila de enteros está ordenada de manera ascendente si, en el sentido que va desde el tope de la pila hacia el resto de elementos, cada elemento es menor al elemento que le sigue. La pila debe quedar en el mismo estado que al invocarse la función. Indicar y justificar el orden del algoritmo propuesto.

## Resolucion

Lo que podria hacer es desapilar y comparar con el elemento anterior (que sería el tope) y luego volver a apilarlo. Realizar esto sucesivamente hasta que la pila esté vacía, es decir es un caso de recursión donde el caso base es cuando la pila está vacia. Si la pila está vacía entonces es ascendente.

```go
func esAscendente(pila Pila[int]) bool {
    ascendente := true

    if pila.EstaVacia() {
        return true
    }

    elementoActual := pila.Desapilar()
    elementoAnterior := pila.VerTope()
    
    if elementoActual < elementoAnterior {
        return false
    }
    
    ascendente = ascendente && esAscendente(pila)
    pila.Apilar(elementoActual)
    return ascendente
}
```

Esto presenta problemas:
- Si hay un solo elemento
- Si orden ascendente

Rehago el ejercicio

```go
func esAscendente(pila Pila[int]) bool {
    var resultado bool

    // verificar si no hay elementos en pila
    if pila.EstaVacia() {
        return true
    }
    elementoActual := pila.Desapilar()

    // verificar si pila hay un solo elemento
    if pila.EstaVacia() {
        return true
    }
    elementoAnterior := pila.VerTope()
    
    // comparar los dos elementos
    if elementoActual > elementoAnterior {
        resultado = false
    } else {
        resultado = esAscendente(pila)
    }

    pila.Apilar(elementoActual)
    return resultado
}
```

Y en este caso lo que hago es como siempre ir adelante con todos los elementos así que es O(n). Otra opción podría utilizar una pila auxiliar

```go
func esAscendente(pila Pila[int]) bool {
    pilaAux := Pila[int]{}
    var resultado bool
    var elementoActual, elementoAnterior int

    // hacer comparacion
    for !pila.EstaVacia() && resultado {
        elementoActual = pila.Desapilar()
        elementoAnterior = pila.VerTope()
        resultado = elementoActual > elementoAnterior
        pilaAux.Apilar(elementoActual)
    }

    // volver a apilar
    for !pilaAux.EstaVacia(){
        pila.Apilar(pilaAux.Desapilar())
    }

    return resultado
}

```