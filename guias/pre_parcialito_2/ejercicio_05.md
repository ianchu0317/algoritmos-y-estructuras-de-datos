## Consigna
Implementar una primitiva para el árbol binario (no ABB) que determine si el mismo cumple propiedad de AVL. 
La estructura del árbol es: 
```
type arbol struct {
	izq  *arbol
	der  *arbol
	dato int
}
```

```
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
    return ab.izq.cumpleAVL() && ab.der.cumpleAVL()
}  


func (ab *arbol) Altura() int {
    if ab == nil {
        return 0
    }
    return 1 + max(ab.izq.Altura(), ab.der.Altura())
}

```

