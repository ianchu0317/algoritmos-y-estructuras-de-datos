

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
        return self.es_dirigido
    
    def cantidad(self) -> int:
        """
        grafo.cantidad() -> int
        Devuelve la cantidad de vertices que hay en el grafo
        """
        return self.cantidad
    
    # Métodos de vertice
    def agregar_vertice(self, vertice: str):
        """
        grafo.agregar_vertice(v)
        Toma un vertice `v` y lo agrega al grafo
        """
        # Crear una lista vacia de aristas para el vertice
        self.vertices[vertice] = []
    
    def borrar_vertice
    