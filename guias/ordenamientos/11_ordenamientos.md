
## Consigna

Implementar un algoritmo que, dado un arreglo de n números enteros cuyos valores van de 0 a K (constante conocida), procese dichos números en tiempo O(n+K), devuelva alguna estructura que permita consultar cuántos valores ingresados están en el intervalo (A, B), en tiempo O(1). Explicar cómo se usaría dicha estructura para poder realizar tales consultas.

## Resolucion

Okey, sabiendo que nuestro array tiene n elementos y que los valores son acotados de [0, K], podemos utilizar Counting Sort para este problema. Es más, como nos pide que sea O(n+K), coincidiendo con dicho algoritmo.

Counting sort consiste en 3 pasos: 
1. Calcular frecuencia O(n)
2. Calcular array suma de inicios O(K)
3. Insertar los elementos en su posición

Si consigo llegar hasta el segundo paso y devuelvo el array de la suma de los elementos (array de inicios), entonces puedo saber cuántos elementos estuvieron antes de un numero en una sola operación O(1). Si repetimos esto para dos numeros sería 2 * O(1) = O(1). A lo que quiero llegar es que si hacemos 
`cant numeros antes de A - cant numeros antes de B = cant numeros entre A y B`

Y la cantidad de numeros tanto para A y para B pertenecientes a [0, K] lo podemos obtener del array de los inicios o sumas.

Acá está la implementación en Go

```go
// Complejidad O(n+k)
func devolverArraySuma(arr []int, k int) []int {
    // Calcular las frecuencias O(n)
    frecuencias := make([]int, k)
    for _, num := range arr {
        // O(1)
        frecuencias[i]++
    }
    // Calcular las sumas O(k)
    sumas := make([]int, k)
    for i := 1; i < k; i++ {
        sumas[i] = sumas[i-1] + frecuencias[i-1] // O(1)
    }
    return sumas    // O(1)
}

func cantNumEntre(a, b int, sumas []int) int {
    return sumas[a] - ssumas[b]  // O(1)
}
```