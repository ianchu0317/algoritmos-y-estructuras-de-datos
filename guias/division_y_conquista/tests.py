
# Online Python - IDE, Editor, Compiler, Interpreter

lista = [1, 2, 3, 4, 5, 6]

lista2 = [4, 5, 6, 1, 2, 3]

def estaOrdenado(arr, largo):
    if largo <= 1:
        return True
    if largo == 2:
        return arr[0]< arr[1]
    
    mitad = largo // 2
    izq_ordenado = estaOrdenado(arr[:mitad], mitad)
    der_ordenado = estaOrdenado(arr[mitad:], largo-mitad)
    conexion = arr[mitad-1] <= arr[mitad]
    
    return izq_ordenado and der_ordenado and conexion
"""    
def elementoDesordenado(arr):
    if len(arr) == 2:
        arr[0] > arr[1]
    mitad = len(arr)//2
    izq_ordenado = estaOrdenado(arr[:mitad], mitad)
    der_ordenado = estaOrdenado(arr[mitad:], len(arr)-mitad)
    
    if not izq_ordenado:
        return elementoDesordenado(arr[:mitad])
    else if not der_ordenado:
        return"""

def raizCuadrada(n, inicio, final):
    mitad = (inicio + final) // 2
    if mitad**2 <= n and (mitad+1)**2 > n:
        return mitad
    else:
        if mitad**2 < n:
            return raizCuadrada(n, mitad+1, final)
        else:
            return raizCuadrada(n, inicio, mitad-1)
    
print(lista)
print("Esta ordenado: ", estaOrdenado(lista, len(lista)))
print(lista2)
print("Esta ordenado: ", estaOrdenado(lista2, len(lista2)))
print("Raiz cuadrada de ", raizCuadrada(25, 1, 25))