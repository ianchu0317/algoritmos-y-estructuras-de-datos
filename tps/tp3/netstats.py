#!/usr/bin/python3
from grafo import Grafo
import biblioteca as bib

class Netstats:
    """
    Netstats modela una red con grafos, donde:
    - Cada vertice son POSTS
    - Cada arista son links a otros Post desde un Post
    
    **Ejemplo**
    - P1 tiene link a [P2, P3... Pn]
    Entonces P1 -> P2
             P1 -> P3...
    """
    
    def __init__(self):
        self.red = Grafo(True)
        self.operaciones = [
            'camino', 'mas_importantes', 'conectados'
            'ciclo', 'lectura', 'diametro', 'rango', 
            'comunidad', 'navegacion', 'clustering'
        ]
    
    
    def listar_operaciones(self):
        """
        Devuelve una lista de operaciones disponibles
        
        Complejidad O(1)
        """
        return self.operaciones
    
    
    def camino(self, origen, destino) -> str:
        """
        Toma un origen y un destino.
        Devuelve una cadena con un camino en formato origen->destino.
        En caso de que no hay recorrido devuelve 'No se encontro recorrido'
        """
        camino_minimo = bib.camino_minimo(self.red, origen, destino)
        if not camino_minimo:
            return "No se encontro recorrido"
        
        cadena_camino_minimo = " -> ".join(camino_minimo)   # O(P) en el peor de los casos
        resultado = f"{cadena_camino_minimo}\nCosto: {len(camino_minimo)}"
        
        return resultado
    
    
    def mas_importantes():
        pass
    
    def conectados():
        pass
    
    def ciclo():
        pass
    
    def lectura():
        pass
    
    def diametro():
        pass
    
    def rango():
        pass
    
    def comunidad():
        pass
    
    def navegacion():
        pass
    
    def clustering():
        pass

if __name__ == '__main__':
    netstat = Netstats()
    print(netstat.listar_operaciones())