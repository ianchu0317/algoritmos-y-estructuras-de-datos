## Consigna

 Indicar Verdadero o Falso, justificando de forma concisa en cualquier caso.

a. Podríamos mejorar el orden de complejidad de QuickSort si contaramos con más información sobre cómo son los datos a ordenar.

b. No siempre conviene utilizar Counting Sort para ordenar un arreglo de números enteros.

c. Que un algoritmo de ordenamientos sea estable implica que el algoritmo ordena sobre el arreglo original (sin utilizar otro adicional). Por ejemplo, Counting Sort no es estable.

## Resolucion

a) Verdadero. La complejidad de Quick Sort no se podría mejorar ya que igualmente por ejemplo con un conjunto de datos ya ordenados, tendría que llamar la recursividad. Y la complejidad es O(n log n)

b) Verdadero. El Counting Sort sirve para ordenar un arreglo de números enteros de un rango acotado. Si el enunciado no menciona esa información adicional no se puede aplicar Counting Sort.

c) Falseo. Un algoritmo de ordenamientos estable quiere decir que el ordenamiento no afecta el orden de aparición de los elementos iguales luego de ordenar. Que el algoritmo sea in-place quiere decir que se ordena sobre el arreglo original. El Counting Sort sí es estable, pero NO es in-place.

---

## Corrección
a) VERDADERO. Con más información podríamos:
- Elegir mejor pivot (mediana de 3, mediana de medianas)
- Cambiar a algoritmo lineal si conocemos el rango (Counting Sort)
- Usar híbridos: Quick Sort + Insertion Sort para arrays pequeños

b) VERDADERO. Counting Sort NO conviene cuando:
- Rango muy grande (ej: números de 0 a 1,000,000 con pocos elementos)
- Espacio O(k) sería mayor que O(n log n) de Quick/Merge Sort
- Números no enteros (floats, strings)