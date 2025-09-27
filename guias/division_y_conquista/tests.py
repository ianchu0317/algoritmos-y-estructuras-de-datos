
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
print(lista)
print("Esta ordenado: ", estaOrdenado(lista, len(lista)))
print(lista2)
print("Esta ordenado: ", estaOrdenado(lista2, len(lista2)))


# Ejercicio 12 hallar raiz cuadrada (parte entera)
def raizCuadrada(n, inicio, final):
    mitad = (inicio + final) // 2
    if mitad**2 <= n and (mitad+1)**2 > n:
        return mitad
    else:
        if mitad**2 < n:
            return raizCuadrada(n, mitad+1, final)
        else:
            return raizCuadrada(n, inicio, mitad-1)

print("Raiz cuadrada de ", raizCuadrada(25, 1, 25))


# Ejercicio 13 hallar posicion de elemento pico
def posicionPico(arr, inicio, final):
    mitad = (inicio + final) // 2
    print(arr[mitad])
    # Caso base
    if arr[mitad-1] < arr[mitad] > arr[mitad+1]:
        return mitad
    else:
        if arr[mitad-1] > arr[mitad]:
            return posicionPico(arr, inicio, mitad-1)
        elif arr[mitad+1] > arr[mitad]:
            return posicionPico(arr, mitad+1, inicio)

lista_pico = [1, 2, 3, 1, 0, -2, -1, -50]
print("Posicion pico es: ", posicionPico(lista_pico, 0, len(lista_pico)-1))


