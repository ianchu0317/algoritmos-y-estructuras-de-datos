## Consigna

 Indicar si las siguientes afirmaciones son verdaderas o falsas. En caso de ser verdaderas, justificar, en caso de ser falsas poner un contraejemplo:

a. Si dos árboles binarios tienen el mismo recorrido inorder, entonces tienen la misma estructura.

b. Si dos árboles binarios tienen el mismo recorrido preorder, entonces tienen la misma estructura.

c. Si dos árboles binarios de búsqueda (e idéntica función de comparación) tienen el mismo recorrido preorder, entonces tienen la misma estructura.


## Resoluciones

a. Falso. Si dos arboles binarios tienen mismo inorder no necesariamente tienen misma estructura. Por ejemplo si tengo inorder A-B-C entonces sabiendo que inorder es izquierda-nodo-derecha, una solucion puede ser B raiz, A izquierda y C derecha. Pero otra solucion es C raiz, B izquierda de C, A izquierda de B

b. Falso. Si es arbol **binario** tienen mismo preorder no necesariamnete tienen misma estructura. Por ejemplo si tengo preorder A-B-C. Sabiendo que preorder es nodo-izq-der entonces una estructura con el preorder puede ser A raiz, B izquierda de A, C derecha de B. Pero otra solucion puede ser A raiz, B derecha de A, C derecha de B.

c. Verdadero. Si es ABB entonces sí tiene misma estructura. Es como ir recorriendo por el orden que se fue insertando. Si tengo ABB y su orden pre-order puedo reconstruir el arbol original, ya que es unica.

_Correccion C: sabemos que pre-order mantiene primer elemento como raiz, luego los elementos restantes son del subarbol izquierdo y/o subarbol derecho. Por la propiedad de ABB izquierda<nodo<derecha, cada elemento luego de raiz, se va reubicando automaticamente en su posicion. De esta manera se reconstruye el mismo arbol_