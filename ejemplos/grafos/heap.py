import heapq

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
