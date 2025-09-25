## Consigna
Suponer que se tiene un arreglo de n elementos ordenados, seguido de f(n) elementos desordenados. Cómo ordenarías el arreglo según si f(n) es

a. f(n)=O(1)
b. f(n)=O(log n)
c. f(n)=O(sqrt(n))

---

## Interpretación de la consigna:

Tienes un array que está **parcialmente ordenado**:

```
[elementos ordenados] + [elementos desordenados]
      n elementos           f(n) elementos
```

### Ejemplos concretos:

**Caso a) f(n) = O(1)** (constante):
```
Array total: [1,2,3,4,5,6,7,8] + [15,9]
            ↑─── n=8 ordenados ───↑  ↑─f(n)=2 desordenados─↑
```

**Caso b) f(n) = O(log n)** :
```
Array total: [1,2,3,4,5,6,7,8] + [15,9,12]
            ↑─── n=8 ordenados ───↑  ↑─f(n)=3≈log(8) desordenados─↑
```

**Caso c) f(n) = O(√n)**:
```
Array total: [1,2,3,4,5,6,7,8] + [15,9,12,20]
            ↑─── n=8 ordenados ───↑  ↑─f(n)=4≈√8 desordenados─↑
```

**La pregunta es:** ¿Cuál es la **estrategia más eficiente** para ordenar todo el array en cada caso?

**Estrategias posibles:**
1. **Insertion Sort** en los f(n) elementos
2. **Merge** de las dos partes  
3. **Ordenar todo desde el principio**


---

a. Para este primer caso, como los elementos desordenados es constante, entonces podríamos aplicar Inserción para el arreglo, ya que si n tiende un valor grande y los valores a cambiar son constantes podríamos decir que son casi despreciables, entonces el algoritmo tendería O(n) ya que recorre desde el primer elemento hasta el último, salvo en las pocas excepciones donde tendría que insertar un elemento.

b. Para el caso de f(n)=O(log(n)) elementos desordenados, si ordenamos el lado de f(n) y hacemos merge, sabemos que ordenar con mergeSort f(n) sería O(log(n) * log(n)) ya que en vez de n elementos tenemos log(n) elementos y es O(log(2n))=O(log(n)) y luego una vez ordenado f(n) lo mergeamos con la parte ordenada un algoritmo de O(n) y finalmente entonces nos queda O(n)+O(log(n)) = O(log(n)) nos queda esto

c. Si aplico un algoritmo de n log n sería rapido

---

## Correccion

b) Para f(n) = O(log n):

Estrategia: Ordenar f(n) + Merge
- Ordenar f(n) elementos: O(log n × log(log n))
- Merge con n elementos: O(n + log n) = O(n)
- Total: O(n + log n × log(log n)) = O(n)

Alternativa: Insertion Sort directo
- Cada elemento de f(n) se inserta en O(n) peor caso
- Total: O(log n × n) = O(n log n)

¡Mejor la primera estrategia!


c) Para f(n) = O(√n):

Estrategia: Ordenar f(n) + Merge
- Ordenar f(n) elementos: O(√n × log √n) = O(√n × log n)
- Merge: O(n + √n) = O(n)
- Total: O(n + √n log n)

Comparar con ordenar todo desde cero: O((n + √n) log(n + √n)) = O(n log n)

Como √n log n < n log n para n grande, ¡conviene la estrategia híbrida!