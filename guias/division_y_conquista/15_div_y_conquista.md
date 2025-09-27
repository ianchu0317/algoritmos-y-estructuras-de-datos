## Consigna

Un algoritmo sencillo para multiplicar matrices de n×n demora O(n³). El algoritmo de Strassen (que utiliza División y Conquista) lo hace en O(n^log_2(7)). La profesora Manterola quiere implementar un algoritmo de División y Conquista que sea aún más veloz, donde divida al problema en A subproblemas de tamaño de n/4, y que juntar las soluciones parciales sea O(n²). ¿Cuál es el máximo A para que el orden del algoritmo sea menor que el del algoritmo de Strassen?
Justificar.

## Resolucion
Primero armamos la ecuacion de recurrencia hipotetica que se quiere llegar: `T(n)=A*T(n/4) + O(n²)` donde, si aplicamos teorema maestro tenemos: B=4, C=2.

si log_4(A) < 2 entonces -> O(n²)
si log_4(A)= C= 2 entonces -> O(n²log(n)) 
si log_4(A) > C entonces -> O(n^log_4(A)), A > 16

De estos algoritmos el que nos conviene más es que la complejidad sea O(n²) por lo qeu log_4(A) < 2 -> A <16. 

La segunda O(n²log(n)) es valida para n pequeños y para ultima opcion es valido para A>16 y menor a 49 > A > 16. 


El máximo A es 48 ya que aunque 16 estariamos con O(n²), pero con A=48 seguimos mejor que Strassen y es el valor máximo de A para que sea menor que algoritmo de strassen