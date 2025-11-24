from grafo import Grafo
from collections import deque
import heapq


def camino_minimo(grafo: Grafo, origen, destino) -> list:
    """
    Devuelve un diccionario de distancias y de padres del origen a destino.
    con el camino mínimo desde origen hasta destino de un grafo no pesado.
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


def diametro(grafo) -> int:
    """
    Toma un grafo y devuelve su Diametro.
    
    El diametro de un grafo se define como la máxima distancia de todas las distancias minimas.
    Complejiadad O(V*(V + E))
    """
    max_dist = 0
    for v in grafo:
        # O(v + e)
        distancias, padres = camino_minimo(grafo, v, None)
        # Recorrer los vertices O(v)
        for w in grafo:
            if w == v:
                continue
            # Si distancia(v -> w) mayor a diametro actual, actualizar diametro
            if max_dist < distancias[w]:
                max_dist = distancias[w]
    return max_dist


def en_rango(grafo: Grafo, origen, rango: int) -> int:
    """
    Toma de parametro un grafo, vértice origen y un rango.
    
    Devuelve la cantidad de vértices que se encuentran a ese rango del origen.
    """



if __name__ == '__main__':
    pass