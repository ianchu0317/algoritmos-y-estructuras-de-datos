
class Grafo:
    def __init__(self, es_dirigido: bool, vertices=[]):
        """
        Grafo(es_dirigido) toma dos parametros y devuelve un grafo dirigido.
        
        Ejemplo:
            Grafo(es_dirigido=False) -> Crea un grafo no dirigido
            Grafo(es_dirigido=True) -> Crea un grafo dirigido
        """
        # Estructura interna
        self.vertices = dict()      # Diccionario de listas
        self.pesos = dict()         # Diccionario de diccionario
        self.es_dirigido = es_dirigido
        
        # Cargar datos iniciales
        for v in vertices:
            self.vertices[v] = []
            self.pesos[v] = dict()
    
    
    def __iter__(self):
        """
        Inicializar iterador para vértices del grafo.
        
        Sintaxis:
            for v in grafo:
                ...
        """
        return iter(self.vertices)
    
    
    def hay_vertice(self, v) -> bool:
        """
        Chequea si vertice 'v' está en el grafo.
        Devuelve True si está en grafo, False en caso contrario.
        """
        return v in self.vertices
    
    
    def hay_arista(self, v, w) -> bool:    
        """
        Toma dos vértices y chequea si hay arista entre v y w, donde v es origen y w destino en grafo dirigido.
        Devuelve True si cumple, False en caso contrario.
        Si alguno de los vertices no existe devuelve error con "No existe vertice en grafo"
        """
        if not self.hay_vertice(v) or not self.hay_vertice(w):
            raise ValueError("No existe vertice en grafo")
        return w in self.vertices[v]
    
    
    def agregar_vertice(self, v):
        """
        Toma un vertice 'v' y lo agrega al grafo.
        Si vertice ya existe devuelve error "Vertice ya esta en grafo"
        """
        if self.hay_vertice(v):
            raise ValueError("Vertice ya esta en grafo")
        self.vertices[v] = []
        self.pesos[v] = dict()
    
        
    def agregar_arista(self, v, w, peso_vw=0):
        """
        Toma dos vertices 'v' 'w' y los une.
        Si se quiere operar con grafo no pesado, no pasarle nada al parametro peso_vw (por defecto 0).
        Si algun vertice no existe devuelve "Algun vertice no existe"
        Si ya existe arista entonces devuelve "Ya existe arista".
        """
        if not self.hay_vertice(v) or not self.hay_vertice(w):
            raise ValueError("Algun vertice no existe")
        
        if w in self.vertices[v]:
            raise ValueError("Ya existe arista")

        self.vertices[v].append(w)
        self.pesos[v][w] = peso_vw
        if not self.es_dirigido:
            self.vertices[w].append(v)
            self.pesos[w][v] = peso_vw
        
        
    def adyacentes(self, v) -> list:
        """
        Toma un vertice v y devuelve los adyacentes del vertice (aristas).
        Si no existe el vertice devuelve "No existe vertice en grafo".
        """
        if v not in self.vertices:
            raise ValueError("No existe vertice en grafo")
        return self.vertices[v]
    
    
    def peso_arista(self, v, w):
        """
        Toma dos vértices y devuelve el peso de su arista si existe. 
        Si algun vertice ingresado no existe devuelve "Algun vertice no existe".
        Si no hay arista entonces devuelve "No existe arista en grafo"
        """
        if not self.hay_vertice(v) or not self.hay_vertice(w):
            raise ValueError("Algun vertice no existe")
        if not self.hay_arista(v, w):
            raise ValueError("No existe arista en grafo")
        return self.pesos[v][w]
    
        
    def obtener_vertices(self) -> list:
        """
        Devuelve una lista con los vertices del grafo
        """
        return list(self.vertices.keys())

    
    def borrar_vertice(self, v):
        """
        Borra un vertice y todas sus aristas del grafo. 
        Si es dirigido también borra sus apariciones en otros vertices.
        Si no existe vertice en grafo, devuelve "No existe vertice en grafo"
        """
        
        if not self.hay_vertice(v):
            raise ValueError("No existe vertice en grafo")
        
        self.vertices.pop(v)
        for _, ady in self.vertices.items():
            if v in ady:
                ady.remove(v)
        # Eliminar los pesos
        self.pesos.pop(v)
        for w, peso_wv in self.pesos.items():
            if v in peso_wv:
                peso_wv.pop(v)
    
    
    def borrar_arista(self, v, w):
        """
        Toma dos vertices del grafo y SI tienen aristas borrar la arista, sino nada.
        Si algun vertice ingresado no esta en grafo, deveuvle "Algun vertice no existe"
        Si arista no existe, devuelve: no existe arista en grafo
        """
        if not self.hay_vertice(v) or not self.hay_vertice(w):
            raise ValueError("Algun vertice no existe")
        if not self.hay_arista(v, w):
            raise ValueError("No existe arista en grafo")

        if w in self.vertices[v]:
            self.vertices[v].remove(w)
            self.pesos[v].pop(w)
        
        if not self.es_dirigido:
            if v in self.vertices[w]:
                self.vertices[w].remove(v)
                self.pesos[w].pop(v)
