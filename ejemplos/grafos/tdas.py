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



class UnionFind:
    def __init__(self, elems: list):
        self.grupos = dict()
        for e in elems:
            self.grupos[e] = e        
    
    def find(self, v):
        """
        UnionFInd.find(v) devuelve el grupo donde pertenece V
        """
        if self.grupos[v] == v:
            return v
        # Actualizar el valor y devolver valor del grupo
        self.grupos[v] = self.find(self.grupos[v])
        return self.grupos[v]
    
    def union(self, u, v):
        """
        UnionFind.union(u, v) une a U al grupo de V
        """
        nuevo_grupo = self.find(u)
        otro = self.find(v)
        self.grupos[otro] = nuevo_grupo
    