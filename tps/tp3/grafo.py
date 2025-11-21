
class Grafo:
    def __init__(self, es_dirigido: bool):
        """
        Grafo(es_pesado, es_dirigido) toma dos parametros y devuelve un grafo dirigido y/o pesado
        """
        self.vertices = dict()      # Diccionario de listas
        self.pesos = dict()         # Diccionario de diccionario
        self.es_dirigido = es_dirigido
    
    
    def agregar_vertice(self, v):
        """
        agregar_vertice(v) toma un vertice 'v' y lo agrega al grafo.
        Si vertice ya existe devuelve error "Vertice ya esta en grafo"
        """
        if v in self.vertices:
            print("Vertice ya esta en grafo")
            return
        self.vertices[v] = []
        self.pesos[v] = dict()
    
        
    def agregar_arista(self, v, w, peso_vw: 0):
        """
        agregar_arista(v, w, peso_vw) toma dos vertices 'v' 'w' y los une.
        Por defecto el peso_vw es 0 para grafo no pesado.
        Si ya existe arista entonces devuelve "Ya existe arista".
        """
        if v not in self.vertices or w not in self.vertices:
            print("Algun vertice ingresado no existe")
            return
        self.vertices[v].append(w)
        self.pesos[v][w] = peso_vw
        if not self.es_dirigido:
            self.vertices[w].append(v)
            self.pesos[w][v] = peso_vw
        
        
    def adyacentes(self, v) -> list:
        """
        adyacentes(v) toma un vertice v y devuelve los adyacentes del vertice (aristas)
        """
        if v not in self.vertices:
            print("No existe vertice en grafo")
            return None
        return self.vertices[v]
    
    
    def peso_arista(self, v, w):
        """
        peso_arista(v, w) toma dos airistas y devuelve su peso si es un grafo pesado
        """
        if v not in self.vertices or w not in self.vertices:
            print("Algun vertice ingresado no existe")
            return
        return self.pesos[v][w]
    
        
    def obtener_vertices(self) -> list:
        """
        obtener_vertices() devuelve una lista con los vertices del grafo
        """
        return self.vertices.keys()
    