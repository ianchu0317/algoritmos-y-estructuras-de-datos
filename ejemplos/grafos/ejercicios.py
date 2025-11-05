from grafos import Grafo
import random

"""
# Ejercicio: dado un grafo no dirigido, ver si hay ciclo o no
Procedimiento: Mejor utilizar dfs -> 
si estoy en un nodo y su adyacente ya visitado y no es el padre,entonces hay ciclo, sino no
"""

def hay_ciclo(grafo: Grafo) -> bool:
    # Estructuras auxiliares 
    padres = dict()
    visitados = set()
    # Elegir un vertice random (como es no dirigido no importa)
    v = random.choice(grafo.obtener_vertices())
    padres[v] = ""
    visitados.add(v)
    
    return _dfs_hay_ciclo(v, grafo, padres, visitados)

def _dfs_hay_ciclo(v, grafo: Grafo, padres: dict, visitados: set) -> bool:
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            _dfs_hay_ciclo(v, grafo, padres, visitados)
        elif w != padres[v]:
            return True
    return False

