# clase lunes 1 de Septiembre

## Teoria division y conquista

### Buscar maximo en arreglo desordenado (recursividad)
- siempre pensar primero en el caso base y luego el algoritmo (en este caso si `len(arreglo)= 1`)
- 
```go
func maximo(vect []int) int {

}
```
- COn este algoritmo llegando al caso base estoy recorriendo igualmente por cada elemento
- 

### Ordenamiento
#### Merge sort
- caso base facilito como busqueda maximo. SI hay un solo elemento ya esta ordenado
- existen teoremas para calcular complejidad -> se calcula a mano primero para ver que existen metodos de calcularlos manualmente 


## Practica

### Interfaces (repaso)
- Las interfaces definen un comportamiento ( como un template, un comportamiento a seguir )
- errores
- error vs panic
    - error es mas por tecnico
    - Panic es error conceptual

### Interfaz any
- NO se puede hacer nadad dentro de una función ya que no hay metodos asociados porque no sabemos de que tipo de dato estamos operando

### Enums
- Los enums son los macros en C `#define`

### Generics
- Una misma funcion para cualquier tipo de datos 
- Por ejemplo invertir un array, el array puede ser de cualquier tipo