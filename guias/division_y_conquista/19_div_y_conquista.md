## Consigna

 Se sabe, por el teorema de Bolzano, que si una función es continua en un intervalo [a, b], y que en el punto a es positiva y en el punto b es negativa (o viceversa), necesariamente debe haber (al menos) una raíz en dicho intervalo. Implementar una función func raiz(f func(int)int, a int, b int) int que reciba una función (univariable) y los extremos mencionados y devuelva una raíz dentro de dicho intervalo (si hay más de una, simplemente quedarse con una). La complejidad de dicha función debe ser logarítmica del largo del intervalo [a, b]. Asumir que por más que se esté trabajando con números enteros, hay raíz en dichos valores: Se puede trabajar con floats, y el algoritmo será equivalente, simplemente se plantea con ints para no generar confusiones con la complejidad. Justificar la complejidad de la función implementada.

 ## Resolucion

 ### Puntos clave
 - solución logarítmica -> una búsqueda
 - intervalo cerrado -> parecido a lineal

 ### Planteo
 Lo que estamos buscando es un numero en un rango -> podemos pensarlo como una busqueda binaria dentro de un rango de numeros.

 El caso base o el elemento que estamos buscando es f(x)=0, siendo X el elemento que buscamos. 


Hacemos como busqueda binaria: vamos por el elemento del medio y decidimos si partir hacia derecha o izquierda, y así sucesivamente hasta encontrar el elemento. -> la condición para ver si es derecha o izquierda dependerá de si en el tramo a la función era positiva o no.

acá estaría la implementacion

```go
// suponiendo tramo a positivo y tramo b negativo
func raiz(f func(int) int, a, b int) int {
    mitad := (a+b) / 2
    if f(mitad) == 0 {
        return mitad
    } else {
        if f(mitad) > 0 {
            return raiz(f, mitad + 1, b)
        } else {
            return raiz(f, a, mitad - 1)
        }
    }
}
```
La complejidad de este ejercicio es de O(log(n)) ya que su ecuacion de recursividad es T(n)=T(n/2)+O(1)