## Consigna
Explicar por qué el siguiente siguiente código no es de división y conquista.

ˋˋˋgo
// Algoritmo ¿por D&C? para obtener el máximo de un arreglo
func maximo(arreglo []int) int {
    if len(arreglo) == 1 {
        return arreglo[0]
    }
    max_restante := maximo(arreglo[0:len(arreglo)-1])
    if arreglo[len(arreglo) - 1] > max_restante {
        return arreglo[len(arreglo) - 1]
    } else {
        return max_restante
    }
}
ˋˋˋ

## Resolucion 
Esta función no es de division y conquista ya que recorre linealmente el arreglo buscando el máximo. Es decir, va recortando el ultimo elemento del arreglo para comparar con el resto del máximo. Es cierto que utiliza la recursividad como recurso de iteración, pero aún así va descartando de a un elemento y llamando la recursion con arreglo -1, por lo que recorre n-veces recursión siendo n el tamaño del arreglo a encontrar el maximo. 

Para que sea de división y conquista, al llamar la función recursiva deberia hacerlo con una proporcioin del arreglo.

Los algoritmos de division y conquista tienden a dividir el problema entre sub-problemas y resolverlos para luego combinarlos. En este caso no "divide" el problema para resolverlas en varias instancias de recursion. 

Los algoritmos de division y conquista tienden a tener una complejidad similar: T(n)=2O(n/2) + O(1). En este caso la complejidad del algoritmo es T(n)=T(n-1)+O(1)