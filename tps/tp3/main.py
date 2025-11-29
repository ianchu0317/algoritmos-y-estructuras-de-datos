#!/usr/bin/python3

import sys
from netstats import Netstats


def operar(netstat: Netstats, comando, argumentos):
    if operacion == "listar_operaciones":
        print(netstat.listar_operaciones())

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
            operacion = input()
            print(operacion)
            # Operar argumento
            operar(netstat, operacion[0], operacion[1:])
        except EOFError:
            break

