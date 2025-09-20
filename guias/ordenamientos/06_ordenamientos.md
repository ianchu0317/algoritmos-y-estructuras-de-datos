## Consigna
Realizar el seguimiento de ordenar por Radix Sort el siguiente arreglo de cadenas que representan versiones. Cada versión tiene el formato a.b.c, donde cada valor a, b y c puede tener un valor entre 0 y 99. Considerar que se quiere que quede ordenado primero por la primera componente (a), luego por la segunda (b) y finalmente por la tercera (c). Se puede asumir que a nunca será 0 salvo que el número sea efectivamente 0. Es decir, la notación es correcta. Tener en cuenta que, por ejemplo 1.1.3 < 1.1.20, 2.20.8 < 3.0.0.

```
["4.3.2", "5.1.2", "10.1.4", "2.1.20", "2.2.1", "4.2.3", "2.1.5", "8.1.2", "5.30.1", "10.0.23"]
```

## Resolucion

Como queremos ordenar primero por primera componente (a), luego por segunda (b), y por tercera (c), entonces el orden que tenemos que aplicar en Radix Sort es justamente la inversa c -> b -> a por orden de importancia (del menos significativo al más significativo).

Como función auxiliar podemos utilizar el counting sort ya que sabemos que los números toman valor entre 0 y 99, el cual es un rango acotado de ENTEROS, y además es un algoritmo estable.


**Seguimiento**

Ordenar por C
```
["2.2.1", "5.30.1", "4.3.2", "5.1.2", "8.1.2", "4.2.3", "10.1.4", "2.1.5", "2.1.20", "10.0.23"]
```

Ordenar por B
```
["10.0.23", "5.1.2", "8.1.2", "10.1.4", "2.1.5", "2.1.20", "2.2.1", "4.2.3", "4.3.2", "5.30.1"]
```

Ordenar por A
```
["2.1.5", "2.1.20", "2.2.1", "4.2.3", "4.3.2", "5.1.2", "5.30.1", "8.1.2", "10.0.23", "10.1.4"]
```

El algoritmo toma una complejidad de O(3*(n+100)), ya que son tres campos a ordenar y como utilizamos counting sort son 100 el rango que toma (0-99). 