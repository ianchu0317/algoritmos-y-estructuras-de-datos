
"""
Resolvé los siguientes problemas en forma clara y legible. Podés incluir tantas funciones auxiliares como creas necesarias. Los
algoritmos a implementar deben ser de la mejor complejidad posible dadas las características del problema.
1. Se quiere modelar un sitio que conecte ONGs con personas y se cuenta con la siguiente información:
Para cada persona a qué ONGs donó.
Para cada persona a qué otras personas sigue.
Se considera que si una persona sigue a otra la distancia entre ellas es 1; y que si dos personas donaron a la misma ONG la
distancia entre ellas también es 1.
a) Modelar un grafo que represente estos datos
b) Dadas dos personas y el grafo modelado en el punto anterior, encontrar el camino más corto que las conecta.
"""
def ej1_3p_2c2025(grafo, origen, destino):
    """
    **Modelacion**
    Se modela un grafo dirigido pesado donde:
        - Vertices son personas
        - Aristas son los seguimientos y distancias entre personas
        - Peso de arista: son las distancias entre personas
    
    -> Si p1 dono a ONG1 y p2 dono a ONG1 entonces p1-> p2 y p2 -> p1
    -> Si p1 sigue a p2 entonces p1 -> p2
    
    La idea es encontrar un camino minimo dado p1: origen y p2: destino
    La trampa en este ejercicio es aunque es un grafo pesado, los pesos son los mismos, así que no tiene sentido usar dijkstra
    Utilizo BFS Complejidad O(V + E) -> O(Personas + Relaciones)
    """
    # O(1) Crear estructuras auxiliares
    cola = Cola()
    visitados = set()
    padres = dict()
    
    cola.encolar(origen)
    visitados.add(origen)
    padres[origen] = None
    
    # O(V + E) Visito cada vertice y arista buscando el menor camino
    while not cola.esta_vacia():
        v = cola.desencolar()
        
        if v == destino:
            return reconstruir_camino(padres, origen, destino)
        
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.encolar(w)
    
    # Si llegó aca no hay camino 
    return []

def reconstruir_camino(padres: dict, origen, destino):
    """
    O(V) en el peor caso
    """
    camino = [destino]
    while destino != origen:
        destino = padres[destino]
        camino.append(destino)
    camino.reverse()
    return camino

"""
2. Juan, un fanático de The Red Hot Chili Peppers, y un gran ayudante de diversas materias, está obligando a sus corregidos
a instruirse en el mundo de la música y les insiste que escuche a los Red Hot. Los alumnos después de un par de reuniones
escuchandolo quejarse de lo poco que saben de música, un poco cansados de esquivar la propuesta, finalmente la aceptan. Juan
antes de decirles qué álbumes escuchar, les aclara que no sean nabos, y que hay ciertos albumes que deben escuchar antes que
otros: imaginate que alguien escuche The Getaway antes que By the Way! Esta información la brinda con el siguiente formato:
para cada álbum menciona una lista de álbumes que hay que escuchar antes que ese.
a) Modelar este problema con grafos, indicando claramente qué son los vértices y qué son las aristas.
b) Implementar un algoritmo que le de a los alumnos un orden posible para escuchar los albumes que les recomendó Juan.
"""


"""
3. Sobre el grafo del dorso:
a. Realizar el seguimiento del algoritmo de Prim, arrancando desde el vértice A.
b) Encontrar, si es posible, una arista en este grafo al que se le pueda asignar un valor más bajo, pero aún así sea imposible
que esa arista sea parte de un árbol de tendido mínimo válido del grafo. En ese caso, indicar cuál es dicha arista y qué
valor se le asignaría para cumplir la condición.
"""