## Consigna
Si en el ejercicio anterior en vez de quererse los 3 elementos más grandes, se quisieran los K elementos más grandes ¿cómo se debería proceder? ¿Cuál terminaría siendo la complejidad del algoritmo?

## Resolucion

**Caso trivial**

Una forma de pensar es: si tengo un heap y voy desencolando K-veces entonces cada vez que saco hago downheap k-veces lo que hace la complejidad K*log(n). Entonces mi objetivo es encontrar una complejidad mejor que esa.

Si solo necesito devolver K máximos, entonces tendría que recorrer el arreglo que seria O(1)


**Pensar complejidad mejor que KLog(n)**
