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
    #    -> realizar para todos los vertices del grafo (por si hay componente no conexo)
    for v in grafo:
        if v not in visitados:
            padres[v] = None
            visitados.add(v)
            hay_ciclo, inicio, fin = _dfs_hay_ciclo(v, grafo, padres, visitados)
            if hay_ciclo:
                return reconstruir_camino(padres, inicio, fin)
    return None

def _dfs_hay_ciclo(v: str, grafo: Grafo, padres: dict, visitados: set):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            padres[w] = v
            resultado = _dfs_hay_ciclo(w, grafo, padres, visitados)
            if resultado[0]:
                return resultado
        elif w != padres[v]:
            return True, v, w
    return False, None, None


def reconstruir_camino(padres: dict, inicio, fin) -> list:
    camino = []
    camino.append(inicio)
    while inicio != fin:
        inicio = padres[inicio]
        camino.append(inicio)
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



"""
Ejercicio 21:
Contamos con un grafo dirigido que modela un ecosistema. 
En dicho grafo, cada vértice es una especie, y cada arista (v, w) indica que v es depredador natural de w. 

Considerando la horrible tendencia del ser humano por llevar a la extinción especies, 
algo que nos puede interesar es saber si existe alguna especie que, si llegara a desaparecer, 
rompería todo el ecosistema: quienes la depredan no tienen un sustituto (y, por ende, pueden desaparecer también) 
y/o quienes eran depredados por esta ya no tienen amenazas, por lo que crecerán descontroladamente. 

Implementar un algoritmo que reciba un grafo de dichas características 
y devuelva una lista de todas las especies que cumplan lo antes mencionado. 

Indicar y justificar la complejidad del algoritmo implementado.
"""


def especies_ecosistema(ecosistema: Grafo) -> set:
    """
    especies_ecosistema devuelve una lista de especies a tener en cuenta

    los puntos claves del enunciado son:
        - Si un depredador solo tiene una presa, entonces presa es importante 
            -> si v solo tiene 1 grado salida, entonces w es importante 
        - Si una presa solo tiene un depredador, entonces depredador es importante
            -> si w solo tiene 1 grado entrada, entonces v es importante
    
    Pasos:
        - calcular grados de entrada O(v + e) y salida O(v + e) de cada vertice
        - Iterar cada vertice y sus adyacentes para chequear condiciones O(v + e)
        - Devolver resultado O(1)
    Complejidad O(v + e)
    """
    
    depredadores = cantidad_depredadores(ecosistema)
    presas = cantidad_presas(ecosistema)
    especies = set()
    
    for animal in ecosistema.obtener_vertices():
        for presa, _ in ecosistema.adyacentes(animal):
            if presas[animal] == 1:
                especies.add(presa)
            if depredadores[presa] == 1:
                especies.add(animal)
    
    return especies

def cantidad_presas(ecosistema: Grafo) -> dict:
    presas = dict()
    for animal in ecosistema.obtener_vertices():
        presas[animal] = len(ecosistema.adyacentes(animal))
    return presas

def cantidad_depredadores(ecosistema: Grafo) -> dict:
    depredadores = dict()
    for animal in ecosistema.obtener_vertices():
        depredadores[animal] = 0
    for animal in ecosistema.obtener_vertices():
        for presa, _ in ecosistema.adyacentes(animal):
            depredadores[presa] += 1
    return depredadores



"""
Ejercicio 23:
El diámetro de una red es el máximo de las distancias mínimas entre todos los vértices de la misma. 
Implementar un algoritmo que permita obtener el diámetro de una red, para el caso de un grafo no dirigido y no pesado.

Indicar el orden del algoritmo propuesto.
"""
def obtener_diametro(grafo: Grafo) -> int:
    """
    devolver diametro de un grafo (no dirigido y no pesado)
    
    Para cada vertice hallar la mayor distancia minima
    -> para cada v vertices realizar bfs para hallar el maximo distancia minimo de ese vertice a otros
    Para los valores hallados de esos hallar el maximo.
    
    Complejidad V * O(v + e) = O(v² + v*e)
    """
    
    diametro = 0
    
    for v in grafo.obtener_vertices():
        max_distancia_v = bfs_distancia_max(v, grafo)
        if max_distancia_v > diametro:
            diametro = max_distancia_v
        
    return diametro

def bfs_distancia_max(v, grafo) -> int:
    cola = Cola()
    visitados = set()
    distancias = dict()
    
    cola.encolar(v)
    visitados.add(v)
    distancias[v] = 0
    max_dist = 0
    
    while len(cola) > 0:
        v = cola.desencolar()
        
        for w, _ in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.encolar(w)
                distancias[w] = distancias[v] + 1
                if distancias[w] > max_dist:
                    max_dist = distancias[w]
                    
    return max_dist


"""
Ejercicio 32:
Se tiene un Grafo no dirigido G, y un árbol de tendido mínimo T, de G. 
Se obtiene una de las aristas de G, que no se encuentra en T, y se le reduce el peso. 
Dar un algoritmo lineal que permita determinar el nuevo árbol de tendido mínimo T'
(notar que T' podría ser igual a T). 

Justificar la complejidad del algoritmo.
"""

def obtener_mst_arista_nueva(mst: Grafo, nueva_arista):
    """
    Como es mst, una propiedad es que hay un solo camino de un vertice a otro.
    Si agrego otro awrista (v, w, peso_vw) entonces hay un nuevo camino a v->w y genera un ciclo.
    La idea entonces es eliminar al arista con mayor distancia de ese ciclo.
    
    Sin agregar el nuevo vertice todavía, hacer un bfs de v->w,
    y si en ese camino hay algun arista con peso mayor a nuevo_peso entonces peso=nuevo_peso y listo
    
    La complejidad queda lineal ya que es un BFS normal. 
    En el peor de los casos se visitan todos los vertices y aristas así que es O(V + E)
    
    
    ** Procedimiento ** 
    - BFS hasta encontrar arista con peso mayor al peso de nuevo arista en camiono de v->w
    - Si hay, borrar ese vertice y agregar el nuevo
    - Si no hay, devolver mst normal
    """
    
    v, w, nuevo_peso = nueva_arista
    v_max, w_max = buscar_arista_con_peso_mayor(v, w, nuevo_peso, mst)
    nuevo_mst = mst
    
    if v_max:
        nuevo_mst.borrar_arista(v_max, w_max)
        nuevo_mst.agregar_arista(v, w, nuevo_peso)
    
    return nuevo_mst
    

def buscar_arista_con_peso_mayor(origen, destino, nuevo_peso, grafo: Grafo):
    v_max = None
    w_max = None
    cola = Cola()
    visitados = set()
    
    peso_max = nuevo_peso
    cola.encolar(origen)
    visitados.add(origen)
    
    while not cola.esta_vacia():
        v = cola.desencolar()
        
        if v == destino:
            return v_max, w_max
        
        for w, peso_vw in grafo.adyacentes(v):
            if w not in visitados:
                cola.encolar(w)
                visitados.add(w)
                if peso_vw > peso_max:
                    peso_max = peso_vw
                    v_max = v
                    w_max = w
                
    return v_max, w_max


"""
Ejercicio 33: 
Como todos sabemos, Taller de Programación I implica programar mucho. 
El equipo de Federico esta desarrollando un Age of Empires y le asignaron el modulo de movimientos. 
Dispone de un mapa, podríamos decir una grilla, donde se puede ver para cada celda su contenido 
(por ejemplo si está libre, si es agua, un árbol, etc). Fede tiene que implementar un algoritmo que, 
a partir de una unidad (por ejemplo un soldado), en una celda en especifico encuentre el camino a una celda objetivo. 
Considerar que entre celdas puede haber diferencia de alturas 
por lo que puede costarle mas (o menos) a una unidad ir de una celda a otra. 
Por supuesto que no puede ser cualquier camino, si no el que haga que la unidad llegue más rápido a su objetivo.

a. Modelar el problema usando Grafos especificando de forma completa todos sus componentes.

b. Implementar un algoritmo que a partir de una celda de origen y una de destino, 
retorne el camino que tiene que hacer la unidad, indicando y justificando la complejidad final.
"""

def camino_rapido(grilla: dict, grafo: Grafo, origen, destino):
    """
    Como el enunciado ya menciona camino, entonces la idea es buscar el camino rapido: caminos minimos.
    Luego el dato menciona que puede costar más (pesos positivos) o costar menos (pesos negativos).
    Teniendo esto en cuenta, podemos saber más o menos que se trata de Bellman ford el algoritmo a implementar
    
    **Modelacion**
    Para modelar esto sabemos que cada celda esta conctada a otra celda por medio de caminos, entonces:
    Vertices: Celdas del mapa
    Aristas: Caminos que conectan distintas celdas
    PEso de arista: Costo de recorrer el camino para ir de una celda a otra
    
    **Pasos**
    - Obtener aristas disponibles O(v + e)
    - Bellman ford con las aristas disponibles "LIBRES" O(v * e)
    - Reconstruir camino O(E)
    
    La complejidad final es:
    O((v-1)*e) = O(v*e), que en terminos de la modelacion sería
    O(celdas*caminos)
    
    -> CAMBIAR MODELADO POR DIJKSTRA, donde PESOS SON TRIEMPO DE RECORRER CAMINO >0
    (me da fiaca cambiar)
    """
    
    aristas = obtener_aristas_disponibles(grilla, grafo)
    dist = dict()
    padres = dict()

    for v in grafo:
        dist[v] = float('inf')
    dist[origen] = 0
    padres[origen] = None
    
    for _ in range(grafo.cantidad() - 1):
        for arista in aristas:
            v, w, peso_vw = arista
            nueva_dist = dist[v] + peso_vw
            if nueva_dist < dist[w]:
                dist[w] = nueva_dist
                padres[w] = v
    
    return reconstruir_camino(padres, origen, destino)
            
    

def obtener_aristas_disponibles(grilla: dict, grafo: Grafo):
    aristas = []
    for v in grafo:
        if grilla[v] == "LIBRE":
            for w in grafo.adyacentes(v):
                if grilla[w] == "LIBRE":
                    aristas.append((v, w, grafo.peso(v, w)))
    return aristas



"""
EJERCICIO 34:

Daniel está a punto de casarse y tiene un problema: 
gastó casi todo su dinero en la luna de miel. 
Contrató un salón para la fiesta donde sólo hay 2 mesas (muy, muy grandes, pero 2 en fin). 
Debe repartir a los n invitados entre las dos mesas, y su esposo le indicó una condición: 
en cada mesa debe sentarse gente que se lleve bien entre todos ellos. 
Daniel cuenta con la información de quién se lleva bien con quién, 
y necesita poder determinar si hay alguna forma de separar en dos grupos de gente donde en cada grupo 
todos se lleven bien entre sí.

a. Modelar este problema utilizando grafos, indicando claramente qué son los vértices y qué las aristas.

b. Implementar un algoritmo que reciba un grafo como el modelado en el punto (a) 
y devuelva ambos grupos de personas. Indicar y justificar la complejidad del algoritmo implementado.

IMPORTANTE: tener en cuenta que resolver el problema de forma directa puede ser difícil.
Recomendamos plantearse el problema inverso: 
poder separar en dos grupos tal que en ningúno de los grupos haya un par que no se lleven bien.
"""

def organizar_mesas(grafo: Grafo):
    """
    **Modelacion**
    El grafo modelado es un grafo no dirigido:
    Vertice: personas invitadas
    aristas: Quien no se lleva bien con quien
    
    Entonces por ejemplo 
    P1 no se lleva bien con p2 => arista(p1, p2)
    
    Con este modelado cada conjunto bipartito seria una mesa.
    
    **PASOS**
    Utilizar un dfs para pintar cada conjunto
    
    **COmplejidad**
    Como es un dfs la complejidad es O(v +e) quie en este caso seria O(personas + relaciones malas)
    En el peor de los casos todos van en la mesa 1 si e es 0, osea si nadie se llva mal con nadie 
    y hay v conjuntos conexos
    """
    conj_1 = set()
    conj_2 = set()
    visitados = set()
    
    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            conj_1.add(v)
            if not dfs_hallar_conjuntos(v, grafo, conj_1, conj_2, visitados):
                return False
    return conj_1, conj_2

def dfs_hallar_conjuntos(v, grafo: Grafo, conj_1: set, conj_2: set, visitados: set):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            if v in conj_1:
                conj_2.add(w)
            else:
                conj_1.add(w)
            if not dfs_hallar_conjuntos(w, grafo, conj_1, conj_2, visitados):
                return False            
        # Caso hipotetico de que no se puede armar mesas
        else:
            if v in conj_1 and w in conj_1:
                return False
            elif v in conj_2 and w in conj_2:
                return False
    return True