## Consigna
Implementar un algoritmo que, por división y conquista, permita obtener la parte entera de la raiz cuadrada de un numero n, en tiempo O(log(n)), por ejemplo para n=10 debe devolver 3 y para n=25 debe devolver 5

## Punto clave
- El algoritmo debe ser O(log(n)), lo que quiere decir que `C=0` `B=2`, `A=1`, haciendo que `log_2(1)= C=0` y dando como complejidad O(log(n))


## Resolucion

Con los datos de que C=0, B=2 y A=1, pude pensar que bueno todo lo que hay dentro de la funcion debe ser O(1), y debo ir dividiendo el numero por dos ya que B=2, y solo hay un llamado de recursividad A=1. Siguiendo estos datos, fui probando con cada número y siempre llego a 1 (si realizo división de numero impar, dejo siempre la parte entera por defecto).

Siguiendo esta logica y probando los numeros dividiendo recursivamente de a dos, veo que la cantidad de veces que llama la recursión es la cantidad resultado que tengo que devolver, pero tambien veo que algunos números me dan su raiz y otros no: con 25 me da 5, con 10 me da 4, con 16 me da 5, y con 36 me da 6. Aunque algunos me dieron mal, la diferencia es de uno más y uno menos (eso quiere decir que voy por buen camino creo), así que resuelvo con el codigo primero:

```go
func raizCuadrada(n int) int {
    // Caso base todos llegan acá
    if (n == 1) {
        return 1
    }
    return 1 + raizCuadrada(n/2)
}
```
Por ahí podría aplicar una condicional en el medio para ver si devolver 1 o 0 o si sumar 1 o no dependiendo si es par o impar pero esta solucion esta cerca siento.

---

## Rehacer

Si decimos que sqrt(n)=k, estamos diciendo que k²<=n< (k+1)².
Siguiendo esto entonces una forma de buscarlo es ir linealmente hasta que cumpla la condicion. 
Pero sabemos también que k debe estar entre [1,n] matemáticamente hablando, por lo que podemos verlo como "un rango" y aplicar busqueda binaria en ese rango.

Siempre va a haber una respuesta para esta solución ya que no puede no haber una raiz cuadrada para un numero porque es multiplicacion de dos. Así que no hace falta chequear fuera de rango como hariamos en una busqueda binaria tradicional

```go
func raizCuadrada(n, inicio, final int) int {
    mitad := (inicio + final) / 2
    
    // Caso k² <= n < (k+1)²
    if mitad*mitad <= n && (mitad+1)*(mitad+1) > n {
        return mitad
    } else {
        if mitad*mitad < n {
            return raizCuadrada(n, mitad+1, final)
        } else {
            return raizCuadrada(n, inicio, mitad-1)
        }
    }
}
```