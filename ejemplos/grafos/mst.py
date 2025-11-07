from grafos import Grafo
from tdas.heap import Heap
import random



"""
Algoritmo de PRIM

- Crear estructuras auxiliares: Heap, visitados, arbol (nuevo grafo)
- Encolar todos los vertices del grafo actual
- Elegir un vertice random y encolar los aristas -> marcar vertice como visitado

- Mientras la cola (heap) no está vacia:
    - Desencolar un arista
    - SI destino de arista no esta visitado:
        - Agregar arista al arbol
        - Ver adyacentes del v:
            - encolar
        - Marcar V como visitado
- Devolver arbol hecho :)

Complejidad O(e * log(e))
"""

def mst_prim(grafo: Grafo) -> Grafo:
    # Estructuras auxiliares
    arbol = Grafo(False)
    visitados = set()
    heap = Heap()
    
    
    # Encolar Todos los vertices al grafo del arbol O(V) 
    for v in grafo.obtener_vertices():
        arbol.agregar_vertice(v)
    
    v = random.choice(grafo.obtener_vertices())
    visitados.add(v)
    
    # Encolar Todos los aristas al heap E1 * log(E)
    for w, peso in grafo.adyacentes(v):
        heap.encolar((peso, v, w))
    
    # E * log(E)
    while not heap.esta_vacia():
        peso, u, v = heap.desencolar()
        if v not in visitados:
            arbol.agregar_arista(u, v, peso)
            visitados.add(v)
            for w, peso in grafo.adyacentes(v):
                if w not in visitados:
                    heap.encolar((peso, v, w))
                
    return arbol



# Crear Arbol de prueba

def crear_grafo_mst() -> Grafo:
    grafo = Grafo(False)
    
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_vertice("D")
    grafo.agregar_vertice("E")
    grafo.agregar_vertice("F")
    
    grafo.agregar_arista("A", "B", 2)
    grafo.agregar_arista("A", "C", 4)
    grafo.agregar_arista("C", "B", 2)
    grafo.agregar_arista("C", "E", 3)
    grafo.agregar_arista("E", "B", 2)
    grafo.agregar_arista("D", "B", 4)
    grafo.agregar_arista("E", "D", 3)
    grafo.agregar_arista("E", "F", 2)
    grafo.agregar_arista("F", "D", 2)

    return grafo
    

if __name__ == '__main__':
    grafo_mst = crear_grafo_mst()
    
    grafo_mst_prim = mst_prim(grafo_mst)
    print(grafo_mst_prim.adyacentes("F"))