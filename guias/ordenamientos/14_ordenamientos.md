## Consigna

Implementar en Go un algoritmo de RadixSort para ordenar un arreglo de Alumnos (estructuras) en función de sus notas en parcialitos, de menor a mayor. Los alumnos tienen su nombre y las notas (numéricas, 0-10) de los 3 parcialitos (tenemos las notas finales). El arreglo debe quedar ordenado primero por el promedio de notas. No importan los decimales, nada más si tuvo “promedio 9”, “promedio 8”, etc., es decir la parte entera del promedio. Luego, en caso de igualdad en este criterio, los alumnos deben quedar ordenados por la nota del parcialito 1, en caso de persistir la igualdad, la del parcialito 2, y finalmente por la del 3. En ningún caso se utiliza el nombre para desempatar. Suponer que cualquier algoritmo de ordenamiento auxiliar que se requiera ya se encuentra implementado. Sí justificar por qué se utiliza el o los algoritmos auxiliares utilizados, y no otros. Indicar y justificar la complejidad del algoritmo. El desarrollo de la complejidad debe estar completo, no se aceptan resultados parciales. Hacer un seguimiento sobre el siguiente ejemplo. No es necesario hacer el seguimiento de cómo funciona el o los algoritmos auxiliares.

```
Juan Pérez, 7, 4, 5              Hector Fernández, 8, 9, 10
María Gómez, 5, 7, 4             Martina Giménez, 7, 4, 4
Germán González, 4, 8, 5         Mirtha Legrand, 10, 8, 10
Elián Valenzuela, 4, 2, 0        Ricardo Bonafide, 4, 6, 8
Bobo Weghorst, 0, 0, 0 
```

## Resolución

### Planteo
Si queremos ordenar por promedio -> parcial 1 -> parcial 2 -> parcial 3, entonces en nuestro radixSort la función tiene que ordenar primero parcial 3 -> parcial 2 -> parcial 1 -> promedios ya que ordenamos desde orden menos significativo al más significativo. 

### Ordenamiento auxiliar
El ordenamiento auxiliar a utilizar es el Counting Sort ya que las notas van de 0 a 10, es decir un rango acotado de 11 elementos en total. Esto se aplica tanto para las notas de parciales como para promedios. Además el algoritmo es estable, lo que es una condición para radix sort


### Implementación

```go
func OrdenarAlumnos(alumnos []Alumno) []Alumno {
    alumnos = ordenarParcial3(alumnos)
    alumnos = ordenarParcial2(alumnos)
    alumnos = ordenarParical1(alumnos)
    alumnos = ordenarPromedios(alumnos)
    return alumnos
}
```

### Complejidad
En cada Ordenamiento de notas la complejidad es de O(n + 11) ya que 11 es la cantidad diferentes que puede tomar la nota y por ende el tamaño de los arreglos auxiliares utilizando Counting Sort.

Entonces si aplicamos 4 veces Counting sort serían `4 * O(n + 11) = O(n + 11) = O(n)`, por lo que el algoritmo para ordenar alumnos con el criterio dado tiene una complejidad de O(n).

### Seguimiento

Ordenamiento por parcial 3 (ultima posicion)
```
Elián Valenzuela, 4, 2, 0
Bobo Weghorst, 0, 0, 0
María Gómez, 5, 7, 4
Martina Giménez, 7, 4, 4
Juan Pérez, 7, 4, 5
Germán González, 4, 8, 5
Ricardo Bonafide, 4, 6, 8
Hector Fernández, 8, 9, 10              
Mirtha Legrand, 10, 8, 10
```

Ordenamiento por parcial 2 (segunda posicion)
```
Bobo Weghorst, 0, 0, 0
Elián Valenzuela, 4, 2, 0
Martina Giménez, 7, 4, 4
Juan Pérez, 7, 4, 5
Ricardo Bonafide, 4, 6, 8
María Gómez, 5, 7, 4
Germán González, 4, 8, 5
Mirtha Legrand, 10, 8, 10
Hector Fernández, 8, 9, 10
```

Ordenamiento por parcial 1 (primera posicion)
```
Bobo Weghorst, 0, 0, 0
Elián Valenzuela, 4, 2, 0
Ricardo Bonafide, 4, 6, 8
Germán González, 4, 8, 5
María Gómez, 5, 7, 4
Martina Giménez, 7, 4, 4
Juan Pérez, 7, 4, 5
Hector Fernández, 8, 9, 10
Mirtha Legrand, 10, 8, 10
```

Ordenamiento por promedios
```
Bobo Weghorst, 0, 0, 0
Elián Valenzuela, 4, 2, 0
Germán González, 4, 8, 5
María Gómez, 5, 7, 4
Martina Giménez, 7, 4, 4
Juan Pérez, 7, 4, 5
Ricardo Bonafide, 4, 6, 8
Hector Fernández, 8, 9, 10
Mirtha Legrand, 10, 8, 10
```