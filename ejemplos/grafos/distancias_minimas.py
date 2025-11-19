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
        visitados.add(v)

        for w, dist in grafo.adyacentes(v):
            nueva_distancia = distancias[v] + dist 
            if distancias[w] == None or distancias[w] > nueva_distancia:
                distancias[w] = nueva_distancia
                padres[w] = v
                heap.encolar((nueva_distancia, w))
    return distancias, padres

"""
# Algoritmo Bellman-Ford
- Inicializar estructuras a utilizar: padres, distancias
- Inicializar distancias: todos en infinito menos el primer elemento
- Obtener vertices

Repetir v - 1 veces:
    - Para cada arista (v, w, peso) de ese vertice:
        - si la distancia es menor a distancia actual: -> if dist[v] + peso_vw < dist[w]
            - actualizar padres + distancia
- devolver distancias, padres

Complejidad O(v + e) + O(V*E)
"""

def bellman_ford(grafo: Grafo, v):
    distancias = dict()
    padres = dict()
    aristas = obtener_aristas(grafo)
    
    for w in grafo.obtener_vertices():
        distancias[w] = float('inf')
    distancias[v] = 0
    padres[v] = None
    
    for _ in range(grafo.obtener_cantidad() - 1):
        for a in aristas:
            v, w, dist_vw = a
            if distancias[v] + dist_vw < distancias[w]: 
                distancias[w] = distancias[v] + dist_vw
                padres[w] = v
                
    return distancias, padres        


def obtener_aristas(grafo: Grafo) -> list:
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w, peso in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, peso))
        visitados.add(v)
    return aristas


"""
Centralidad: Betweeness

La centralidad de un vertice es la cantidad de veces que aparece dentro de caminos minimos en un grafo.

Procedimiento:
    - Para todos los posibles (v, w) dentro del grafo le busco el camino minimo
        - "Reconstruiyo el camino" pero a cada vertice del medio le agrego uno a la centralidad
"""

def obtener_centralidad(grafo: Grafo) -> list:
    centralidad = dict()
    for v in grafo.obtener_vertices():
        centralidad[v] = 0
    
    for v in grafo.obtener_vertices(): 
        for w in grafo.obtener_vertices():
            # No tiene sentido de un vertice a si mismo
            if v == w:
                continue
            
            distancia, padres = camino_minimo_dijkstra(grafo, v, w)
            
            # Si no hay camino, skip
            if not padres:
                continue
            
            # Actualizar centralidad para vertices en medio
            actual = padres[w]
            while actual != v:
                centralidad[actual] += 1
                actual = padres[actual]
    return centralidad
            

def camino_minimo_dijkstra(grafo: Grafo, origen, destino):
    cola = Heap()
    distancias = dict()
    padres = dict()
    visitados = set()
    
    for v in grafo.obtener_vertices():
        distancias[v] = float('inf')
    distancias[origen] = 0
    padres[origen] = None
    cola.encolar((0, origen))
    
    while not cola.esta_vacia():
        _, v = cola.desencolar()

        if v == destino:
            return distancias, padres        
        if v in visitados:
            continue
        visitados.add(v)
        
        for w, peso_vw in grafo.adyacentes(v):
            nueva_dist = distancias[v] + peso_vw
            if nueva_dist < distancias[w]:
                distancias[w] = nueva_dist
                padres[w] = v
                cola.encolar((nueva_dist, w))
    # Si llega hasta aca no hay camino hasta alla
    return None, None


"""
Grafo de ejemplo imagenes/distancias_ejemplo
"""
def crear_grafo() -> Grafo:
    grafo = Grafo(False)
    
    for x in range(7):
        grafo.agregar_vertice(x)
        
    grafo.agregar_arista(0,1,2)
    grafo.agregar_arista(0,2, 6)
    grafo.agregar_arista(1,3,5)
    grafo.agregar_arista(2, 3, 8)
    grafo.agregar_arista(3, 5, 15)
    grafo.agregar_arista(3, 4, 10)
    grafo.agregar_arista(4, 6, 2)
    grafo.agregar_arista(5, 6, 6)
    
    return grafo

if __name__ == '__main__':
    grafo = crear_grafo()
    dist_dijkstra, padres_dikstra = dijkstra(grafo, 0)
    print(dist_dijkstra, padres_dikstra)
    dist_bellmanford, padres_bellmanford = bellman_ford(grafo, 0)
    print(dist_bellmanford, padres_bellmanford)
    print(obtener_centralidad(grafo))