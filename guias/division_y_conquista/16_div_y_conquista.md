## Consigna
Implementar una función (que utilice división y conquista) de orden 
O(nlogn) que dado un arreglo de n números enteros devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces. Justificar el orden de la solución. Ejemplos:

```
[1, 2, 1, 2, 3] -> false
[1, 1, 2, 3] -> false
[1, 2, 3, 1, 1, 1] -> true
[1] -> true
```

## Resolucion

Como dato nos dice que tiene que ser O(n log(n)) por lo que la solucion debe ser de `T(n)=kT(n/k) + O(n)`, pero como para simplificar dejamos A=B=2 -> de esta parte vemos que es como el mismo problema para un subarray, osea que los subarray tienen el mismo problema que el array principal. Acerca de O(n) podria intentar bueno recorrer el array poniendolo en un array auxiliar y contar las frecuencias pero el ejercicio me pide unicamente de division y conquista.

Podriamos comenzar analizando el caso mas simple `[1]->true`

Esto es lo que pense que deberia hacer la funcion:
- Función hayCantidad(arreglo [])
    - Chequear largo del arreglo y calcular mitad
    - Recorrer el array si hay elementos más de la mitad
    - Partir mitad en subarreglos
    - Caso base si len(arr) == 1 entonces devolver true. 
    