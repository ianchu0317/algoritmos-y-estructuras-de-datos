#!/usr/bin/python3

import sys
from netstats import Netstats



if __name__ == '__main__':
    netstat = Netstats()
    archivo_tsv = sys.argv[1]  # Suponer que se pasa bien argumento
    
    # Leer tsv
    with open(archivo_tsv, "r") as archivo:
        contenido = archivo.read()
        print("- linea es: ", contenido)
    # Cargar servidor


