package arboles

/*
Implementar una primitiva para el AB que reciba dos arreglos (o listas) de cadenas.
El primer arreglo corresponde al preorder de un árbol binario.
El segundo al inorder del mismo árbol (ambos arreglos tienen los mismos elementos, sin repetidos).
La función debe devolver un árbol binario que tenga dicho preorder e inorder.
Indicar y justificar el orden de la primitiva (tener cuidado con este punto). */

type Arbol struct {
	valor int
	izq   *Arbol
	der   *Arbol
}

func (ab *Arbol) ConstruirArbol(preorder, inorder []int) {
	if len(preorder) == 0 {
		return
	}

	// Hallar elemento de medio
	valorMedio := preorder[0]
	// Hallar indx medio en inorder
	var indxMedio int
	for i := range inorder {
		if inorder[i] == valorMedio {
			indxMedio = i
			break
		}
	}
	// Separar inorder y preorder de izq y der
	inorderIzq := inorder[:indxMedio]
	inorderDer := inorder[indxMedio+1:]
	preorderIzq := preorder[1 : 1+len(inorderIzq)]
	preorderDer := preorder[1+len(inorderIzq):]

	// Construir arbol
	ab.valor = valorMedio
	ab.izq.ConstruirArbol(preorderIzq, inorderIzq)
	ab.der.ConstruirArbol(preorderDer, inorderDer)
}
