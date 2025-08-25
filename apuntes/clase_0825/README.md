# Clase 25 de Agosto

## Dudas

- Puede una funcion tener multiples return o mejor utilizar guard clauses ?
- Hay alguna convención en el formato de las funciones?
- La declaracion de bucle `for`
```go
// metodo 1
func sort1(vec []int) {
    var pivot int
    for i := range len(vec) {
        pivot = i
    }
}

// metodo 2 no convencional pero limpio y funciona
func sort2(vec []int) {
    for i := range len(vec) {
        pivot := i
    }
}
```

--- 

## Teoria: Abstraccion y tipos de datos abstractos

### Tipo de datos abstractos (tda)
- Tipo de dato creado para modelar algo
- Tiene funciones asociadads a este tipos de datos (primitivas, funcionalidades, métodos)
- Importa que modela (su comportamiento y qué hace) pero no cómo se implementa
    - El que usa el TDA no importa cómo se implementa sino qué
    - Lo que quiere hacer el tda que lo haga, no importa cómo y no hay que importarnos tmp (clases y objetos)
    - Se tiende a pedir valores internos de tda
- UNa persona desarrolla y otra la utiliza

### Abstraccion: que importa y que no
- cambios de tda no debe cambiar su implementacion (el valor de return y los formatos deben ser iguales)
- Está mal acceder componentes internos de un tda

### Invariantes de tda
- OBlicaciones de tda que si o si tiene que cumplir debido a la especificación 

### TDA Basicos
#### TDA Vector
- insertar al final, eliminar al final
- modificar segun posicion
- eliminar posicion
- redimension (interno)
- _listas de Python es un TDA vector_

#### TDA Pila (stack)
**Características / primitivas**
- Estructura LIFO: Last In First Out
- Apliar / desapilar
- Las primitivas son de O(1)

**Aplicaciones**
- Memoria Stack

#### TDA Cola (Queue)
Uso para organizar la llegada de datos (el orden en el que llega el orden en el que salen)

- Las primitivas son de O(1)

#### Pruebas (testing)
- Hacer pruebas para los TDA (abrir PR)
- Demostracion de que funciona (segun las especificaciones), sin meterse en el 
- Sirve para mantener. SI tengo pruebas anteriores y corrijo algo en el TDA, las pruebas anteriores me sirven para ver si se rompió algo o no
- Sirve de especificación y documentación 

**Cómo hacer testing**
- Hacer una prueba de cada tipo específico
- Modularizar código para mejor debugging -> poner los tests de funcionalidad en una función por funcionalidad. Osea no todos los tests en una sola funcion (ver tp0_tests)


---

## Practica
### Estructuras (struct)
- Guardar tipos de datos distintos en un tipo nuevo
```go
type Punto struct {
    x,y int
}

p1 := Punto{}
p2 := &Punto{}

p2.func() // no hace falta operacion especial para acceder a mismos metodos o campos
p1.func()
```

A los struct se pueden asociar métodos.

```go
// Método modifica struct
func (p Punto) Distancia() int {
    // code
}

// Método no modifica struct
func (p *Punto) Modificar() int {
    // code
}
``` 

**Anidación de struct**
Se puede acceder a campos de hijos directamente desde instancia de padre

### Manejo de errores
```go
entero, err := strconv.Atoi(str)
if err != nil {
    // ejecutar codigo de error
} else {
    // ejecutar codigo si no hay error
}
```

Declaracion de errores en una funcion
```go
func funcion(slice []int, e int) (int, error) {
    // code
    // devolver normal
    return e, nil
    // devolver con error
    return e, errors.New("Texto blabla")
}
```