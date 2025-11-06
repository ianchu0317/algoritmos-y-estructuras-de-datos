from collections import deque
import heapq

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


class Pila:
    def __init__(self):
        self.datos = deque()
        self.cantidad = 0
    
    def esta_vacia(self):
        return self.cantidad == 0
    
    def cantidad(self):
        return self.cantidad
    
    def apilar(self, elem):
        self.cantidad += 1
        self.datos.append(elem)
        
    def desapilar(self):
        self.cantidad -= 1
        return self.datos.pop()



class Heap:
    def __init__(self, max_heap=False):
        self.datos = []
        self.max_heap = max_heap
    
    def esta_vacia(self):
        return len(self.datos) == 0
    
    def cantidad(self):
        return len(self.datos)
    
    def apilar(self, elem):
        if not self.max_heap:
            heapq.heappush(self.datos, elem)
    
    def desapilar(self):
        if not self.max_heap:
            return heapq.heappop()
