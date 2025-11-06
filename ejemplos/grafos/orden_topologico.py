import random
from grafos import Grafo
from collections import deque


"""
Orden topológico con BFS:
- Es un BFS normal pero encola sólo grados de entradas 0.

Algoritmo BFS Topológico:
    - Calculo grados de entrada  O(V + E)
    - encolar grados de entrada 0 O(v + e)
    - mientras cola no esta vacia:  O(V)
        - Desencolar un vertice
        - Ver sus adyacentes    O(E)
            - Para cada adyacente restar 1
            - Si es 0 encolar
                - Actualizar orden
Complejidad: O(v + E)
"""

def orden_bfs(grafo: Grafo) -> list:
    """
    orden_bfs toma un grafo y devuelve una lista de orden
    """
    g_entrada = calcular_grados_entrada(grafo)
    cola = deque()
    orden = []
    
    for v in grafo.obtener_vertices():
        if g_entrada[v] == 0:
            cola.append(v)
            orden.append(v)
    
    while len(cola) > 0:
        v = cola.popleft()
        for w, _ in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.append(w)
                orden.append(w)

    return orden
            


def calcular_grados_entrada(grafo: Grafo) -> dict:
    """
    calcular_grados_entrada
    devuelve un diccionario con grados de entrada de cada vertice
    """
    grados = dict()
    for v in grafo.obtener_vertices():
        grados[v] = 0
    for v in grafo.obtener_vertices():
        for w, _ in grafo.adyacentes(v):
            grados[w] = grados[w] + 1
    return grados



"""
Orden Topológico con DFS
- Es un DFS pero eligiendo vértices random

Elegir vertices random hasta que todos los vertices estén visitados:
- Para cada vértice (V):
    - Si no está visitado::::
    - Ver adyacentes -> para cada adyacente (W)
        - Si no está visitado
            - Marcar como visitado
            - Llamada recursiva para ese adyacente
    - Apilar el vértice (V)

- Revertir desapilar todo y devolver orden
"""

def orden_dfs(grafo: Grafo) -> list:
    visitado = set()
    pila = deque()
    orden = []
    
    for v in grafo.obtener_vertices():
        if v not in visitado:
            visitado.add(v)
            _visitar_dfs(v, grafo, visitado, pila)

    while len(pila) > 0:
        orden.append(pila.pop())

    return orden

def _visitar_dfs(v, grafo: Grafo, visitado: set, pila: deque):
    for w, _ in grafo.adyacentes(v):
        if w not in visitado:
            visitado.add(w)
            _visitar_dfs(w, grafo, visitado, pila)
    pila.append(v)    

"""
Funciones Auxiliares XD
"""
    
def crear_grafo_ejemplo() -> Grafo:
    grafo = Grafo(True)
    # Variables (valor a poner) de vertices
    analisis_ii = "Análisis Matemático"
    fundamentos = "Fundamentos de Programacion"
    ids = "Introduccion a desarrollo de Software"
    algebra = "Algebra Lineal"
    orga = "Organizacion del computador"
    algo_ii = "Algoritmos y Estructuras de datos"
    paradigmas = "Paradigmas de programacion"
    proba = "Probabilidad y estadística"
    sisops = "Sistemas Operativos"
    teoria_algo = "Teoría de Algoritmos"
    materias = [analisis_ii, fundamentos, ids, algebra, orga, algo_ii, paradigmas, proba, sisops, teoria_algo]
    # Insertar vertices de grafo
    for v in materias:
        grafo.agregar_vertice(v)
    # Insertar Aristas (relaciones / correlativas)
    grafo.agregar_arista(analisis_ii, proba)
    grafo.agregar_arista(fundamentos, orga)
    grafo.agregar_arista(fundamentos, algo_ii)
    grafo.agregar_arista(ids, paradigmas)
    grafo.agregar_arista(ids, teoria_algo)
    grafo.agregar_arista(algebra, proba)
    grafo.agregar_arista(orga, sisops)
    grafo.agregar_arista(algo_ii, paradigmas)
    grafo.agregar_arista(algo_ii, teoria_algo)
    return grafo

if __name__ == '__main__':
    grafo_materias = crear_grafo_ejemplo()
    print(grafo_materias.obtener_vertices())
    print(orden_bfs(grafo_materias))
    print(orden_dfs(grafo_materias))