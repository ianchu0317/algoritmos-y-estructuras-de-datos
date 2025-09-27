## consigna
¿Puede aplicarse counting sort a un arreglo de floats cuyos valores se encuentran entre 0 y 1? En caso de poderse, explicar cómo lo harías. En caso de no poderse, explicar por qué y mencionar un algoritmo que resuelva el problema, en un orden semejante al de counting sort. ¿Necesitarías más información para aplicar dicho algoritmo, o con saber que el rango es de 0 a 1 es suficiente?

## Resolucion

No se puede directamente ordenar un arreglo de floats con Counting Sort ya que no es un valor Entero. El counting Sort para que funcione ponemos el número como índice de un arreglo auxiliar tanto para contar la frecuencia como para la sumatoria y calcular inicios, pero si el valor es flotante no podemos realizar directamente sin establecer otras condiciones.

Una opción alternativa al counting sort es Bucket sort, ya que podemos dividir los numeros flotantes en distintos buckets: los números entre 0 y 0.1 y 0.1 y 0.2 así, es decir 10 buckets en total. Luego en cada bucket aplicamos algún ordenmiento auxiliar por ejemplo merge sort. Este algoritmo no está restringido por ser números flotantes y tiene una complejidad similar al Counting sort: O(n * log(n/10)), pero como es log(n/b) crece lento y queda despreciable quedandonos O(n) similar a Counting sort O(n+k). 

Adicionalmente estaría tener en cuenta información como distribución de datos para aplicar las condiciones de la separación de cada bucket para que queden casi del mismo tamaño.