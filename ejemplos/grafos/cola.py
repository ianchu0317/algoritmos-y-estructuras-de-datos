from collections import deque

class Cola:
    def __init__(self):
        self.datos = deque()
        self.cantidad = 0
    
    def esta_vacia(self):
        """
        esta_vacia devuelve True si la cola esta vacia, False en caso contrario
        """
        return self.cantidad == 0
    
    def encolar(self, elem):
        """
        encolar(e) toma un elemento y lo agrega a la cola
        """
        self.datos.append(elem)
        self.cantidad += 1
        
    def desencolar(self):
        """
        desencolar() -> devuelve el primer elemento de la cola
        """
        self.cantidad -= 1
        return self.datos.popleft()
        
    def cantidad(self):
        return self.cantidad
    