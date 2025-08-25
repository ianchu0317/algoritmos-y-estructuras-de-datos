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

## Teoria

## Practica