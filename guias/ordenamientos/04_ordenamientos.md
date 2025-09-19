## Consigna
Se tiene un arreglo de estructuras de la forma `type Evento struct {anio long, evento string}`, que indica el año y evento de un hecho definido a lo largo de la historia de la Tierra. Indicar y justificar cuál sería un algoritmo de ordenamiento apropiado para utilizar para ordenar dicho arreglo por año. Indicar también, si en vez de ordenar por año se decide ordenar por evento (lexicográficamente). Si se quiere ordenar por año y dentro de cada año, por evento: ¿Deben utilizarse para ambos campos el mismo algoritmo de ordenamiento? ¿Que característica/s deben cumplir dicho o dichos algoritmos para que quede ordenado como se desea? ¿En qué orden deben aplicarse los ordenamientos?
```go
type Evento struct {
    anio long
    evento string
}
```

## Resolucion
1. **Ordenamiento por cada año a lo largo de historia**
Sabemos que hasta ahora (2025) es un rango acotado. Si tomamos los años 0-2025 tenemos un rango acotado de valores enteros, por lo que se podría ordenar por counting sort para optimizar el ordenamiento.
La complejidad en este caso sería de O(n+2025), que si tenemos n elementos muy grandes, 2025 se hace despreciable y se tiende a una complejidad lineal O(n).

En cambio, dependiendo si queremos añadir los eventos a lo largo de la historia de la Tierra, incluyendo los años antes de Cristo (años negativos), el ordenamiento se tendría que hacer con merge sort o quick sort, ya que counting sort en el caso anterior se utiliza para numeros positivos. O también podríamos utilizar Bucket sort con ordenamiento auxiliar merge o quick: dividimos los eventos cada 100 años, 1000 años, por lo que tendíramos +20 buckets y aplicamos merge sort por ejemplo. En ese caso la complejidad sería O(n * log(n/20)), pero como log(n/20) crece lento tiende a O(n).

2. **Ordenamiento por evento lexico-gráfico (diccionario)**
Si se quiere ordenar por evento (que es string), como es cadena tiene varias letras, podemos ordenar según las letras pero para eso aplicamos Radix Sort para cada letra desde el ultimo hasta el primero ya que el primero es el más significativo. La función para ordenar según la letra, como sabemos que el abecedario es limitado, entonces podriamos utilizar un Counting Sort como  ordenamiento auxiliar. El problema con esta resolución, es que se va a ordenar la cantidad de veces coincidiendo con la cantidad de letras que haya la cadena del evento, lo cual no es eficiente por ejemplo para un seguimiento.
La complejidad con este método es de O(d*(n+k)), que para este caso es O(letras*(n+24)), siendo n la cantidad de eventos a ordenar y 24 las letras de abecedario. Pero Si cambiamos a ascii para realizar el counting sort, el rango cambia pero es despreciable al lado de la cantidad de eventos que hay en la historia, por lo que es O(n)

3. **Ordenar por año, y por evento**
En este ordenamiento lo más importante a priori es el año, y el menos significativo es el evento, así que ordenamos primero por evento y luego por año utilizando Radix Sort. Dentro de Radix Sort, utilizamos un Radix Sort para ordenar por eventos, el cual utiliza un Counting sort para ordenar según las letras, y luego aplicamos Bucket Sort con Merge Sort para ordenar por años.
Algo así quedaría:
```go

// Radix Sort para ordenar por evento O()
func ordenarPorEvento(eventos []Evento) []Evento {
    evnetos = OrdenarPorLetra(eventos)
    // repetir por cada letra
    return eventos
}

// Merge Sort
func ordenarPorAño

func OrdenarEventos(eventos []Evento) {
    eventos = ordenarPorEvento(eventos)     // O(letras * (n+24))
    eventos = ordenarPorAño(eventos)        // O(n * log(n/b))
    return eventos
}
```
Con los valores justos el algoritmo tiende a O(n), que es bueno para grandes conjuntos de datos para este caso.
Para el caso de ordenar por Evento y ordenar por Año, si se detalla que los años son positivos, podrían los dos casos utilizar el mismo algoritmo de counting sort para ordenar.


--- 

Correcciones: 
1. El counting sort puede con años negativos
2. Los eventos se ordenan con merge sort ya que no sabemos el largo de la cadena. Luego aplicamos comparacion de lexicografica estandar
3. Bucket sort no es necesario ya que si tengo rango necesario entonces no hace falta. Además si no son MILLONES de años, no necesito reducir, con un merge sort esta bien

Entonces:
- Ordenar por evnetos: O(n Log N) con merge sort
- Ordenar por años: O(n) con Counting Sort
- Ordenar los dos: O(n log N ) -> gana el mayor
