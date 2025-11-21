
class Grafo:
    def __init__(self, es_pesado: bool, es_dirigido: bool):
        """
        Grafo(es_pesado, es_dirigido) toma dos parametros y devuelve un grafo dirigido y/o pesado
        """
        self.vertices = dict()
        self.es_pesado = es_pesado
        self.es_dirigido = es_dirigido
    
    def agregar_vertice(self, v):
        """
        agregar_vertice(v) toma un vertice 'v' y lo agrega al grafo.
        Si vertice ya existe devuelve error "Vertice ya esta en grafo"
        """
        
    def agregar_arista(self, v, w):
        """
        agregar_arista(v, w) toma dos vertices 'v' 'w' y los une
        Si ya existe arista entonces devuelve "Ya existe arista"
        """
        
    def adyacentes(self, v) -> list:
        """
        adyacentes(v) toma un vertice v y devuelve los adyacentes del vertice (aristas)
        """
    
    def peso_arista(self, v, w) -> int:
        """
        peso_arista(v, w) toma dos airistas y devuelve su peso si es un grafo pesado
        """
        
    def obtener_vertices(self) -> list:
        """
        obtener_vertices() devuelve una lista con los vertices del grafo
        """