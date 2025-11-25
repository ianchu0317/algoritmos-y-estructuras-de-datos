"""
Escribir una función en Go que, dado un arreglo de `n` cadenas y un entero positivo `k` cadenas más largas. Se 
espera que el orden del algoritmo sea O(n+klogn). Justificar el orden.
"""

import heapq


def k_cadenas_largas(k: int, cadenas: list) -> list:
    res = []
    
    # Heapify o crear max-heap desde arr O(n)
    heapq.heapify_max(cadenas)
    
    # Sacar k elementos del heap de maximos k log(n)
    i = 0
    while i < k and len(cadenas) > 0:
        print(heapq.heappop_max(cadenas))
        i += 1
        #res.append(heapq.heappop_max(cadenas))
    return res


if __name__ == '__main__':
    cadenas = ["auto", "tiranosaurio", "skibidi", "pelota", "ian", "jijija", "lol"]
    print(k_cadenas_largas(3, cadenas))