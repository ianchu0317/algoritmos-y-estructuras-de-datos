from grafos import Grafo



def orden_bfs(grafo: Grafo):
    """
    orden_bfs toma un grafo y devuelve su recorrido de padres en orden topologico
    """
def crear_grafo_ejemplo() -> Grafo:
    grafo = Grafo(True)
    # Variables (valor a poner) de vertices
    analisis_ii = "Análisis Matemático"
    fundamentos = "Fundamentos de Programacion"
    ids = "Introduccion a desarrollo de Software"
    algebra = "Algebra Lineal"
    orga = "Organizacion del computador"
    algo_ii = "Algoritmos y Estructuras de datos"
    paradigmas = "Paradigmas de programacion"
    proba = "Probabilidad y estadística"
    sisops = "Sistemas Operativos"
    teoria_algo = "Teoría de Algoritmos"
    materias = [analisis_ii, fundamentos, ids, algebra, orga, algo_ii, paradigmas, proba, sisops, teoria_algo]
    # Insertar vertices de grafo
    for v in materias:
        grafo.agregar_vertice(v)
    # Insertar Aristas (relaciones / correlativas)
    grafo.agregar_arista(analisis_ii, proba)
    grafo.agregar_arista(fundamentos, orga)
    grafo.agregar_arista(fundamentos, algo_ii)
    grafo.agregar_arista(ids, paradigmas)
    grafo.agregar_arista(ids, teoria_algo)
    grafo.agregar_arista(algebra, proba)
    grafo.agregar_arista(orga, sisops)
    grafo.agregar_arista(algo_ii, paradigmas)
    grafo.agregar_arista(algo_ii, teoria_algo)
    return grafo

if __name__ == '__main__':
    grafo_materias = crear_grafo_ejemplo()
    print(grafo_materias.obtener_vertices())