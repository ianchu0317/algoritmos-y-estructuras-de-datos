from grafos import Grafo
from tdas.cola import Cola
from tdas.heap import Heap
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


"""
Ejercicio 35:
Implementar un algoritmo que reciba un grafo no dirigido 
y determine la cantidad mínima de aristas que debería agregársele para que el grafo sea conexo.
Obviamente, si el grafo ya es conexo el algoritmo debe devolver 0. 
Indicar y justificar la complejidad del algoritmo implementado.
"""

def calcular_min_aristas_grafo_conexo(grafo: Grafo):
    """
    La idea es calcular la cantidad de componentes conexos en un grafo no dirigido.
    Y la cantidad de aristas a agregar es Componentes - 1.
    Es como el caso de un arbol E = V - 1.
    
    Complejidad O(v + e) utilizo un dfs
    """
    
    visitados = set()
    comp_conexo = 0
    for v in grafo:
        if v not in visitados:
            visitados.add(v)
            _dfs(v, grafo, visitados)
            comp_conexo += 1
    return comp_conexo - 1
            
def _dfs(v, grafo: Grafo, visitados: set):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            _dfs(w, grafo, visitados)


"""
Ejercicio 3, parcialito 1C2025:
Se tiene un grafo no dirigido sin pesos, que representa una huerta de zanahorias. Cada vértice representa a una zanahoria y
cada arista (v, w) indica que v y w se encuentran lo suficientemente cerca como para que si una se pudre, entonces podrá
“contagiar a la otra”. Luego, esta nueva zanahoria contagiada infectará a las otras que se encuentren cercanas. En una unidad
de tiempo, una zanahoria podrida infecta a todas las que tiene a su alcance. Suponiendo que existe una única componente
conexa en dicho grafo, implementar una función que reciba un grafo de dichas características y una zanahoria podrida inicial,
e indique cuál o cuáles son las últimas zanahorias en pudrirse. Indicar y justificar la complejidad de la función.
"""

def ej3_p3_1c2025(grafo: Grafo, origen):
    """
    La idea es hacer un bfs calculando la distancia de cada zanahoria a la infectada.
    Las ultimas en infectarse son las que se encuentran en mayor distancia.
    
    Osea utilizar bfs para calcular orden y luego guardar orden maximo. O(v + e)
    Recorrer lo guardado para ver vertices que se encuentran en mayor distancia. O(v)
    """
    
    cola = Cola()
    visitados = set()
    orden = dict()
    
    orden[origen] = 0
    max_dist = 0
    visitados.add(origen)
    cola.encolar(origen)
    
    while not cola.esta_vacia():
        v = cola.desencolar()
        
        for w, _ in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                cola.encolar(w)
                if max_dist < orden[w]:
                    max_dist = orden[w]
    
    vertices = []            
    for v, dist in orden.items():
        if dist == max_dist:
            vertices.append(v)
    return vertices



"""
Ejercicio 2, parcialito 3 1c2025:

En una pequeña ciudad costera, todos los habitantes usan el AlgoMaps, una aplicación que les ayuda a decidir sus rutas de
viaje diarias. Conociendo su funcionamiento, sabemos que siempre ofrece la ruta del camino más corto de un punto a otro
dadas las rutas existentes, sabiendo que cuenta con la información de la distancia cubierta por cada camino entre punto y
punto. En el futuro se planifican ciertas obras en diferentes puntos de la ciudad que seguramente afecten a la red de transporte.
Para evitar afectar en demasía a los ciudadanos, se busca analizar la información disponible para planificar las obras. Sería
deseable evitar bloquear el transporte de los puntos más importantes de la ciudad.
a. Explicar cómo la información disponible se puede modelar con un grafo, y explicar sus características. 
¿Qué información debemos solicitarle a AlgoMaps?
b. Desarrollar una función puntosImportantes(grafo, k) que reciba el grafo definido en el punto (a) y que devuelva los k
puntos más importantes en la ciudad, según lo modelado. Indicar y justificar su complejidad.
"""


def ej2_p3_1c2025(grafo: Grafo, k: int) -> list:
    """
    **Modelado**
    Se puede modelar al problema con un grafo dirigido y pesado donde:
    - Vertices: los distintos puntos de la ciudad
    - Aristas: Los distintos caminos de la ciudad
    - Peso de aristas: Tiempo de recorrido de un camino
    
    Teniendo en cuenta eso, los puntos importantes son los puntos que con mas frecuencia están en los caminos minimos,
    es decir los k-puntos con mayor centralidad del grafo 
    
    **Procedimiento**
    Obtener los puntos más transitados / importantes (calcular centralidad)
    Ordenar los puntos obtenidos de menor a mayor
    Obtener los k puntos mas importantes
    
    -> Se cuenta con distancia 
    
    **Complejiad**
    O(v²*(elog(v) + v) + klog(k))
    """
    centralidad = dict()
    for v in grafo:
        centralidad[v] = 0
    
    for v in grafo:
        for w in grafo:
            if w == v:
                continue
            _, padres = dijkstra(grafo, v, w)
            
            if not padres:
                continue
            actual = padres[w]
            while actual != v:
                centralidad[actual] += 1
                actual = padres[actual]
    
    heap = Heap()
    i = 0
    for v, freq in centralidad.items():
        if i < k:
            heap.encolar((freq, v))
            i += 1
        elif heap.ver_primero()[0] < freq:
            heap.desencolar()
            heap.encolar((freq, v))
    resultado = []
    while not heap.esta_vacia():
        _, v = heap.desencolar()
        resultado.append(v)
    return resultado



"""
Ejercicio 1 parcialito 3 r1, 1c2025:
Resolvé los siguientes problemas en forma clara y legible. Podés incluir tantas funciones auxiliares como creas necesarias. Los
algoritmos a implementar deben ser de la mejor complejidad posible dadas las características del problema.
1. En nuestra huerta de zanahorias se está construyendo un sistema de riego. La plantación completa es muy extensa y entre
algunas zonas no pueden ubicarse mangueras, por lo que no es tarea sencilla determinar cómo conectar las mangueras para
hacer llegar el agua a todas las zonas.
Se tiene un grafo pesado en el que se representan como vértices estas zonas, y las aristas determinan la distancia entre zona y
zona (en caso de poder conectarse entre sí). Definir una función que, dado el grafo mencionado, devuelva la mejor forma de
conectar las distintas zonas, se requiere:
minimizar la cantidad de metros de manguera a utilizar.
que ninguna zona quede desconectada del resto de la red.
"""

def ej1_p3_r1_1c2025(grafo: Grafo):
    """
    *Modelacion*
    Para este ejercicio se modela un grafo pesado no dirigido, donde:
    - Vertices -> zonas con plantacion
    - Aristas -> mangueras para conectar con agua distintas zonas
    - Peso aristas -> distancias entre distintas zonas
    
    El objetivo es encontrar como conectar zonas con menor cantidad de mangueras posibles.
    Osea conectar todos los vertices, pero con menor cantidad de distancias posibles (mst)
    
    En este caso utilizo prim para no tener que implementar ordenamientos
    
    *Complejidad*
    - Crear mst y cargar los vertices O(v)
    - Crear un heap O(1)
    - Cargar el arista del vertice random al heap: O(elog(e)) -> en el peor de los casos tiene todos los aristas del grafo
    - Se itera para cada arista del grafo: E aplicando operaciones de encolar y desencolar en el heap O(e * log(e)) 
    tambien se aplican operaciones O(1) pero se acotan
    
    Entonces complejidad final O(e log e) -> O(e log(v²)) -> O(2elog(v)) -> O(e log (v))
    """
    mst = Grafo(False)
    heap = Heap()
    visitados = set()
    
    for v in grafo:
        mst.agregar_vertice(v)
    
    v = random.choice(grafo.obtener_vertices())
    visitados.add(v)
    for w, peso_vw in grafo.adyacentes(v):
        heap.encolar((v, w, peso_vw))
    
    while not heap.esta_vacia():
        v, w, peso_vw = heap.desencolar()
        
        if w in visitados:
            continue
        visitados.add(w)
        mst.agregar_arista(v, w, peso_vw)
        
        for x, peso_wx in grafo.adyacentes(w):
            if x not in visitados:
                heap.encolar((w, x, peso_wx))
    return mst

"""
Ejercicio 2 parcialito 3 r1, 1c2025:
Un ayudante de esta materia quiere comenzar a leer los libros de Sanderson. Este autor tiene distintas sagas de libros
interconectadas en un mismo universo, sumado a distintos libros independientes. Una opción podría ser leerlos orden de salida
cronológica, pero hay algo más importante, la consideración del autor. En la web de Sanderson se puede encontrar, por cada
libro, un listado de cuáles son los libros que es necesario leer con anterioridad.
a. Describir detalladamente cómo se podría representar esta información en un grafo que nos sea útil para resolver este
problema.
b. Implementar una función que, dado un grafo definido como en el punto a, devuelva un orden posible que nuestro ayudante
pueda seguir. Indicar y justificar la complejidad de la función.
"""

def ej2_p3_r1_1c2025(grafo: Grafo):
    """
    *Modelación*
    Podemos modelar el ejercicio con un grafo dirigido no pesado, en donde:
    - Vertice: los distintos libros de Sanderson
    - Aristas: libro a leer con anterioridad.
    Entonces así por ejemplo:
    L1 se lee antes que L2: L1 -> L2
    
    La idea con esto es del listado de información, crear un grafo y realizar orden topológico en ese grafo
    
    *Complejidad*
    - Crear cola y estructuras auxiliares O(1)
    - Hallar grado de entrada de cada vertice O(v + e) ya que visito cada vertice y sus adyacentes
    - Encolar vertices de grado 0: O(v) en el peor de los casos son todos los vertices
    - Ir desencolando todos los vertices y encolar sus adyacentes O(v + e). 
        Adentro se realizan asignacion de variables, comparaciones O(1)
    - Devolver orden O(1)
    => Complejidad total O(v + e) que en este caso es O(libros + Informacion)
    Para cada libro: l1...Ln hay n-vertices
    Para cada informacion: L1 leer antes que l2 hay un arista
    V = Libro
    E = Informacion
    """
    
    cola = Cola()
    orden = []
    g_entrada = grados_entrada(grafo)
    
    for v in grafo.obtener_vertices():
        if g_entrada[v] == 0:
            cola.encolar(v)
            
    while not cola.esta_vacia():
        v = cola.desencolar()
        orden.append(v)       
        for w, _ in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.encolar(w)
    return orden
    
    
def grados_entrada(grafo: Grafo):
    g_entrada = dict()
    
    for v in grafo.obtener_vertices():
        g_entrada[v] = 0
    
    for v in grafo.obtener_vertices():
        for w, _ in grafo.adyacentes(v):
            g_entrada[w] += 1
    return g_entrada


"""
Ejercicio 1, p3 r2 1c2025:
En nuestra huerta tenemos unos rociadores de insecticidas automáticos. Cada rociador cuenta con la dosis apropiada para
cubrir hasta un máximo de 5 plantaciones a su alrededor. Es necesario averiguar si algún rociador tiene más de 5 plantaciones
alrededor ya que de tener una mayor cantidad la dosis sería insuficiente. Se tiene un grafo en donde los vértices son los
rociadores y plantas, es no pesado y dirigido (el origen de una arista es el rociador y el destino es una planta en su rango).
Implementar una función que reciba este grafo y devuelva, en caso que un rociador tenga más de 5 plantaciones alrededor,
el conjunto de plantaciones alrededor de dicho rociador (si hay más de un rociador que cumpla esta condición, devolver la
información correspondiente a cualquiera de estos). En caso contrario, devolver None. Indicar y justificar la complejidad de la
función.
"""

def ej1_p3_r2_1c2025(grafo: Grafo):
    """
    *Modelado*
    Grafo dirigido no pesado donde:
    - Vertices: Plantas / roceadores
    - Aristas: distancia de roceado
    
    La idea es encontrar para cada vértice, si hay algun vértice (roceador) que tenga grados de salida mayor a 5.
    Si hay algun vértice (roceador) que cumpla condicion; devolver las plantas (los adyacentes al roceador).
    
    Para eso hago un dfs para cada vertice no visitado, si hay algun vertice que cumple eso devolver sus ady.
    
    **Complejiad**
    Estoy hacienod un dfs donde visito todos los vertices y los aristas exactamente una vez, mientras voy realizando
    operaciones como condicionales y asignaciones de variables O(1).
    
    O(v + e) complejidad final -> O(roceadores + plantas)
    """
    
    for v in grafo.obtener_vertices():
        plantaciones = grafo.adyacentes(v)
        if len(plantaciones) > 5:
            return plantaciones
    return None
