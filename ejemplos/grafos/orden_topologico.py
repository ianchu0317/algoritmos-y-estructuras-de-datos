from grafos import Grafo
from collections import deque


"""
Orden topológico con BFS:
- Es un BFS normal pero encola sólo grados de entradas 0
"""

def orden_bfs(grafo: Grafo):
    """
    orden_bfs toma un grafo y devuelve su recorrido de padres en orden topologico
    """
    


def calcular_grados_entrada(grafo: Grafo) -> dict:
    """
    calcular_grados_entrada
    devuelve un diccionario con grados de entrada de cada vertice
    """
    grados = dict()
    for v in grafo.obtener_vertices():
        grados[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.obtener_vertices():
            if (v, 0) in grafo.adyacentes(w):
                grados[v] += 1
    return grados



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
    g_entrada = calcular_grados_entrada(grafo_materias)
    print(g_entrada)