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
    
    cant_izq = contar_apariciones(arr, elem_izq)
    cant_der = contar_apariciones(arr, elem_der)
    
    if not hay_izq and not hay_der:
        return False, -1
    
    if elem_izq == elem_der:
        pass
    
    return True, elem_izq

def contar_apariciones(arr, num):
    contador = 0
    for x in arr:
        if x == num:
            contador += 1
    return contador

def hay_mayoria(arr):
    res, _, _ = _hay_mayoria(arr)
    return res

if __name__ == '__main__':
    arr1 = [1, 2, 3, 1, 1, 1] # -> True
    print(hay_mayoria(arr1))
    