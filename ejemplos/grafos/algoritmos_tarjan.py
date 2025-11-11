from grafos import Grafo


"""
Algoritmo de Tarjan: 
- Inicializar estructuras auxiliares
- Inicializar variable inicial
- LLamada para primer vertice (y todos los demas no visitados)
- DFS
    - Actualizar variables del vertice actual
    - Para cada adyacente no visitado:
        - Actualizar variables vertice a visitar
        - LLamado recursivo
        - Actualizar mas_bajo vertice adyacente
        - Chequear si el vertice V es articulacion
CONDICION:
 Sea inicio de ciclo y tiene mas de un hijo
 - No sea raiz de dfs 
"""


def tarjan_puntos_articulacion(grafo: Grafo):
    # Inicializar estructuras a utilizar
    puntos_articulacion = set()
    visitados = set()
    mas_bajo = dict()
    orden = dict()
    padres = dict()
    
    # Realizar para todos los vertices del grafo no visitados (por si hay comp no conexo)
    for v in grafo.obtener_vertices():
        if v not in visitados:
            padres[v] = None
            orden[v] = mas_bajo[v] = 0
            visitados.add(v)
            dfs_puntos_articulacion(v, grafo, puntos_articulacion, visitados, padres, mas_bajo, orden)
    return puntos_articulacion

def dfs_puntos_articulacion(v, grafo: Grafo, pt_a: set, visitados: set, padres: dict, mas_bajo: dict, orden: dict):
    hijos = 0
    for w, _ in grafo.adyacentes(v):
        if w not in visitados:
            hijos += 1
            orden[w] = mas_bajo[w] = orden[v] + 1
            padres[w] = v
            visitados.add(w)
            dfs_puntos_articulacion(w, grafo, pt_a, visitados, padres, mas_bajo, orden)
            
            ### CONDICIONES
            # Si es raiz de dfs y tiene mas de un hijo
            if padres[v] is None and hijos > 1:
                pt_a.add(v)
                
            # Si no es raiz de dfs y el mas bajo de w es mayor 
            # -> quiere decir que solo puede pasar por ese camino (v) - > solo se incrementa con ese camino
            if padres[v] is not None and mas_bajo[w] >= orden[v]:
                pt_a.add(v)
                
            # Actualizar mas bajo de V entre mas bajo de W            
            mas_bajo[v] = min(mas_bajo[w], mas_bajo[v])
            
        # SI hay uno adyacente visitado y no es padre entonces me quedo con su mb (el mas bajo)
        elif w != padres[v]:
            mas_bajo[v] = min(mas_bajo[w], orden[v])

def crear_grafo_pto_articulacion():
    grafo = Grafo(False)
    
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_vertice("D")
    grafo.agregar_vertice("E")
    
    grafo.agregar_arista("A", "B")
    grafo.agregar_arista("B", "C")
    grafo.agregar_arista("C", "E")
    grafo.agregar_arista("E", "D")
    grafo.agregar_arista("D", "C")
    
    return grafo

def crear_grafo_cfc():
    pass


if __name__ == '__main__':
    grafo_pt_a = crear_grafo_pto_articulacion()
    print(tarjan_puntos_articulacion(grafo_pt_a))
