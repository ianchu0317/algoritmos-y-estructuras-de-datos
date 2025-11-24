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


def diametro(grafo: Grafo) -> list:
    """
    Toma un grafo y devuelve el camino de su diametro.
    
    El diametro de un grafo se define como la máxima distancia de todas las distancias minimas.
    
    ****
    
    Complejiadad O(V*(V + E))
    """
    max_dist_padres = dict()
    max_dist = 0
    
    for v in grafo:
        distancias, padres = bfs_distancias_padres(grafo, v)
        for w in distancias:
            if v == w:
                continue
            # Si la distancia v -> w mayor a diamtro actual, actualizar
            if max_dist < distancias[w]:
                max_dist = distancias[w]
                # Para optimizar copiar una sola vez                
                if max_dist_padres != padres:
                    max_dist_padres = padres
                max_dist_origen = v
                max_dist_destino = w 
    return reconstruir_camino(max_dist_padres, max_dist_origen, max_dist_destino)


def bfs_distancias_padres(grafo: Grafo, origen):
    """
    Toma un grafo y un vertice de origen y devuelve 
    las distancias minimas y dicc de padres desde ese vertice hasta todos los demas vertices.
    
    **Ejemplo**
    ```
    distancias, padres = bfs_distancias_padres(grafo, v) # Distancias es un dict
    ```
    """
    distancias = dict()
    cola = deque()
    visitados = set()
    padres = dict()
    
    padres[origen] = None
    distancias[origen] = 0
    visitados.add(origen)
    cola.append(origen)
    
    while len(cola) > 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                distancias[w] = distancias[v] + 1
                visitados.add(w)
                cola.append(w)
                
    return distancias, padres
    


def en_rango(grafo: Grafo, origen, rango: int) -> int:
    """
    Toma de parametro un grafo, vértice origen y un rango.
    
    En caso de no encontrar, devuelve una lista vacia.
    
    Devuelve la cantidad de vértices que se encuentran a ese rango del origen.
    """
    cant_en_rango = 0
    distancias, _ = bfs_distancias_padres(grafo, origen)
    for v in distancias:
        if distancias[v] == rango:
            cant_en_rango += 1
    return cant_en_rango



if __name__ == '__main__':
    pass