from grafos import Grafo
from tdas.heap import Heap


"""
# Algoritmo Dijkstra
- Crear estructuras auxiliares: visitados, heap, padres, distancias
- Inicializar estructuras:
    - Marco todos los vertices con distancia infinito
    - El primer vertice tiene distancia 0
    - Encolar primer vertice con distancia 0
- Mientras la cola no esta vacia:
    - v = desencolar
    Si v está analizado lo skipeo
    - Para cada adyacente (w) de V:
        - Si distancia[w] > distancia[v] + peso_vw
            - Distancia[w] = distancia[v] + peso_vw
            - padres[w] = v
            - cola.encolar(distancia[v] + peso_vw, w)
    - Marcar v como visitado
- Devolver distancias y padres

Complejidad O(V + (V + E)*log(V))
"""

def dijkstra(grafo: Grafo, v):
    # Estructuras auxiliares a utilizar
    visitados = set()
    distancias = dict()
    heap = Heap()
    padres = dict()
    
    # Inicializar estructuras auxiliares
    for w in grafo.obtener_vertices():
        distancias[w] = None
    distancias[v] = 0
    padres[v] = None
    heap.encolar((0, v))
    
    # Desencolar mientras no esta vacia
    while not heap.esta_vacia():
        _, v = heap.desencolar()

        if v in visitados:
            continue

        for w, dist in grafo.adyacentes(v):
            if w not in visitados:
                nueva_distancia = distancias[v] + dist 
                if distancias[w] == None or distancias[w] > nueva_distancia:
                    distancias[w] = nueva_distancia
                    padres[w] = v
                    heap.encolar((nueva_distancia, w))
        visitados.add(v)
    return distancias, padres


