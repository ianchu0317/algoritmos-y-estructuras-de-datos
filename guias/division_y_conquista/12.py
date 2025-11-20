
"""
Se tiene un arreglo de N>=3 elementos en forma de pico, esto es: estrictamente creciente hasta una determinada posición 
p, y estrictamente decreciente a partir de ella (con 0<p<N−1). Por ejemplo, en el arreglo [1, 2, 3, 1, 0, -2] la posición del pico es p=2. Se pide:
Implementar un algoritmo de división y conquista de orden 
O(logn) que encuentre la posición 
p del pico: func PosicionPico(v []int, ini, fin int) int. La función será invocada inicialmente como: PosicionPico(v, 0, len(v)-1), y tiene como pre-condición que el arreglo tenga forma de pico.
Justificar el orden del algoritmo mediante el teorema maestro.
"""

def posicion_pico(arr, ini, fin):
    medio = (ini + fin) // 2
    if arr[medio - 1] < arr[medio] > arr[medio + 1]:
        return medio
    else:
        if arr[medio] < arr[medio+1]:
            return posicion_pico(arr, medio, fin)
        else:
            return posicion_pico(arr, ini, medio) 
        
        
if __name__ == '__main__':
    arr1 = [1, 2, 3, 5, 7, 9, 11, 1, 0, -2] # -> 6
    arr2 = [1, 2, 3, 1, 0, -2] # -> 2
    print(posicion_pico(arr1, 0, len(arr1) - 1))
    print(posicion_pico(arr2, 0, len(arr2) - 1))