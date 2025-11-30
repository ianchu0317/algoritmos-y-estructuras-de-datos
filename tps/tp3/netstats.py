#!/usr/bin/python3
from grafo import Grafo
import biblioteca as bib

# Constantes
OPERACION_LISTAR_OPERACIONES = 'listar_operaciones'
OPERACION_CAMINO = 'camino'
OPERACION_MAS_IMPORTANTES = 'mas_importantes'
OPERACION_CONECTADOS = 'conectados'
OPERACION_CICLO = 'ciclo'
OPERACION_LECTURA = 'lectura'
OPERACION_DIAMETRO = 'diametro'
OPERACION_RANGO = 'rango'
OPERACION_COMUNIDAD = 'comunidad'
OPERACION_NAVEGACION = 'navegacion'
OPERACION_CLUSTERING = 'clustering'

LISTA_OPERACIONES=f"""{OPERACION_CAMINO}
{OPERACION_MAS_IMPORTANTES}
{OPERACION_CONECTADOS}
{OPERACION_CICLO}
{OPERACION_LECTURA}
{OPERACION_DIAMETRO}
{OPERACION_RANGO}
{OPERACION_COMUNIDAD}
{OPERACION_NAVEGACION}
{OPERACION_CLUSTERING}"""


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
        self.page_rank = None
    
    
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
        return LISTA_OPERACIONES
    
    
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
    
    
    def mas_importantes(self, n):
        """
        Devuelve las n-paginas más importantes de la red.
        En caso de que n < grafo, devuelve "N es mayor a la cantidad de paginas".
        
        **Complejidad**: O(l(P + L) + O(V log n))
        (falta optimizar para repetidas operaciones en una misma red)
        """
        if not self.page_rank:
            self.page_rank = bib.page_rank(self.red, 15, 0.85)
        return ", ".join(bib.obtener_top_n(self.page_rank, n))
    
    
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
    
    
    def ciclo(self, pagina, n: int):
        """
        Toma una pagina y un valor entero n y devuelve si hay ciclo en esa pagina de largo n.
        Si no encuentra entonces devuelve: "No se encontro recorrido"
        **Complejidad**: O(P^n)
        """
        recorrido = bib.ciclo_largo_n(self.red, pagina, n)
        if not recorrido:
            return "No se encontro recorrido"
        return " -> ".join(recorrido)
    
    
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
    
    
    def comunidad(self, pagina):
        """
        Devuelve la comunidad a la que pertenezca la ṕágina por parámetro.
        
        En caso de encontrar comunidad devuelve una lista de comunidades separados por coma.
        En caso contrario, "No se encontro comunidad"
        
        **Ejemplo**
        netstat.comunidad("Chile")
        **Complejidad**: O(k(V + E)) -> O(V + E) ya que k=10 en este caso
        """    
        comunidades = bib.comunidades(self.red)
        for comunidad in comunidades:
            if pagina in comunidad:
                return ", ".join(comunidad)
        return "No se encontro comunidad"
    
    
    def navegacion(self, pagina) -> str:
        """
        Toma una pagina de origen y devuelve el camino con las paginas accediendo al primer link.
        Se detiene cuando no hay mas links o hasta haber llegado a 20 paginas.
        
        **Complejidad**: O(n)
        """
        nav_primer_link = bib.camino_primer_adyacente(self.red, pagina)
        return " -> ".join(nav_primer_link)
    
    
    def clustering(self, pagina=None):
        """
        Toma una pagina (opcional) y devuelve su coeficiente de clustering.
        En caso de no indicar ninguna pagina, devuelve coeficiente de clustering de toda la red (promedio)
        
        **Parametros**
        - pagina (opcional)
        """
        if not pagina:
            return bib.clustering(self.red)
        return bib.clustering_vertice(self.red, pagina)
        

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
    
    # Test page rank
    grafo_pr = Grafo(True, vertices=["p1", "p2", "p3", "p4", "p5"])
    grafo_pr.agregar_arista("p1", "p4")
    grafo_pr.agregar_arista("p2", "p1")
    grafo_pr.agregar_arista("p2", "p4")
    grafo_pr.agregar_arista("p3", "p2")
    grafo_pr.agregar_arista("p4", "p3")
    grafo_pr.agregar_arista("p4", "p5")
    grafo_pr.agregar_arista("p5", "p4")
    
    print(bib.page_rank(grafo_pr, 10, 0.85, 6))
    
    # Test netstat mas_importantes (page_rank)
    print(netstat.mas_importantes(5))
    
    # Test ciclo de largo n
    print(bib.ciclo_largo_n(grafo_pr, "p2", 3))
    
    # Encontrar recorrido con netstat
    print(netstat.ciclo("Portugal", 3))     # Portugal -> Colomb -> Taiwan -> Portugal