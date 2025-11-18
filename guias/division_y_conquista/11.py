"""
Se tiene un arreglo tal que [1, 1, 1, ..., 0, 0, ...] (es decir, unos seguidos de ceros). Se pide:
una función de orden O(logn) que encuentre el índice del primer 0. Si no hay ningún 0 (solo hay unos), debe devolver -1.
demostrar con el Teorema Maestro que la función es, en efecto, O(logn).
"""
def _primer_indice(arr, ini, fin):
    if ini > fin:
        return -1
    if ini == fin:
        if arr[ini] == 0:
            return ini
        else:
            return -1
    
    mid = (ini + fin) // 2
    if arr[mid] == 0:
        return _primer_indice(arr, ini, mid)
    else:
        return _primer_indice(arr, mid+1, fin)

def primer_indice(arr):
    return _primer_indice(arr, 0, len(arr) - 1)

if __name__ == '__main__':
    arr1 = [1, 1, 0, 0, 0] # →  2
    arr2 = [0, 0, 0, 0, 0] # →  0
    arr3 = [1, 1, 1, 1, 1] # → -1
    arr4 = [1, 1, 1, 1, 1, 1, 1, 0, 0] # -> 7
    
    print(primer_indice(arr1))
    print(primer_indice(arr2))
    print(primer_indice(arr3))
    print(primer_indice(arr4))