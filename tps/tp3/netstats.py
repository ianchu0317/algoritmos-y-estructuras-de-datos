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
    
    
    def agregar_post(self, post: str, links: list):
        """
        Toma un post y sus links y los agrega a la red.
        
        **Ejemplo**
        post = "Argentina"
        links = ["Francia", "Portugal", "Espania"] 
        netstat.agregar_post(post, links)
        
        lista = ["Argentina", "Francia", "Portugal", "Espania"]
        netstat.agregar_post(lista[0], lista[1:])
        """
        # Agregar Publicacion a la red si no existe
        if not self.red.hay_vertice(post):
            self.red.agregar_vertice(post)
        # Agregar enlace a la red si no existe
        for link in links:
            if not self.red.hay_vertice(link):
                self.red.agregar_vertice(link)
            if not self.red.hay_arista(post, link):
                self.red.agregar_arista(post, link)
    
    
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
        resultado = f"{cadena_camino_minimo}\nCosto: {len(camino_minimo)-1}"
        
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
    paises1 = ["Argentina", "Francia", "Portugal", "España"]
    paises2 = ["Portugal", "China", "Colombia"]
    paises3 = ["Colombia", "Taiwan"]
    
    netstat.agregar_post(paises1[0], paises1[1:])
    netstat.agregar_post(paises2[0], paises2[1:])
    netstat.agregar_post(paises3[0], paises3[1:])
    
    # Test de caminos minimos
    print(netstat.camino("Argentina", "Taiwan"))   # Argentina -> Portugal -> Colombia -> Taiwan  (Costo: 3)
    print(netstat.camino("Portugal", "Taiwan"))    # Portugal -> COlombia -> Taiwan (Costo 2) 
    print(netstat.camino("Argentina", "España"))   # Argentina -> España (Costo 1)
    
    