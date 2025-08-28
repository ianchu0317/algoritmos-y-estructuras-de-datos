# Clase 28 de Agosto

## Teoría

### Repaso TDA
- (TDA) -> Estructuras de datos
- Tipo de dato que creo para modelar algo


### Implementación de estructuras
- Hay que ver distintos temas: 
    - algunas son mas faciles
    - Algunas son complejas (temporal, espacial)
    - Cuan fácil es entender la solución, 
    - Tiempo de implementación

### Estructuras sobre arreglos

**TDA Pila**

TDA Pila sobre arreglo, lo que hay que hacer no necesario reemplazar el ultimo elemento cada vez que elimino. Sino que cambio lógicamente osea la variable que cuenta la lógica de largo del arreglo

Cada vez que apilo y des-apilo es una operación O(1)

No utilizar append porque en GO no existe achicar memoria. Además al copiar todos los elementos es un O(n). Osea si aumento el tamaño de a 3 (capacidad), entonces cada vez que se llena tengo que agrandar de nuevo y correr una complejidad de O(n)

-> Analisis de complejidad amortizado


**Como achicar?**
- Podria pensar que achico luego de que este 1/4 ocupado nada mas, asi todavia tengo la mitad libre. -> el problema con esto era que hacia expansion y achicacion 


### ERstructuras enlazadas
- Facil agrandar y achicar
- Dificil indexar
- No hay redimensiṕon
- Hay que mantener correctamente las referencias


- Pila.cantidad está mal por ejemplo para el usuario porque es algo interno de la estructura de datos

## Práctica

### Repaso TDA
- Nos importa que hace pero no cómo se hace. -> de ahi viene la palabra abstracción 
- No manipular cosas internas de TDA
- Las cosas internas de TDA no se exportan
- solo nos importan lo que hace el tda


### Invariantes de primitivas
- Sirven para especificar tda
- Compuestas por precondiciones y postcondiciones de primitivias (contratos) (documentación)

### Interfaz vs implementación
- Interfaz: que dice que hace -> firmas y documentacion
- Implementación: 


#### Interfaces
- Interfaz es un conjunto de firmas de primitvas
- Que el que implementa y la persona que usa sepan que hacer y que no hacer
- Evitar exponer cómo está implementado

Pueden haber las funciones con mismo nombre ya que pueden ser distintos métodos de otras estructuras

```go
type figura interface {
    area() float64
    perimetro() float64
}

type rect struct {
    ancho, alto float64
}

func (r rect) area() float64 {
    return r.ancho * r.alto
}

func (r rect) perimetro() float64 {
    return 2 * r.ancho + 2 * r.alto
}
// como rect cumple area y perimetro entonces es una figura
// Tener interfaces permite crear funciones generales para una interfaz
// Entender la interfaz como un clase padre y la herencia
func medidas(f figura) {
    fmt.Pritln(f)
    fmt.Println(f.area())
    fmt.Println(f.perimetro())
}
```

- Las interfaces contiene que hacen las primitivas

```go
type Puerta struct {
    estado int;  // si está capitalizado es un atributo público, si es minuscula es privado
}

```