## Consigna
Implementar una primitiva para el árbol binario (no ABB) que determine si el mismo cumple propiedad de AVL. 
La estructura del árbol es: 
```go
type arbol struct {
	izq  *arbol
	der  *arbol
	dato int
}
```

## Resolucion
```go
func (ab *arbol) cumpleAVL() bool {
    if nodo == nil {
        return true
    }
    // Calcular diferencia entre altura izquierda y derecha
    altIzq := ab.izq.Altura()   // O(n)
    altDer := ab.der.Altura()   // O(n)
    diferencia := altIzq - altDer

    if diferencia > 1 || diferencia < -1 {
        return false
    }
    // cada llamado es n/2 ya que es izquierda y derecha (suponiendo balanceado)
    return ab.izq.cumpleAVL() && ab.der.cumpleAVL()
}  


func (ab *arbol) Altura() int {
    if ab == nil {
        return 0
    }
    return 1 + max(ab.izq.Altura(), ab.der.Altura())
}
// Con esta implementacion si bien está bien y es correcta, la complejidad es O(n log n)
// La complejidad se puede calcular mediante la ecuacion de recurrencia:
// T(n) = 2T(n/2) + O(n) -> Loga(b) = 1 = C= 1 -> n^c*logb(a) = nlog(n)
```

Otra manera de calcular esto es devolviendo multiples valores al mismo tiempo.
Podemos observar las dos funciones que acabamos de implementar y fijar las similitudes.

```go
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func (ab *arbol) cumpleAVL() (bool, int) {
    // Si es hoja entonces es AVL y altura 0
    if (ab == nil) {
        return true, 0
    }
    // Obtener altura de izquierda y derecha y chequear si es AVL o no
    izqAVL, izqAlt := ab.izq.cumpleAVL()
    derAVL, derAlt := ab.der.cumpleAVL()

    // Si diferencia de altura es mayor a 1 entonces no es AVL 
    if abs(izqAlt - derAlt) > 1 {
        return false, 0
    }
    // si alguno de los subarboles no es AVL entonces salimos
    if !izqAVL || !derAVL {
        return false, 0
    } 

    return true, 1+max(izqAlt, derAlt)
}

// Con esta implementacion la complejidad en este caso acudimos a ec recurrencia
// T(n) = 2*T(n/2) + O(1) -> A=2, B=2, C=0, LogB(A)>C entonces n^logB(A)= O(n)
// Tambien podemos pensarlo que hace las dos cosas al mismo tiempo
```