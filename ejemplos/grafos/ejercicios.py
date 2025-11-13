from grafos import Grafo
from tdas.cola import Cola
import random

"""
# Ejercicio: dado un grafo no dirigido, ver si hay ciclo o no
Procedimiento: Mejor utilizar dfs -> 
si estoy en un nodo y su adyacente ya visitado y no es el padre,entonces hay ciclo, sino no
- > la funcion auxiliar devuelve si hay ciclos y donde comenzar a reconstruir
"""

def hay_ciclo(grafo: Grafo) -> list:
    # Estructuras auxiliares 
    padres = dict()
    visitados = set()
    # Elegir un vertice random (como es no dirigido no importa)
    v = random.choice(grafo.obtener_vertices())
    padres[v] = ""
    visitados.add(v)
    
    # Reconstruir camino si hay ciclo
    hay_ciclo, w = _dfs_hay_ciclo(v, grafo, padres, visitados)
    if hay_ciclo:
        return reconstruir_camino(w, padres)
    return None

def _dfs_hay_ciclo(v: str, grafo: Grafo, padres: dict, visitados: set) -> tuple[bool, str]:
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            _dfs_hay_ciclo(v, grafo, padres, visitados)
        elif w != padres[v]:
            return True, v
    return False, ""


def reconstruir_camino(v: str, padres: dict) -> list:
    camino = []
    while v != "":
        camino.append(v)
        v = padres[v]
    return camino


"""
Ejercicio 15:
Un autor decidió escribir un libro con varias tramas que se puede leer de forma no lineal. 
Es decir, por ejemplo, después del capítulo 1 puede leer el 2 o el 73; 
pero la historia no tiene sentido si se abordan estos últimos antes que el 1.

Siendo un aficionado de la computación, el autor ahora necesita un orden para publicar su obra, 
y decidió modelar este problema como un grafo dirigido, en dónde los capítulos son los vértices 
y sus dependencias las aristas. Así existen, por ejemplo, las aristas (v1, v2) y (v1, v73).

Escribir un algoritmo que devuelva un orden en el que se puede leer la historia sin obviar ningún capítulo. 
Indicar la complejidad del algoritmo.
"""

def orden_lectura(libro: Grafo) -> list:
    """
    orden_lectura() toma un libro y devuelve una lista [v1, v2, v3, ... v_n]
    
    Ya que el grafo es dirigido, donde cada vertice es un capitulo y las aristas son las dependencias,
    la idea es realizar un recorrido topológico. Para mi approach prefiero bfs topológico:
        - Contar grados de entradas de cada capitulo (la cantidad de capitulos dependentientes de cada) O(v + e)
        - Encolar los capitulos que no dependen de nadie. O(v) -> todos los vertices en el peor caso
        - Mientras haya elementos pendientes en la cola O(v + E):
            - Restar los adyacentes y encolarlos si ya no dependen de nadie
            - Agregar la lista el capitulo actual
        - Devolver la lista de orden
    Complejidad: O(v + e)
    """
    grado = calcular_dependencia_capitulo(libro)
    cola = Cola()
    
    for v in libro.obtener_vertices():
        if grado[v] == 0:
            cola.encolar(v)
    
    orden_capitulos = []
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w, _ in libro.adyacentes(v):
            grado[w] -= 1
            if grado[w] == 0:
                cola.encolar(w)
        orden_capitulos.append(v)
        
    return orden_capitulos


def calcular_dependencia_capitulo(grafo: Grafo) -> dict:
    grado = dict()
    for v in grafo.obtener_vertices():
        grado[v] = 0 
    for v in grafo.obtener_vertices():
        for w, _ in grafo.adyacentes(v):
            grado[w] += 1
    return grado


"""
Ejercicio 17: Implementar una función que reciba un grafo no dirigido, y que compruebe la siguiente afirmación: 
“La cantidad de vértices de grado IMPAR es PAR”. 
Indicar y justificar el orden del algoritmo si el grafo está implementado como matriz de adyacencia.
"""

def comprobar_afirmacion(grafo: Grafo) -> bool:
    """
    comprobar_afirmacion(grafo) toma un grafo y comprueba si cumple "La cantidad de vertices de grado impar es par"
    
    Para resolver primero calculamos el grado de cada vertice O(V²), 
    luego iteramos cada vertice contando la cantidad de vertices con grado impar. O(v)
    Finalmente devolvemos el booleano con la afirmacion -> True si es afirmativo, False si es Falso
    
    La complejidad final es O(v²)
    """
    grados = calc_g_entrada_no_dirigido(grafo)  #  O(v²)
    cantidad_v_grado_impar = 0
    
    # Iterar para cada vertice O(v)
    for v in grafo.obtener_vertices():
        if grados[v] % 2 == 1:
            cantidad_v_grado_impar += 1
    
    return cantidad_v_grado_impar % 2 == 0

def calc_g_entrada_no_dirigido(grafo: Grafo) -> dict:
    grados = dict()
    # Iterar para cada vertice O(v)
    for v in grafo.obtener_vertices():
        # Calcular cantidad de aristas O(v)
        grados[v] = len(grafo.adyacentes(v))
    return grados