## Consigna
Dada una lista, hacer una función que devuelva una nueva lista con todos los elementos de la lista original que no estén repetidos, y en el mismo orden relativo a la lista original. 

```
Ejemplos: 
[1, 2, 1, 3, 4, 2, 1, 5, 0, 5] → [3, 4, 0]
[1, 0, 3, 2] → [1, 0, 3, 2]
```

## Resolucion
- Iterar una vez en lista original para hallar elementos repetidos (guardar en hash O(1))
- Iterar segunda vez en lista original para hallar elementos unicos y armar nueva lista


```go
func elementosNoRepetidos(lista Lista[K]) Lista[K] {
    resultado := CrearListaEnlazada[K]()
    repetidos := CrearHash[K, bool]()

    // Iterar la lista y ver para cada elemento si esta repetido 
    lista.Iterar(func (elemento K) bool {
        esRepetido := repetidos.Pertenece(elemento)
        repetidos.Guardar(elemento, esRepetido)
        return true
    })
    // Iterar la lista segunda vez para enlistar elementos no repetidos
    lista.Iterar(func (elemento K) bool {
        if !repetidos.Obtener(elemento) {
            resultado.InsertarUltimo(elemento)
        }
        return true
    })
    return resultado
}

    // Complejidad: O(1) + O(n) + O(n) = O(n)
    // Es o(n) la primera iteracion ya que visita cada elemento de la lista original y luego para cada elemtno aplica operaciones O(1) comparando y guardando en hash. Lo mismo para segunda iteracion
```
