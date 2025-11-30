#!/usr/bin/python3

import sys
from netstats import Netstats

sys.setrecursionlimit(100000)

def operar(netstat: Netstats, operacion, argumentos):
    # Acordar formatear argumentos con coma -> Y los ints
    if operacion == "listar_operaciones":
        salida = netstat.listar_operaciones()
    elif operacion == "camino":
        origen, destino = argumentos[0].split(",")
        salida = netstat.camino(origen, destino)
    elif operacion == "mas_importantes":
        n = int(argumentos[0])
        salida = "Falta implementar mas importantes"
        #salida = netstat.mas_importantes(n)
    elif operacion == "conectados":
        pagina = argumentos[0]
        salida = "Falta implementar conectados"
        #salida = netstat.conectados(pagina)
    elif operacion == "ciclo":
        pagina, n = argumentos[0].split(",")
        salida = "Falta implementar ciclo"
        #salida = netstat.ciclo(pagina, int(n))
    elif operacion == "lectura":
        paginas = argumentos[0].split(",")
        salida = netstat.lectura(paginas)
    elif operacion == "diametro":
        salida = netstat.diametro(argumentos[0])
    elif operacion == "rango":
        salida = netstat.rango(argumentos[0])
    elif operacion == "comunidad":
        salida = netstat.comunidad(argumentos[0])
    elif operacion == "navegacion":
        salida = netstat.navegacion(argumentos[0])
    elif operacion == "clustering":
        if not argumentos:
            salida = netstat.clustering()
        else:
            salida = netstat.clustering(argumentos[0])    
    else:
        salida = "No existe operacion"
    print(salida)



if __name__ == '__main__':
    netstat = Netstats()
    archivo_tsv = sys.argv[1]  # Suponer que se pasa bien argumento
    
    # Leer tsv
    with open(archivo_tsv, "r") as archivo:
        contenido = archivo.read()
    
    # Cargar servidor -> Para cada linea de tsv, separar a lista
    for dato in contenido.split("\n"):
        linea_lista_tsv = dato.split("\t")
        netstat.agregar_post(linea_lista_tsv[0], linea_lista_tsv[1:])
    
    # Esperar comandos -> Hasta que no haya inputs entonces break
    while True:
        try:
            # Leer argumento
            linea_operacion = input()
            operacion = linea_operacion.split(" ", 1)   # Split solo el primer espacio
            # Operar argumento
            operar(netstat, operacion[0], operacion[1:])
        except EOFError:
            break

