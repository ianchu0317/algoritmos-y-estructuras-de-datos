
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
  
  mitad = (ini + fin)//2
  
  izq = arr[ini:mitad+1]
  der = arr[mitad+1:fin+1]
  
  # DEBUG contadores
  # print(f"ini: {ini}, mid: {mitad}, fin: {fin}")
  
  if len(izq) > len(der):
    izq = arr[ini:mitad]
  
  # DEBUG subarray a la balanza
  # print(f"arrIzq: {izq}, arrDer: {der}")
    
  if ini == fin:
    return ini
  
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
  print(encontrar_joya(arr1))
  print(encontrar_joya(arr2))