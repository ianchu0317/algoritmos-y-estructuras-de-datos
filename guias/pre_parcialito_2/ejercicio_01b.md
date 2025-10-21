## Consigna
Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo, ordenadas tal que si insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el árbol original, pero solamente queremos las claves que estén en el sub-árbol de un elemento E.


## Resolucion
- Buscar nodo subarbol
- Agregar clave de todos los nodos del subarbol

```go

// Suponer que buscarClave devuelve un puntero al puntero que apunta al nodo padre


func (ab *abb[K, V]) ClavesSubarbol(padre K) Lista[K] {
    lista := CrearListaEnlazada[K]()
    ptPadre := ab.buscarClave(padre)   // O(log n)
    (*ptPadre).agregarClaves(lista) // O(n) -> en el peor caso padre puede ser raiz
    return lista
}

func (nodo *nodoAbb[K, V]) agregarClave(lista Lista[K]) {
    if nodo == nil {
        return 
    }
    lista.InsertarUltimo(nodo.clave)
    nodo.izq.agregarClave(lista)
    nodo.der.agregarClave(lista)
}

// Complejidad O(log n tengo que encontrar el nodo)
```