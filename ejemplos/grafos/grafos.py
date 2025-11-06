

class Grafo:
    def __init__(self, dirigido: bool):
        self.dirigido = dirigido
        self.vertices = dict()
        self.cantidad = 0
    
    # *** Metodos comunes ***
    
    def es_dirigido(self):
        """
        grafo.es_dirigido() -> bool
        Devuelve True si el grafo es dirigido, 
        False en caso contrario
        """
        return self.dirigido
    
    def obtener_cantidad(self) -> int:
        """
        grafo.cantidad() -> int
        Devuelve la cantidad de vertices que hay en el grafo
        """
        return self.cantidad
    
    # Métodos de vertice
    def agregar_vertice(self, vertice):
        """
        grafo.agregar_vertice(v)
        Toma un vertice `v` y lo agrega al grafo
        """
        # Crear una lista vacia de aristas para el vertice
        self.vertices[vertice] = []
        self.cantidad += 1
    
    
    def borrar_vertice(self, vertice):
        """
        grafo.borrar_vertice(v)
        Borra el vertice `v` del grafo. Borra todos los grados del vertice.
        """
        for v, aristas in self.vertices:
            if v != vertice:
                for arista, peso in aristas:
                    if arista == vertice:
                        aristas.remove((arista, peso))
        self.vertices.pop(vertice)
        self.cantidad -= 1
        
    # Aristas
    def agregar_arista(self, v, w, peso=0):
        """
        grafo.agregar_arista(v, w, peso)
        Toma dos vertices V, W, y agrega flecha desde v hasta W si es dirigido. 
        Peso es 0 por defecto.
        """
        self.vertices[v].append((w, peso))
        if not self.es_dirigido():
            self.vertices[w].append((v, peso))
    
    def borrar_arista(self, v, w):
        """
        grafo.borrar_arista(v, w)
        Borra arista que va de V a W. Si no es dirigido borra la otra relación tambien.
        """
        for arista, peso in self.vertices[v]:
            if arista == w:
                self.vertices[v].remove((arista, peso))
                break
        
        if not self.es_dirigido:
            for arista, peso in self.vertices[w]:
                if arista == v:
                    self.vertices[w].remove((arista, peso))
                    break
    
    def adyacentes(self, v) -> list:
        """
        grafo.adyacentes(v)
        Devuelve una lista de las aristas de V en formato [(w_1, peso_1)...(w_n, peso_n)]
        """
        return self.vertices[v]
    
    def obtener_vertices(self) -> list:
        """
        obtener_vertices
        devuelve una lista con todos los vertices del grafo
        """
        return list(self.vertices.keys())
    
    
    