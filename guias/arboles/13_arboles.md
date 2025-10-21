## Consigna

Implementar una función que reciba un arreglo ordenado y devuelva un arreglo o lista con los elementos en orden para ser insertados en un ABB, de tal forma que al insertarlos en dicho orden se asegure que el ABB quede balanceado. ¿Cómo cambiarías tu resolución si en vez de querer guardarlos en un ABB se fueran a insertar en un AVL?

## Mi resolucion
Basicamente para que abb sea balanceado, es devolver un preorder. Del in-order que nos da entonces parto a la mitad que va a ser raiz y luego tengo izquierda y derecha misma cantidad de elementos (puede diferir en uno pero igual queda balanceado). Insertar en la lista recursivamente repitiendo ese mismo procedimiento. 

La complejidad es O(n) porque visito todos los elementos.
```go
package balancear

func balancear[K comparable](claves []K) Lista[K] {
    listaBalanceada := CrearListaEnlazada[K]()
	balancearClaves(claves, listaBalanceada)
	return listaBalanceada
}

func balancearClaves[K comparable](claves []K, lista Lista[K]){
	if len(claves) == 0{
		return
	}
	mitad := len(claves) / 2
	// Enlistar clave
	raiz := claves[mitad]
	lista.InsertarUltimo(raiz)
	clavesMenores := claves[:mitad]
	clavesMayores := claves[mitad+1:]
	// Repetir procedimiento
	balancearClaves(clavesMenores, lista)
	balancearClaves(clavesMayores, lista)
}
// COmplejidad O(n) por teorema maestro T(n)=2T(n/2) + O(1)
```

Para AVL para que me quede igual, tendria que primero insertar raiz, luego UNO de izquierda, luego UNO de derecha, en ese orden Osea del mismo pre-order pero por nivel. Para eso se podria hacer con cola?