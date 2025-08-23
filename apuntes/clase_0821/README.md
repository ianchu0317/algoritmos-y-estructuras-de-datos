# Clase Jueves 21 de Agosto

## Teoría

### Complejidad Computacional (temporal)

#### Repaso recursividad

- El caso base tiene que ser simple (mejor) ya que si es complejo entonces podria estar equivocando en algo 

- El caso general tiene que acercar al caso base

**Fibonacci**
El recursivo de fibonacci repite muchos mismos calculos, entonces _toma mucho tiempo_

#### Búsqueda de máximo en un arreglo
Dado un arreglo de n enteros, obtener maximo

#### Notación Big Oh
**¿Es mejor recursivo o iterativo? ¿Cómo medir el tiempo?**
- El mismo algoritmo puede variar dependiendo del: lenguaje, arquitectura, sistema operativo, etc
- ¿y tomar tiempo con cronometro? -> con tiempo 
- ¿Cantirdad de operaciones?

**Por que medir el tiempo?**
- Comparar algoritmos sin importar lenguaje, épocas, computadoras
- Utilizar cotas Big Oh
- El tiempo que demanda un algoritmo tiene que estar limitada por una contante

**Simplicacion**
- En numeros mayores por ejemplo `n²+n` es lo mismo que `n²`
- Operaciones es O(1)

**Complejidades tipicas**
- Constante 
- Logaritmico
- Lineal 
- Lineal-logaritmico
- Cuadrado

Se pueden plantear comparaciones de complejidades evaluando límites matemáticamente

**Por que se descartan las constantes?**

#### Ecuaciones de recurrencia


## Práctica

### Punteros y Memoria en GO

```go
var val int = 5     // declarar int
var dir *int = &val // declarar ptr de int
var val2 int = *dir // declarar int2 con valor ptr
```

**Para que utilizar punteros**
- Modificar variables locales desde funciones externas
- Trabajar con memoria dinámica

### Arreglos 
- Tamaños fijos no modificables 
```go
var arreglo [4]int = [4]int{1, 2, 3, 4} // declaracion 1
arreglo := [4]int{1, 2, 3}              // declaracion 2
len(arreglo)                            // tamaño de arreglo
```
- En funciones se pasan por valor (normal) -> se tiene que pasar por referencia
```go
func modificar(arreglo *[4]int){
    // code
}

func main(){
    // code
    var arreglo [4]int
    modificar(&arreglo)
}
```

### Slices
- Es un reference type
- Tamaño variable y dinámico
- No tiene datos por sí, sino es un slice de un array
- La capacidad máxima es el tamaño del vector
- La forma de hacer dinámico un arreglo en GO

```go
slice2 := make([]int, 0, 0)

// tamaño indefinido -> crea el array de referencia de tamaño 5
slice := []int{1, 2, 3, 4, 5}

// hacer lo de arriba es lo mismo
numeros := [5]{1, 2, 3, 4, 5}
slice := numeros[:]
```

- Puedo hacer slicing hacia la derecha para limitar
- Si hago slicing hacia la izquierda (recorto numero de inicio), entonces no se puede recuperar
- Como referencian arrays, entonces cuando paso a función entonces se modifican los valores

### Excepciónes 
- `panic("Mensaje de error")` manda error
- `recover()` recover del panic

### Valor nulo (`nil`)
El valor por defecto de valores nulos en slices y punteros es `nil`
