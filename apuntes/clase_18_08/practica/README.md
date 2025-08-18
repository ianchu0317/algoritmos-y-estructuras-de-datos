# Introducción a GO


## Intro y primeras observaciones

- Busca reemplazar lengujes para servidores web
```go
package main // nombre del paquete
import "fmt" // 

func main() {
    fmt.Println("Hola mundo!")
}
```
- Paquete: un conjunto de archivo `.go`
- módulo (proyecto): conj de paquetes (conjuntos de archivos `.go`)
- Es un lenguaje tipado -> cada tipo de dato ocupa un espacio distinto
    - Puede haber pérdidas de datos en la conversión de datos
    - Para conversión de datos es necesario pasarlo a su 'función'

    ```go
    var a uint16 = 500
    var b uint8 = uint8(a)
    ```

## Sintaxis

### Variables

```go
// misma cosa
var a int
a = 5 
// OJO no se puede declarar dos veces la misma variable
var a int = 5
a := 5     // el lenguaje trata de adivinar el tipo de dato

// declaracion de dos variables
var a, b int = 5, 6
a, b = b, a // swap
```

```go
// declaracion de varias variables
var ( // o const
    a int = 5
    b float = 16.0
)
```
EL sintaxis de constantes es

```go
const a int = 7
```

**LOs valores iniciales de las variables sin iniciar tienen valores predeterminados según tipo**

- `0` para números

- `false` para booleanos

- `""` para strings

### Funciones

OBservación: el tipo de dato de la variable va luego. Por ejemplo en los type hints en Python

```go
// Sintaxis devolución de una sola variable
func sumar(a int, b int) int {
    return a + b
}
```

```go
// devolución de dos variables (o más)
// opcion 1
func sumarYRestar(a, b int) (int, int) {
    return a+b, a-b
}
// opcion 2
func sumarYRestar(a, b int) (suma int, resta int) {
    suma = a+b
    resta = a-b
    return
}

// opcion 3
func sumarYRestar(a, b int) (int, int) {
    var suma, resta int
    suma = a+b
    resta = a-b
    return suma, resta
}

// llamada de funcion
suma, resta := sumarYRestar(3, 4)
fmt.Println(suma) // 7
fmt.Println(resta) // -1
```

## Estructuras de control
### Booleanos
Tiene las mismas estructuras de control que en C

### If
```go
if condicion {
    // code
} else if cond2 {
    // code
} else {
    // code
}
```

### for 
```go
for init; cond; incr {
    // code
}

// ejemplo
for i := 0; i < 10; i++ {
    fmt.Printf(“%d”, i)
}

// el ejemplo se puede usar range
for i := range 10 {
    fmt.Printf("%d", i)
}
```
- NO es necesario paréntesis
- No es necesario que las condiciones e incrementos tengan relacion

**NO existe while**, es un ciclo for sin condicion de incremento e inicio
```
for condicion {

}
```

### Shadowing y scopes

```go
b := 5
for i := 1; i < 4; i++ {
    b := i
    fmt.Println(b)
}
fmt.Println(b)
```
Las llaves definen un scope nuevo, por lo que se redefine la variable. 
OJO también con la sintaxis de shortcut `:=`

### Defer
Pedazo de codigo que se ejecuta luego del bloque de codigo en un scope.

```go
package main

import "fmt"

func main() {
  defer fmt.Println("World")
  fmt.Println("Hello")
}
```

Sirve por ejemplo para no olvidarme de algo: entrada/salida de datos (archivos, network, etc)