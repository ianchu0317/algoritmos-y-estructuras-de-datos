## Consigna

Implementar una función que ordene de manera ascendente una
pila de enteros sin conocer su estructura interna y utilizando como estructura auxiliar sólo otra pila auxiliar. Por ejemplo, la pila `[ 4, 1, 5, 2, 3 ]` debe quedar como `[ 1, 2, 3, 4, 5 ]` (siendo el último elemento el tope de la pila, en ambos casos). Indicar y justificar el orden de la función.

## Resolucion
El algoritmo que se me ocurrio es parecido a una insercion y por ende O(n²) de hecho no hay mucho para mejorar ya que la restriccion es de utilización de una pila auxiliar nada mas. 

Los tres pasos principales que pense fue: 
1. Desapilar la pila original
2. Los elementos desapilados ir en su posicion en pila auxiliar (en inversa)
3. Los elementos de auxiliar pasar al original

Medio que el punto 1 y 2 son casi iguales. 

Comienzo sacando un elemento y comparo con el tope de pila auxiliar si no está vacia: si es mayor al tope entonces saco el tope y lo vuelvo a apilar al pila original (repetir esto hasta que haya encontrado su posición). Una vez encontrado lo apilo en aux en su lugar y vuelvo a desapilar la pila original y apilarlo al aux la cantidad que saqué (bucle for y llevar contador). Repetir este proceso para todos los elementos de la pila original. 

Escrito en código sería así

```go
func OrdenarPila(pila Pila[int]) {
    pilaAuxiliar := CrearPilaAuxiliar[int]{}
    
    for !pila.EstaVacia() {
        // Desapilar el elemento
        elementoActual := pila.Desapilar()
        contador := 0
        // Encontrar su posición
        for !pilaAuxiliar.EstaVacia() && pilaAuxiliar.VerTope() < elementoActual {
            elementoAuxiliar := pilaAuxiliar.Desapilar()
            pila.Apilar(elementoAuxiliar)
            contador += 1
        }
        // Insertar el elemento en su posicion
        pilaAuxiliar.Apilar(elementoActual)
        // Re-apilar al aux los elementos sacados
        for i := range contador {
            elementoAuxiliar := pila.Desapilar()
            pilaAuxiliar.Apilar(elementoAuxiliar)
        }
    }
    // Pasar de auxiliar a pila original
    for !pilaAuxiliar.EstaVacia(){
        elementoAuxiliar := pilaAuxiliar.Desapilar()
        pila.Apilar(elementoAuxiliar)
    }
}
```

