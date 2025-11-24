from grafo import Grafo
from collections import deque
import heapq


def camino_minimo(grafo: Grafo, origen, destino) -> list:
    """
    Devuelve una lista con el camino mínimo desde origen hasta destino de un grafo no pesado.
    En caso de no existir el camino se devuelve una lista vacía.
    
    **Parámetros**: origen y destino
    
    **Complejidad**: O(V + E)
    """
    
    padres = dict()
    cola = deque()
    visitados = set()
    
    cola.append(origen)
    padres[origen] = None
    visitados.add(origen)
    
    while len(cola) > 0:
        v = cola.popleft()

        if v == destino:
            return reconstruir_camino(padres, origen, destino)
        
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.append(w)
                padres[w] = v
    
    # Si llega hasta acá no encontró camino
    return []


def reconstruir_camino(padres: dict, ini, fin) -> list:
    """
    Toma un diccionario de padres y devuelve una lista con camino desde 'ini' hasta 'fin'
    
    Complejidad O(V)
    """
    camino = [fin]
    while fin != ini:
        fin = padres[fin]
        camino.append(fin)
    return list(reversed(camino))


def diametro(grafo: Grafo) -> int:
    """
    Toma un grafo y devuelve su Diametro.
    
    El diametro de un grafo se define como la máxima distancia de todas las distancias minimas.
    Complejiadad O(V*(V + E))
    """
    max_dist = 0
    for v in grafo:
        distancias = bfs_distancias(grafo, v)
        max_dist = max(max_dist, distancias.values())
    return max_dist


def bfs_distancias(grafo: Grafo, origen) -> dict:
    """
    Toma un grafo y un origen y devuelve las distancias minimas desde un origen hasta todos los demas vertices.
    
    **Ejemplo**
    ```python
    distancias = bfs_distancias(grafo, v) # Distancias es un dict
    ```
    """
    distancias = dict()
    cola = deque()
    visitados = set()
    
    distancias[origen]
    visitados.add(origen)
    cola.append(origen)
    
    while len(cola) > 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                distancias[w] = distancias[v] + 1
                visitados.add(w)
                cola.append(w)
    return distancias
    



def en_rango(grafo: Grafo, origen, rango: int) -> int:
    """
    Toma de parametro un grafo, vértice origen y un rango.
    
    Devuelve la cantidad de vértices que se encuentran a ese rango del origen.
    """



if __name__ == '__main__':
    pass