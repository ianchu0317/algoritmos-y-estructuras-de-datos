"""
Es el año 1700, y la pirata Barba-ra Verde atacó un barco de la Royal British Shipping & Something, que transportaba una importante piedra preciosa de la corona británica. Al parecer, la escondieron en un cofre con muchas piedras preciosas falsas, en caso de un ataque. Barba-ra Verde sabe que los refuerzos británicos no tardarán en llegar, y deben uir lo más rápido posible. El problema es que no pueden llevarse el cofre completo por pesar demasiado. Necesita encontrar rápidamente la joya verdadera. La única forma de descubrir la joya verdadera es pesando. Se sabe que la joya verdadera va a pesar más que las imitaciones, y que las imitaciones pesan todas lo mismo. Cuenta con una balanza de platillos para poder pesarlas (es el 1700, no esperen una balanza digital).

a. Escribir un algoritmo de división y conquista, para determinar cuál es la verdadera joya de la corona. Suponer que hay una función balanza(grupo_de_joyas1, grupo_de_joyas2) que devuelve 0 si ambos grupos pesan lo mismo, mayor a 0 si el grupo1 pesa más que el grupo2, y menor que 0 si pasa lo contrario, y realiza esto en tiempo constante. 

b. Indicar y justificar (adecuadamente) la complejidad de la función implementada.
"""

JOYA = "a"
BASURA = "b"

def balanza(arr1, arr2):
  if len(arr1) == len(arr2):
    if JOYA in arr1:
      return 1
    elif JOYA in arr2:
      return -1
    else:
      return 0
  
  
def encontrar_joya(joyas):
  return _encontrar_joya(joyas, 0, len(joyas)-1)
  
def _encontrar_joya(arr, ini, fin):
  if ini == fin:
    return ini
  
  mitad = (ini + fin)//2
  
  izq = arr[ini:mitad+1]
  der = arr[mitad+1:fin+1]
  
  # DEBUG contadores
  # print(f"ini: {ini}, mid: {mitad}, fin: {fin}")
  
  if len(izq) > len(der):
    izq = arr[ini:mitad]
  
  # DEBUG subarray a la balanza
  # print(f"arrIzq: {izq}, arrDer: {der}")
    
  
  resultado = balanza(izq, der)
  
  if resultado == 0:
    return mitad
  elif resultado == 1:
    return _encontrar_joya(arr, ini, mitad)
  else:
    return _encontrar_joya(arr, mitad, fin)


if __name__ == '__main__':
  arr1 = [BASURA, BASURA, BASURA, JOYA, BASURA] # -> index_joya = 3
  arr2 = [JOYA, BASURA, BASURA, BASURA, BASURA] # -> index_joya = 0
  arr3 = [BASURA, BASURA, JOYA, BASURA, BASURA, BASURA, BASURA,
          BASURA, BASURA, BASURA, BASURA, BASURA, BASURA] # -> index_joya = 2
  print(encontrar_joya(arr1))
  print(encontrar_joya(arr2))
  print(encontrar_joya(arr3))