from grafos import Grafo
from tdas.heap import Heap
from tdas.cola import Cola

"""
Ejercicio 1: 
Implementar un algoritmo que reciba un grafo dirigido y acíclico 
y determine si dicho grafo admite un único orden topológico. 

Indicar compeljidad de la funcion.
"""

def ejercicio_1(grafo: Grafo):
    """
    El punto clave en este ejercicio es ver que para que haya un solo orden topologico,
    al iterarlo como orden topológico, la estructura auxiliar (pila/cola) debe tener un solo elemento.
    Si no verifica eso devolver False, sino True.
    
    Como es un orden topológico normal, la complejidad queda O(v + e)
    """
    cola = Cola()
    g_entrada = grados_entrada(grafo)
    
    for v in grafo:
        if g_entrada[v] == 0:
            cola.encolar(v)
    
    while not cola.esta_vacia():
        v = cola.desencolar()
    
        if not cola.esta_vacia():
            return False
    
        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.encolar(w)

    return True

def grados_entrada(grafo: Grafo):
    grados = dict()
    for v in grafo:
        grados[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados


"""
Ejercicio 2:
Se tiene una app de tránsito que indica el mejor camino desde un punto A a un punto B.
Esta app trabajoa con un grafo dirigido de las calles de la ciudad,
en donde las esquinas son los vçertices del Grafo, y aristas
son las calles que unen dichas esquinas. Las aristas tienen como peso: el tiempo, y su nombre.
Además hay un informe respecto al estado de las calles que indica para cada uno 3 posibles estados:
- Habilitada: se puede transitar de forma normal
- Congestionada: Se tarda el doble en transitar
- Cortada: imposible transitar

Se pide implementar un algoritmo que reciba el grafo, el informe 
(un diccionario en el que las claves son los nombre de aristas y valores su estado correspondiente), 
un punto A y un punto B, y determine la forma  más rápida para llegar desde A hasta B,
considerando el estado dado por el informe
"""

def ejercicio_2(a, b, grafo: Grafo, informe: dict):
    """
    El punto clave de este ejercicio es resolver con dijkstra aprovechando que son positivos los tiempsos.
    La idea es: ir encolando a la cola de prioridad dependiendo del informe: 
        Si no está habilitada, saltear ese arista
        si está congestionada, multiplicar x2 el tiempo
        Si está habilidada, continuar como corresponde
    La complejidad es O(e log(v) por dijkstra) + O(v) por reconstruir camino
    """
    cola = Heap()
    dist = dict()
    padres = dict()
    visitados = set()
    
    for v in grafo:
        dist[v] = float('inf')
    dist[a] = 0
    padres[a] = None
    cola.encolar((0, a))
    
    while not cola.esta_vacia():
        _, v = cola.desencolar()
        
        if v in visitados:
            continue
        visitados.add(v)

        if v == b:
            break
        
        for w, peso_vw in grafo.adyacentes(v):
            nombre, tiempo = peso_vw
            if informe[nombre] == "CORTADA":
                continue
            elif informe[nombre] == "CONGESTIONADA":
                tiempo = tiempo * 2
            
            if dist[v] + tiempo < dist[w]:
                dist[w] = tiempo + dist[v]
                padres[w] = v
                cola.encolar((dist[w], w))

    return reconstruir_camino(padres, a, b)

def reconstruir_camino(padres, a, b):
    camino = [b]
    while b != a:
        b = padres[b]
        camino.append(b)
    return reversed(camino)


"""
Ejercicio 3:
Implementar un algoritmo que reciba un grafo no dirigido y determine
la cantidad máxima de aristas que se pueden agrgar al mismo de tal forma que no se reduizcan la cantidad de componentes conexas
que hay en el mismo. 

Indicar y justificar la complejidad del algoritmo implementeado.
"""
def ejercicio_3(grafo: Grafo) -> int:
    """
    La idea para este ejercicio es iterar por cada componente:
        Para cada componente:
            Recorrer la compoentne ya sea con BFS, DFS O(v + e)
            - Calcular la cantidad de vertices totales
            - Calcular la cantidad de aristas totales
            Calcular aristas maximos disponibles para componente = v*(v-1)/2 - aristas_componente
    return aristas_maximos_disponibles
    """

    visitados = set()
    cant_max_aristas = 0
    for v in grafo:
        if v not in visitados:
            cant_max_aristas += bfs_hallar_max_aristas(v, grafo, visitados)
    return cant_max_aristas

def bfs_hallar_max_aristas(v, grafo: Grafo, visitados: set):
    cola = Cola()
    vertices = 0
    aristas = 0
    
    visitados.add(v)
    cola.encolar(v)

    while not cola.esta_vacia():
        v = cola.desencolar()
        vertices += 1
            
        for w in grafo.adyacentes(v):
            aristas += 1
            if w not in visitados:
                cola.encolar(w)
                visitados.add(w)
    
    aristas //= 2
    return (vertices * (vertices - 1)//2) - aristas
        
    
    
"""
Ejercicio 4:
Implementar un algoritmo que, dado un grafo pesado con pesos positivos, un vértice v, y otro w,
determine la cantidad de caminos mínimos que hay de v a w dentro del grafo. Considerar que, justamente
podrían haber varios caminos de una misma distancia, que a su vez sean la mínima. 

Inidicar y justificar la complejidad del algoritmo implementado
"""

def ejercicio_4(grafo: Grafo, a, b) -> int:
    """
    La idea en este ejercicio, como nos pide grafo PESADO con PESOS POSITIVOS y hallar CAMINO MINIMO,
    lo primero que tenemos que pensar es en Dijkstra. Primero lo implementamos y luego realizamos el cambio.
    
    Observamos que si v == w entonces quiere decir que ya llegamos al destino, entonces devolver la cantidad de caminos.
    Como se calcula un nuevo camino?
        Cuando tenemos nueva_dist < dist[w] quiere decir que tenemos un nuevo camino mínimo para llegar a w
            Entonces caminos[w] = 1 -> OJO la idea es ser igual la cantidad de caminos para llegar al padre
            Entonces caminos[w] = caminos[v]
        Pero si tenemos nueva_dist == dist[w] quiere decir que hay otro camino con misma distancia para llegar a w
            Entonces caminos[w] += 1
    """
    
    heap = Heap()
    dist = dict()
    visitados = set()
    caminos = dict()
    
    for v in grafo:
        dist[v] = float('inf')
    dist[a] = 0
    heap.encolar((a, 0))
    caminos[a] = 1
    
    while not heap.esta_vacia():
        v, _ = heap.desencolar()

        if v == b:
            return caminos[b]
        
        if v in visitados:
            continue
        visitados.add(v)    
    
        for w, peso_vw in grafo.adyacentes(v):
            nueva_dist = dist[v] + peso_vw
            if nueva_dist < dist[w]:
                dist[w] = nueva_dist
                heap.encolar((w, nueva_dist))
                caminos[w] = caminos[v]
            elif nueva_dist == dist[w]:
                caminos[w] += caminos[v]
    # Si llega acá es que nunca encontró a B
    return 0
