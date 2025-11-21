from grafo import Grafo
from collections import deque
import heapq



def crear_grafo_testing() -> Grafo:
    vertices = [1, 2, 3, 4, 5, 6, 76, 87]
    grafo = Grafo(False, vertices=vertices)
    
    grafo.agregar_arista(1, 2)
    grafo.agregar_arista(3, 1, 3)
    grafo.hay_vertice
    return grafo
    
    
    
if __name__ == '__main__':
    grafo = crear_grafo_testing()
    print(grafo.peso_arista(1, 3))