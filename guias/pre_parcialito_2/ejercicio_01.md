## Consigna
Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo, ordenadas tal que si insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el arbol.

## Resolucion
```go
type nodoAbb[K, V any] struct {
    izq *nodoABB[K, V]
    der *nodoABB[K, V]
    clave K
    dato V
}

type abb struct {
    raiz *nodoAbb[K, V]
}

func (ab *abb[K, V]) ClavesAbb() Lista[K] {
    lista := CrearListaEnlazada[K]()
    abb.raiz.agregarClave(lista)
    return lista
}

// nodo -> izq -> derecha

func (nodo *nodoAbb[K, V]) agregarClave(lista Lista[K]) {
    if nodo == nil {
        return 
    }
    lista.InsertarUltimo(nodo.clave)
    nodo.izq.agregarClave(lista)
    nodo.der.agregarClave(lista)
}
```

