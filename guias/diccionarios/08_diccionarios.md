## Consignas
Se tiene un hash que cuenta con una función de hashing que, recibida una clave, devuelve la posición de su inicial en el abecedario. La capacidad inicial del hash es 26. Para los puntos B, C y D indicar y justificar si las afirmaciones son verdaderas o falsas. Se puede considerar que todas las claves serán palabras (sólo se usan letras para las claves).

a. Mostrar cómo quedan un hash abierto y un hash cerrado (sólo el resultado final) tras guardar las siguientes claves: Ambulancia (0), Gato (6), Manzana (12), Ananá (0), Girasol (6), Zapato (25), Zapallo (25), Manzana (12), Bolso (1). Aclaración: No es necesario hacer una tabla de 26 posiciones, lo importante es que quede claro en cuál posición está cada elemento.

b. En un hash abierto con dicha función de hashing, se decide redimensionar cuando la cantidad alcanza la capacidad (factor de carga = 1). El rendimiento de hash_obtener() es mejor en este caso respecto a si se redimensionara al alcanzar un factor de carga 2.

c. En un hash cerrado con dicha función de hashing, si se insertan n + 1 claves diferentes (considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra distinta, en el primer caso Obtener() es 
O(n) y en el segundo siempre O(1).

d. En un hash abierto con dicha función de hashing, si se insertan n + 1 claves diferentes (considerar que se haya redimensionado acordemente), n con la misma letra inicial, y 1 con otra distinta, en el primer caso Obtener() es 
O(n) y en el segundo siempre O(1).

## Resolucion

### Punto a.
**Hash abierto**

En un abierto se implementan listas de celdas por lo que son de mismo indice luego de hashear clave simplemente se insertan al final de la lista.

```
0: Ambulancia -> Anana
1: Bolso
6: Gato -> Girasol
12: manzana
25: Zapato - Zapallo
```

**Hash cerrado**

En un hash cerrado se implementan en una tabla (arreglo) de celdas. Los elementos con mismo indice luego de aplicar funcion de hasheo simplemente se van hasta proxima posicion no vacia. En caso de estar al final y esta repetido, volver al inicio.


Ambulancia (0), Gato (6), Manzana (12), Ananá (0), Girasol (6), Zapato (25), Zapallo (25), Manzana (12), Bolso (1)
```
0: Ambulancia
1: Anana
2: Zapallo
3: Bolso
6: Gato
7: Girasol
12: Manzana
25: Zapato
```

### Punto b.

Falso. El factor de carga se calcula como `Cantidad almacenada / tamaño de tabla (arreglo)`. Si los elementos guardados son todos los elementos de la misma clave menos uno, y uno distinto, entonces tendria una lista con (n-1) elementos y otra lista con 1 elemento. Entonces la primitiva Hash obtener Tardaria mucho en buscar la celda si tiene la misma clave que todos los demas (n-1). La redimension en este caso, teniendo en cuenta el metodo de hasheo, no afecta, ya que si anteriormente todos los elementos de esa lista (n-1) comienzan con esa letra, luego de redimensionar no cambia el hecho de que siguen comenzando con esa letra y continuan estando en la misma lista, así que no cambia mucho cuando o no se hace la redimension.



### Punto c. 
Falso. No necesariamente ocurre esto. En un hash cerrado si tengo n con misma letra inicial, quiere decir que al hashear obtengo la misma celda, así que para funcion de buscar necesito ir a la proxima celda vacia, osea complejidad O(n). Pero puede sucederr que se sobreescriba la misma celda en la que iria ese único elemento con clave distinto, y tendría que recorrer O(n) para encontrar su posicion. Por ejemplo, si tengo 26 posiciones disponibles, inserto 25 de la misma letra "A", y luego quiero insertar en su letra "B", entonces "B"tendría que recorrer 24 posiciones para hallar su posicion libre. 
Esta afirmacion cumple cuando se inserta antes de que se sobreescriba la celda de ese único elemento.

_Correccion: Depende si por colisiones se sobreescribe la celda de ese unico elemento a insertar. En el peor de los casos la busqueda para los elementos que inician con misma letra es O(n) y para la busqueda de esa inicial distinta tambien es O(n)_


### Punto d.
Verdadero. Se cumple esta afirmacion ya que al hashear y obtener todos pertenecientes a una misma lista, la lista va a tneer tamaño n, ya que tengo n-elementos con misma clave; entonces obtener un elemento con esa misma letra inicial, quiere decir iterar los n-elemtnos de esa lista hasta encontrar la celda que busco. Si tengo una clave distinta y con otra letra inicial, entonces se guardará en otra posicion del arreglo y sería una lista de un elemento, entonces iterar la lista para obtener el valor equivale a obtener el único de elemtno de esa lista (primer y ultimo elemento), una operacion de O(1).