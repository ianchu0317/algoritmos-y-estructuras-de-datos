#!/usr/bin/python3

import sys
import netstats 
from netstats import Netstats

# Configurar maxima recursion
MAX_RECURSION = 100000
sys.setrecursionlimit(MAX_RECURSION)


def operar(wiki: Netstats, operacion, argumentos):
    """
    Toma un servidor tipo Netstat, una operacion y los argumentos para la operacion.
    Ejecuta e imprime por pantalla el resultado.
    """
    
    if operacion == netstats.OPERACION_LISTAR_OPERACIONES:
        salida = wiki.listar_operaciones()
    
    elif operacion == netstats.OPERACION_CAMINO:
        origen, destino = argumentos[0].split(",")
        salida = wiki.camino(origen, destino)
    
    elif operacion == netstats.OPERACION_MAS_IMPORTANTES:
        n = int(argumentos[0])
        salida = wiki.mas_importantes(n)
    
    elif operacion == netstats.OPERACION_CONECTADOS:
        pagina = argumentos[0]
        salida = wiki.conectados(pagina)
    
    elif operacion == netstats.OPERACION_CICLO:
        pagina, n = argumentos[0].split(",")
        salida = wiki.ciclo(pagina, int(n))
    
    elif operacion == netstats.OPERACION_LECTURA:
        paginas = argumentos[0].split(",")
        salida = wiki.lectura(paginas)
    
    elif operacion == netstats.OPERACION_DIAMETRO:
        salida = wiki.diametro()
        
    elif operacion == netstats.OPERACION_RANGO:
        pagina, n = argumentos[0].split(",")
        salida = wiki.rango(pagina, int(n))
    
    elif operacion == netstats.OPERACION_COMUNIDAD:
        pagina = argumentos[0]
        salida = wiki.comunidad(pagina)
    
    elif operacion == netstats.OPERACION_NAVEGACION:
        pagina = argumentos[0]
        salida = wiki.navegacion(pagina)

    elif operacion == netstats.OPERACION_CLUSTERING:
        if not argumentos:
            salida = wiki.clustering()
        else:
            pagina = argumentos[0]
            salida = wiki.clustering(pagina)    
    else:
        salida = "No existe operacion"
    
    print(salida)



if __name__ == '__main__':
    wiki = Netstats()
    
    # Leer TSV opcional
    if len(sys.argv) > 1:
        archivo_tsv = sys.argv[1]  # Suponer que se pasa bien argumento
        # Leer tsv
        with open(archivo_tsv, "r") as archivo:
            contenido = archivo.read()
        
        # Cargar servidor -> Para cada linea de tsv, separar a lista
        for dato in contenido.split("\n"):
            linea_lista_tsv = dato.split("\t")
            wiki.agregar_post(linea_lista_tsv[0], linea_lista_tsv[1:])
        
    # Esperar comandos -> Hasta que no haya inputs entonces break
    while True:
        try:
            # Leer argumento
            linea_operacion = input()
            operacion = linea_operacion.split(" ", 1)   # Split solo el primer espacio
            # Operar argumento
            operar(wiki, operacion[0], operacion[1:])
        except EOFError:
            break

