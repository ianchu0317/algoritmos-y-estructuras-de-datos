# Ordenar por notación científica
## Consigna
Se tiene una larga lista de números de tres cifras abc que representan números en notación científica de la forma: a,b * 10^c. Por ejemplo 123 representaría el número 1,2 * 10^3.

⚠️ Implemente ⚠️ un algoritmo para ordenar los números según su valor en notación científica. ¿De qué orden es?

Se puede asumir que a nunca será 0 salvo que el número sea efectivamente 0. Es decir, la notación es correcta.

## Resolución
Podemos empezar pensando acerca de la notación `a,b * 10^c`. De esta notación podemos ir viendo las prioridades del número: 
- primero el valor de C -> ya que no es lo mismo `1,2 * 10³` que `1,2 * 10²` 
- El segundo importante es el valor de A -> `1,2*10³` no es lo mismo `2,2*10³`
- Y por último el valor de B es el menos importante ya que es la fraccionaria

Teniendo esto en cuenta vemos que el número abc tiene distinto nivel de importancia dependiendo si es a,b o c, así que podemos aplicar radix sort, comenzando por menos importante, hasta el más importante. 
Como función auxiliar, como cáda dígito sabemos que es en base diez por ende va de 0-9, podemos utilizar counting sort como función auxiliar

Esta sería una implementacion
```go
const RANGO_DIGITOS = 10
func CountingSortN(arr []int, pos int) []int {
    // Contar frecuencias
    frecuencias := make([]int, RANGO_DIGITOS)
    for _, num := range arr {
        digito := (num / pos) % 10
        frecuencias[num]++
    }
    // Contar inicios
    indices := make([]int, RANGO_DIGITOS)
    for i := 1; i < RANGO_DIGITOS; i++ {
        indices[i] = indices[i-1] + frecuencias[i-1]
    }
    // Juntar todo
    // Recorrer array original, 
    // -> buscar sus inicios, agregar ordenado, actualizar
    arrOrdenado := make([]int, len(arr))
    for i, num := range arr {
        digito := (num / pos) % 10
        posicion := indices[digito]
        arrOrdenado[posicion] = num
        indices[digito]++ 
    }
    return arrOrdenado
}

// Ordenar notacion cientifica ABC = a,b * 10^c
func RadixSort(arr []int) []int {
    // Ordenar por B -> digito 
    arr1 := CountingSortN(arr, 10)
    // Ordenar por a
    arr2 := CountingSortN(arr1, 100)
    // Ordenar por c
    arr3 := CountingSortN(arr2, 1)
    return arr3
}
```