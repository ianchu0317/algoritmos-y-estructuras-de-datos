## Consigna

Hacer el seguimiento de counting sort para ordenar por año las siguientes obras:

    1988 - Crónicas del Ángel Gris
    2000 - Los Días del Venado
    1995 - Alta Fidelidad
    1987 - Tokio Blues
    2005 - En Picada
    1995 - Crónica del Pájaro que Da Cuerda al Mundo
    1995 - Ensayo Sobre la Ceguera
    2005 - Los Hombres que No Amaban a las Mujeres

¿Cuál es el orden del algoritmo? ¿Qué sucede con el orden de los elementos de un mismo año, respecto al orden inicial, luego de finalizado el algoritmo? Justificar brevemente.

## Resolucion
1. **Seguimiento Counting Sort**
Como tenemos que aplicar counting sort a los años, entonces por esta primera parte solo nos interesa la parte numérica.

Armamos primero el array con los años: 
```
[1988, 2000, 1995, 1987, 2005, 1995, 1995, 2005]
```

Luego lo que se hace es crear el array de frecuencias. Como el número máximo es de 2005, podemos crear un array de largo mayor a 2025 elementos y contamos las frecuencias de cada elemento:
```
arrayFrecuencia[1987] = 1
arrayFrecuencia[1988] = 1
arrayFrecuencia[1995] = 3
arrayFrecuencia[2000] = 1
arrayFrecuencia[2005] = 2
```
El resto de los elementos del arreglo sin mencionar está en 0.

Ahora contamos las sumatorias para ver su indice de inicio en el arreglo ordenado. Creamos un arrayInicios
```
inicios[0] = 0
...
inicios[1987] = 0
inicios[1988] = 1
inicios[1995] = 2
inicios[2000] = 5
inicios[2005] = 6
...
```

Por último recorremos nuevamente cada año del arreglo original de años y vamos viendo en donde comienza, insertandolos en su posicion. Creamos nuevoArreglo con tamaño del arreglo original (no tamaño de 0 a 2025)
```
nuevoArreglo[inicios[elemento]] = elemento
inicios[elemento]++

nuevoArreglo = [1987, 1988, 1995, 1995, 1995, 2000, 2005, 2005]
```

2. **Orden y complejidad del algoritmo**
El orden de este algoritmo es O(n + 2025), ya que para el paso 1 contamos las frecuencias de cada arreglo O(1) recorriendo cada elemento del arreglo o(n), y recorremos los arreglos auxiliares para los inicios O(K=2025), y finalmente O(n) para poner los elementos del arreglo original en su lugar ordenado.

3. **Orden de aparición de los elementos de mismo año**
Como counting sort es un algoritmo estable, los elementos ordenados con el mismo año permanecerán con la misma prioridad de aparición que en el arreglo sin ordenar: 

```
    1987 - Tokio Blues
    1988 - Crónicas del Ángel Gris
    1995 - Alta Fidelidad
    1995 - Crónica del Pájaro que Da Cuerda al Mundo
    1995 - Ensayo Sobre la Ceguera
    2000 - Los Días del Venado
    2005 - En Picada
    2005 - Los Hombres que No Amaban a las Mujeres
```

Esto lo vemos en el último paso donde recorremos el arreglo original poneindo sus elementos en su indice e incrementando el contador de inicio para poner el que sigue en la posición de atrás.
