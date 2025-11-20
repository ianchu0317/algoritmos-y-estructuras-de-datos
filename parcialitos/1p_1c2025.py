"""
Ejercicio Div y conquista 

Se tiene un arreglo ordenado ascendentemente el cual ha sufrido k rotaciones (el cual es desconocido), y se quiere hallar
el menor elemento del mismo. Implementar una función hallarMenor(array []int) int que lo retorne, utilizando
División y Conquista. ¿Cuál es la complejidad del algoritmo? Justificar utilizando el Teorema Maestro.
"""

def hallar_menor(arr):
  return _hallar_menor(arr, 0, len(arr)-1)

def _hallar_menor(arr, ini, fin):
  mitad = (ini + fin)//2
  print(f"indx -> ini: {ini}, mid: {mitad}, fin: {fin}")
  print(f"valor -> ini: {arr[ini]}, mid: {arr[mitad]}, fin: {arr[fin]}")
  
  if ini == fin:
    return arr[ini]
  if arr[ini] <= arr[mitad] > arr[fin]:
    print("voy derecha\n")
    return _hallar_menor(arr, mitad+1, fin)
  else:
    print("voy izquierda\n")
    return _hallar_menor(arr, ini, mitad)
    

if __name__ == '__main__':
  arr = [7, 8,0, 1, 2, 3, 4, 5, 6, 8]
  print(hallar_menor(arr))