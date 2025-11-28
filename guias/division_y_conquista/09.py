"""
mplementar una función, que utilice división y conquista, de complejidad 
O(nlogn) que dado un arreglo de 
n números enteros devuelva true o false según si existe algún elemento que aparezca más de la mitad de las veces.
Justificar la complejidad de la solución. Ejemplos:

[1, 2, 1, 2, 3] -> false
[1, 1, 2, 3] -> false
[1, 2, 3, 1, 1, 1] -> true
[1] -> true

Aclaración: Este ejercicio puede resolverse, casi trivialmente, ordenando el arreglo con un algoritmo eficiente, o incluso se puede realizar más rápido utilizando una tabla de hash. Para cumplir con la consigna, resolver sin ordenar el arreglo ni con tabla de hash, sino puramente por división y conquista.
"""

def _hay_mayoria(arr):
    if len(arr) == 1:
        return True, arr[0]

    mitad = len(arr)//2
    izq, der = arr[:mitad], arr[mitad:]
    
    hay_izq, elem_izq = _hay_mayoria(izq)
    hay_der, elem_der = _hay_mayoria(der)
    
    if hay_izq and hay_der and elem_izq == elem_der:
        return True, elem_izq
    
    cant_izq = contar_apariciones(arr, elem_izq) if hay_izq else 0
    cant_der = contar_apariciones(arr, elem_der) if hay_der else 0
    
    if cant_izq > mitad:
        return True, elem_izq
    
    if cant_der > mitad:
        return True, elem_der
    
    return False, -1

def contar_apariciones(arr, num):
    contador = 0
    for x in arr:
        if x == num:
            contador += 1
    return contador

def hay_mayoria(arr):
    res, _ = _hay_mayoria(arr)
    return res

if __name__ == '__main__':
    arr1 = [1, 2, 1, 2, 3] #-> false
    arr2 = [1, 1, 2, 3] #-> false
    arr3 = [1, 2, 3, 1, 1, 1] #-> true
    arr4 = [1] # -> true
    print(hay_mayoria(arr1))
    print(hay_mayoria(arr2))
    print(hay_mayoria(arr3))
    print(hay_mayoria(arr4))
    