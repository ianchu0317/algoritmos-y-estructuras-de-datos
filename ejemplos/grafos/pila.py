from collections import deque


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
    