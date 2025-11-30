from grafo import Grafo
from collections import deque
from random import shuffle
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


def camino_primer_adyacente(grafo: Grafo, origen) -> list:
    """
    Toma un grafo y un vértice de origen.
    Devuelve el camino formado yendo por el primer adyacente.
    Se detiene cuando haya más adyacentes o llegado a 20 paginas
    
    **Complejiad**: O(n)
    """
    _LIMITE_PRIMER_ADY = 20
    contador = 0
    camino = [origen]
    while (len(grafo.adyacentes(origen)) > 0) and (contador < _LIMITE_PRIMER_ADY):
        origen = grafo.adyacentes(origen)[0]    
        camino.append(origen)
        contador += 1
    return camino


def obtener_cfcs(grafo: Grafo) -> list:
    """
    Toma un grafo dirigido y devuelve las componentes fuertemente conexas.
    
    **Complejidad**: O(V + E)
    """
    cfcs = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            _dfs_cfcs(v, grafo, visitados, set(), deque(), dict(), dict(), [0], cfcs)
    return cfcs


def _dfs_cfcs(v, grafo: Grafo, visitados: set, 
              apilados: set, pila: deque, orden: dict, 
              mb: dict, contador: list, cfcs: list):
    # Actualizar para vertice
    visitados.add(v)
    orden[v] = mb[v] = contador[0]
    contador[0] += 1
    pila.append(v)
    apilados.add(v)
    
    for w in grafo.adyacentes(v):
        if w not in visitados:
            _dfs_cfcs(w, grafo, visitados, apilados, pila, orden, mb, contador, cfcs)
        if w in apilados:
            mb[v] = min(mb[w], mb[v])
    
    # Crear camino componente conexo
    if mb[v] == orden[v]:
        cfc = set()
        while True:
            w = pila.pop()
            apilados.remove(w)
            cfc.add(w)
            if w == v:
                break
        cfcs.append(cfc)
            

def orden_topologico_vertices(grafo: Grafo, vertices: list) -> list:
    """
    Toma un grafo y una lista de vertices.
    Devuelve un orden topológico de esos vertices según el grafo pasado.
    
    **Consideraciones**
    Para que haya orden topológico entre v_i y v_j de la lista de vertices,
    entonces debe haber un arista entre esos dos vertices.
    
    **Complejidad** O(Vn + En), Vn la cantidad vertices y En la cantidad de aristas que hay
    """
    g_entrada = grados_entrada(grafo, vertices)
    cola = deque()
    orden = list()
    
    for v in vertices:
        if g_entrada[v] == 0:
            cola.append(v)
            
    while len(cola) > 0:
        v = cola.popleft()
        orden.append(v)
        
        for w in grafo.adyacentes(v):
            if w in g_entrada:
                g_entrada[w] -= 1
                if g_entrada[w] == 0:
                    cola.append(w)
    return orden


def grados_entrada(grafo: Grafo, vertices: list) -> dict:
    """
    Toma un grafo y una lista de vertices perteneciente al grafo.
    Devuelve grados de entrada de esos vertices en esa sección del grafo.
    Si grafo.hay_arista(v_i, v_j) entonces v_i -> v_j. 
    Si ningun vertice de la lista pasada conecta a v_i, entonces grado de entrada de v_i es 0. 
    """
    g_entrada = dict()
    for v in vertices:
        g_entrada[v] = 0
    for v in vertices:
        for w in grafo.adyacentes(v):
            if w in g_entrada:
                g_entrada[w] += 1
    return g_entrada


def comunidades(grafo: Grafo) -> list:
    """
    Toma un grafo y devuelve las comunidades dentro del grafo.
    
    **Complejidad**: O(K(V + E)), donde K es la cantidad de iteraciones 
    """
    # Obtener vertices de entrada O(V + E)
    vertices_entrada = obtener_vertices_entrada(grafo)
    
    # Asignar label para cada vertice del grafo
    label = dict()
    i = 0
    for v in grafo:
        label[v] = i
        i += 1
    
    # Repetir hasta condicion de corte O(k)
    _ITERACION = 10
    for x in range(_ITERACION):
        # Obtener orden random O(V)
        orden_random = grafo.obtener_vertices()
        shuffle(orden_random)    
        # Hacer para cada vertice del grafo en orden aleatorio O(V + E)
        for v in orden_random:
            if not vertices_entrada[v]:
                continue
            label[v] = max_freq(label, vertices_entrada[v])
    
    # Reconstruir comunidades
    # Para cada vertice que tiene mismo label es de la misma comunidad
    comunidades_dicc = {}
    for v in grafo:
        v_label = label[v]
        if v_label not in comunidades_dicc:
            comunidades_dicc[v_label] = set()
        comunidades_dicc[v_label].add(v)
    return list(comunidades_dicc.values())


def max_freq(label: dict, v_entrada: list) -> int:
    """
    Toma un diccionario de label y devuelve el mayor label de v_entrada
    
    Complejidad O(E) -> TOdos los grados de entrada de un vertice
    """
    contador_label = dict()
    for v in v_entrada:
        v_label = label[v]
        contador_label[v_label] = contador_label.get(v_label, 0) + 1
    return max(contador_label, key=contador_label.get)



def obtener_vertices_entrada(grafo: Grafo) -> dict:
    """
    Devuelve un diccionario que contiene para cada vertice del grafo, una lista con vertices de entrada.    
    
    **Complejidad**: O(V + E)
    """
    vertices_entrada = dict()
    for v in grafo:
        vertices_entrada[v] = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            vertices_entrada[w].append(v)
    return vertices_entrada


def clustering_vertice(grafo: Grafo, v) -> float:
    """
    Dado un grafo y un vértice del grafo, devuelve el coeficiente de Clustering de ese vertice.
    Si vértice tiene menos de dos adyacentes, entonces devuelve 0.0
    
    **Complejidad** O(k²) pero si K es acotado y con un rango pequeño, entonces O(C) constante y se acota O(1)
    
    Si no es grafo disperso, entonces es O(E) en el peor de los casos
    """
    
    vecinos_v = grafo.adyacentes(v)
    k = len(vecinos_v)
    
    if k < 2:
        return 0.0

    contador_aristas = 0
    # POr cada adyacente, me fijo si hay un par con arista
    for i in range(k):
        w = vecinos_v[i]
        for j in range(k):
            if j == i:
                continue
            u = vecinos_v[j]
            if grafo.hay_arista(w, u):
                contador_aristas += 1
    
    return round(contador_aristas / (k*(k-1)), 3)


def clustering(grafo: Grafo) -> float:
    """
    Toma un grafo y devuelve el coeficiente de clustering promedio de todo el grafo.
    
    **Complejidad**
    Para cada vertice calculo su coeficiente O(v*k²)
    - Si K es acotado (grafo disperso) entonces es O(v)
    - Si es grafo denso entonces (k aprox V) -> O(V³)
        - DUDA RECORDATORIO PREGUNTAR COMPLEJIDAD DE ESTE****
    """
    suma_clustering = 0.0
    cantidad = 0
    for v in grafo:
        suma_clustering += clustering_vertice(grafo, v)
        cantidad += 1
    return round(suma_clustering / cantidad, 3) 



def page_rank(grafo: Grafo, k: int, d: float) -> list:
    """
    Toma un grafo, un valor K (veces a iterar, cuanto más mejora presicion), un valor d (dumping factor),
    y un valor n.
    Devuelve un diccionario con los valores que rankean los vertices del grafo.
    
    **Complejidad**: O(k(V + E) 
    """
    # Estructuras auxiliares a utilizar
    pr = dict()
    v_entrada = obtener_vertices_entrada(grafo) # O(V + E) 
    N = len(grafo.obtener_vertices())           # O(V) obtener cantidad de vertices en grafo
    pr_nuevo = dict()
    
    # Inicializar pr de cada vertice
    for v in grafo:
        pr[v] = 1/N
    
    # Iterar k veces
    for _ in range(k):
        # Iterar O(V + E)
        for v in grafo:
            v_sumatoria = 0
            for w in v_entrada[v]:
                v_sumatoria += pr[w] / len(grafo.adyacentes(w))
            pr_nuevo[v] = (1-d)/N + d*v_sumatoria 
        pr = pr_nuevo       # Actualizar valores en cada iter      
    return pr


def obtener_top_n(diccionario: dict, n: int):
    """
    Toma un diccionario y devuelve una lista ordenada de top n elemntos del diccionario
    Complejidad: O(V(log n))
    """
    # Devolver lista n-elementos importantes -> Meter todo a heap -> min heap top-n
    top_n = []
    i = 0
    for v, rango_v in diccionario.items():
        # Mantener con n-eleemntos
        if i < n:
            heapq.heappush(top_n, (rango_v, v))
            i += 1
        # Mantener top-n
        elif rango_v > top_n[0][0]:
            heapq.heappop(top_n)
            heapq.heappush(top_n, (rango_v, v))    
    
    # Desencolar de heap y revertir
    res = []
    while len(top_n) > 0:
        _, v = heapq.heappop(top_n)
        res.append(v)        
    res.reverse()
    return res


def ciclo_largo_n(grafo: Grafo, origen, n: int) -> list:
    """
    Devuelve un camino de un ciclo de largo n.
    En caso de no encontrar camino devuelve una lista vacia.
    
    Complejidad O(V^n):
    - Busco con dfs hasta la profundidad deseada.
    - SI no encuentro lo borro... En el peor de los casos cada nivel tiene V vertices
        - V x V x ... Vn = V^n
    """
    
    if n < 1:
        return []
    
    if n > len(grafo.obtener_vertices()):
        return []

    padres = dict()
    visitados = set()
    
    padres[origen] = None
    visitados.add(origen)
    
    return _dfs_ciclo_largo_n(grafo, origen, origen, padres, visitados, n)

def _dfs_ciclo_largo_n(grafo: Grafo, v, origen, 
                       padres: dict, visitados: set, 
                       pasos_restantes: int):
    # Cortar recursividad si me paso de n ->> sino se me hace infinito xd
    if pasos_restantes == 0:
        return []
    
    for w in grafo.adyacentes(v):
        # Reconstruir camino si se encontro ciclo
        if pasos_restantes == 1 and w == origen:
            ciclo = reconstruir_camino(padres, origen, v)
            ciclo.append(origen)
            return ciclo
        
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            ciclo = _dfs_ciclo_largo_n(grafo, w, origen, padres, visitados, pasos_restantes-1)
            if ciclo:
                return ciclo
            # Borrar visitados para otros posibles caminos
            visitados.remove(w)
            padres.pop(w)
    return []


if __name__ == '__main__':
    pass