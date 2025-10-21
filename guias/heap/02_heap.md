## Consigna

Implementar una primitiva para el heap (de máximos) que obtenga los 3 elementos más grandes del heap en O(1).

## Resolucion

En este caso es trivial. Si estamos realizando una primitiva de un HEAP, tenemos acceso a su estructura interna que en este caso es un arreglo. Se sabe que el elemento superior es el máximo de todo el heap, y sus hijos izquierdo y derecho son los máximos de subarbol izquierda y derecha, osea los tres primeros elementos más grandes del heap es devolver arr[0], arr[0*2+1] = arr[1], arr[0*2+2] = arr[2]

---

## Correcccion
Es mucho mejor utilizar variables auxiliares en la estructura donde mantenga top 3 actualizado en insertar y borrar. Aumenta en insercion y borrar ya que tengo que actualizar pero la complejidad en devolver es O(1)

