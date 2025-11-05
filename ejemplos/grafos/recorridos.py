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
    grafo = Grafo(False)
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
    """
    - Inicializar estructuras auxiliares
    (Agregar vertice de inicio a estructuras auxiliares)
    
    - Repetir siempre y cuando la cola no este vacia
        - Desencolar el vertice
        - Para cada adyacente del vertice si no esta visitado
            - Encolar a la cola
            - Marcar como visitado -> me aseguro que cada vertice de la cola a visitar no esta visitado
            - Hacer lo que tenga que hacer (padres, distancia, etc)
    """
    cola = deque()
    padres = dict()
    visitados = set()
    
    cola.append(v)
    padres[v] = ""
    
    while len(cola) > 0:
        v = cola.popleft()
        for w, _ in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
                cola.append(w)
    return padres


def dfs(grafo: Grafo, v):
    """
    - Inicializar las estructuras auxiliares
    (Agregar vertice de inicio a estructuras auxiliares (padres, visitados, etc))
    
    - LLamar una funcion auxiliar recursiva para el vertice
        - Para cada vertice adyacente no visitado
            - Marcar como visitado
            - Operaciones
            - Llamada recursiva de ese vértice
    """
    visitados = set()
    padres = dict()
    visitados.add(v)
    padres[v] = ""
    _dfs(v, grafo, visitados, padres)
    return padres

def _dfs(v, grafo: Grafo, visitados: set, padres: dict):
    for w, _ in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            _dfs(w, grafo, visitados, padres)   


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
    print("bfs: ", bfs(grafo1, "A"))
    print("dfs: ", dfs(grafo1, "C"))