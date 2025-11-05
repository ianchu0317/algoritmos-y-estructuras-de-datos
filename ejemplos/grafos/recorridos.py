from grafos import Grafo
from collections import deque

def crear_grafo() -> Grafo:
    """
    Crea un grafo
    A -> D
    B -> D
    C -> D
    A -> B
    A -> C
    """
    grafo = Grafo(True)
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_vertice("D")
    
    grafo.agregar_arista("A", "B")
    grafo.agregar_arista("A", "C")
    grafo.agregar_arista("A", "D")
    grafo.agregar_arista("B", "D")
    grafo.agregar_arista("C", "D")
    
    return grafo


# BFS
def bfs(grafo: Grafo, v):
    cola = deque()
    padres = dict()
    visitados = set()
    
    cola.append(v)
    padres[v] = ""
    
    while len(cola) > 0:
        v = cola.popleft()
        print("Ahora con v: ", v)
        for w, _ in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
                cola.append(w)
                print("w adyacente:", w)
        print()

    return padres


if __name__ == '__main__':
    grafo1 = crear_grafo()
    """
    print(grafo1.obtener_vertices())
    print(grafo1.adyacentes("A")) # A -> B, C, D
    print(grafo1.adyacentes("B")) # B -> D
    grafo1.borrar_arista("A", "B")
    print(grafo1.adyacentes("A")) 
    print(grafo1.adyacentes("B"))
    """
    print(bfs(grafo1, "A"))