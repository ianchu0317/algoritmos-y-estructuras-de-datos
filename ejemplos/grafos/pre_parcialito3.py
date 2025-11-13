from grafos import Grafo
from tdas.heap import Heap

"""
Ejercicio 1: 
Implementar un algoritmo que reciba un grafo dirigido y acíclico 
y determine si dicho grafo admite un único orden topológico. 

Indicar compeljidad de la funcion.
"""

def ejercicio_1(grafo: Grafo):
    """
    El punto clave en este ejercicio es ver que para que haya un solo orden topologico,
    Tiene que ser una lista
    """
    pass


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
            DFS
            - Calcular la cantidad de vertices totales
            - Calcular la cantidad de aristas totales
            Calcular aristas maximos disponibles para componente = v*(v-1)/2 - aristas_componente
    return aristas_maximos_disponibles
    """

    visitados = set()
    cant_max_aristas = 0
    for v in grafo:
        if v not in visitados:
            componente = set()
            
    return cant_max_aristas

"""
Ejercicio 4:
Implementar un algoritmo que, dado un grafo pesado 8con pesos positivos, un vértice v, y otro w,
determine la cantidad de caminos mínimos que hay de v a w dentro del grafo. Considerar que, justamente
podrían haber varios caminos de una misma distancia, que a su vez sean la mínima. 

Inidicar y justificar la complejidad del algoritmo implementado
"""