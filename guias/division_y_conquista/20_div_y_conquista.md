## Consigna
Es el año 1700, y la pirata Barba-ra Verde atacó un barco de la Royal British Shipping & Something, que transportaba una importante piedra preciosa de la corona británica. Al parecer, la escondieron en un cofre con muchas piedras preciosas falsas, en caso de un ataque. Barba-ra Verde sabe que los refuerzos británicos no tardarán en llegar, y deben uir lo más rápido posible. El problema es que no pueden llevarse el cofre completo por pesar demasiado. Necesita encontrar rápidamente la joya verdadera. La única forma de descubrir la joya verdadera es pesando. Se sabe que la joya verdadera va a pesar más que las imitaciones, y que las imitaciones pesan todas lo mismo. Cuenta con una balanza de platillos para poder pesarlas (es el 1700, no esperen una balanza digital).

a. Escribir un algoritmo de división y conquista, para determinar cuál es la verdadera joya de la corona. Suponer que hay una función balanza(grupo_de_joyas1, grupo_de_joyas2) que devuelve 0 si ambos grupos pesan lo mismo, mayor a 0 si el grupo1 pesa más que el grupo2, y menor que 0 si pasa lo contrario, y realiza esto en tiempo constante. 

b. Indicar y justificar (adecuadamente) la complejidad de la función implementada.

## Resolucion
Lo mismo que una busqueda binaria: parto por la mitad y me quedo, y comparo el peso de las dos mitades, busco en la mitad más pesada. Si las mitades pesan lo mismo quiere decir que tengo actualmente la original

```go
func buscarJoya(grupo_de_joyas []int, inicio, final int) joya {
    // Caso base si se encontro
    if inicio == final {
        return inicio
    }

    medio := (inicio + final) / 2
    grupo1 := grupo_de_joyas[inicio:medio]
    grupo2 := grupo_de_joyas[medio:final]
 
    if balanza(grupo1, grupo2) > 0 {
        return buscarJoya(grupo_de_joyas, inicio, medio - 1)
    } else {
        return buscarJoya(grupo_de_joyas, medio + 1, final)
    }
    
}
```

El algooritmo tiene complejidad O(log(n)) igual que la busqueda binaria ya que si aplicamos teorema maestro a la ecuación de recurrencia: T(n)=T(n/2)+O(1) A=1, B=2, C=0 y log_2(1) = C = 0