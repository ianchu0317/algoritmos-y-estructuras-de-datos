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
        ```
        # Utilizando variable/lista
        post = "Argentina"
        links = ["Francia", "Portugal", "Espania"] 
        netstat.agregar_post(post, links)
        # Utilizando sublistas 
        lista = ["Argentina", "Francia", "Portugal", "Espania"]
        netstat.agregar_post(lista[0], lista[1:])
        ```
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
    
    
    def conectados(self, pagina):
        """
        Toma una pagina y muestra todas las paginas a los que podemos llegar desde pagina 
        y que tambien pueden volver a pagina.
        
        **Complejidad**: O(P + L)
        """
        cfcs = bib.obtener_cfcs(self.red)
        for cfc in cfcs:
            if pagina in cfc:
                return ", ".join(cfc)    # O(P)
        return ""   # Si no existe pagina en la red
    
    
    def ciclo():
        pass
    
    
    def lectura(self, paginas: list):
        """
        Permite obtener un orden en el que es válido leer las páginas indicados en lista de paginas pasados por parametro. 
        Solo se tiene en cuenta los artículos mencionados en los parámetros.

        Si pagina_i tiene un link a pagina_j, entonces se lee primero pagina_j.
        Si existe orden, entonces devuelve una cadena separada por coma de un orden.
        Si no se puede leer las paginas en orden, devuelve "NO existe forma de leer las paginas en orden"
    
        **Parametros**: lista([Pagina_i, Pagina_j ..... Pagina_n])
        **Complejidad**: O(n + Ln)
        """
        orden = bib.orden_topologico_vertices(self.red, paginas)
        # Si orden obtenido no es igual, hubo ciclo
        if len(orden) != len(paginas):
            return "No existe forma de leer las paginas en orden"
        # Devolver orden topologico al revess
        return ", ".join(orden[::-1])

    
    def diametro(self) -> str:
        """
        Permite obtener el diametro de toda la red.
        
        **Parametros**: ninguno
        
        **Complejidad**: O(P (P + L))
        
        **Ejemplo**
        ```
        netstat.diametro()
        "Argentina" -> "Portugal" -> "Colombia" -> "Taiwan"
        costo: 4
        ```
        """
        camino_diametro = bib.diametro(self.red)
        cadena_diametro = " -> ".join(camino_diametro)   # O(P) en el peor de los casos
        resultado = f"{cadena_diametro}\nCosto: {len(camino_diametro)-1}"
        return resultado
        
    
    def rango(self, pagina: str, n: int) -> int:
        """
        Permite obtener la cantidad de páginas que se encuentren 
        excactamente a n links desde la pagina recibida por parametro.

        **Parametros**: pagina y n
        
        **Complejidad**: O(P + L)
        
        **Ejemplo**
        netstat.rango(p1, 2) -> 4 'hay 4 en ese rango'
        """
        return bib.en_rango(self.red, pagina, n)
    
    def comunidad():
        pass
    
    
    def navegacion(self, pagina) -> str:
        """
        Toma una pagina de origen y devuelve el camino con las paginas accediendo al primer link.
        Se detiene cuando no hay mas links o hasta haber llegado a 20 paginas.
        
        **Complejidad**: O(n)
        """
        nav_primer_link = bib.camino_primer_adyacente(self.red, pagina)
        return " -> ".join(nav_primer_link)
    
    
    def clustering():
        pass

if __name__ == '__main__':
    netstat = Netstats()
    paises1 = ["Argentina", "Francia", "Portugal", "España"]
    paises2 = ["Portugal", "China", "Colombia"]
    paises3 = ["Colombia", "Taiwan"]
    paises4 = ["Taiwan", "Portugal"]
    
    netstat.agregar_post(paises1[0], paises1[1:])
    netstat.agregar_post(paises2[0], paises2[1:])
    netstat.agregar_post(paises3[0], paises3[1:])
    netstat.agregar_post(paises4[0], paises4[1:])
    
    # Test de caminos minimos
    print(netstat.camino("Argentina", "Taiwan"))   # Argentina -> Portugal -> Colombia -> Taiwan  (Costo: 3)
    print(netstat.camino("Portugal", "Taiwan"))    # Portugal -> COlombia -> Taiwan (Costo 2) 
    print(netstat.camino("Argentina", "España"))   # Argentina -> España (Costo 1)
    print(netstat.camino("Argentina", " "))        # No se encontro camino
    
    # Test de diametros
    print(netstat.diametro())   # Argentina -> Portugal -> Colombia -> Taiwan (costo: 4)
    
    # Test de rango
    #print("Rango: ", netstat.rango("Argentina", 1))
    
    # Test de nav_primer link
    print(netstat.navegacion("Argentina"))
    
    # Test conectividad (cfcs)
    print(netstat.conectados("Colombia"))   # Taiwan -> Colombia -> portugal
    
    # Test lectura (orden topologico)
    print(netstat.lectura(["Portugal", "Argentina", "China"]))  # China -> Portugal -> Argentina