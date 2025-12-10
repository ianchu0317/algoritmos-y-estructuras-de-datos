package parcialitos

import "math"

/*
Implementar un algoritmo que ordene un arreglo con puntos (valores (x, y)) que se encuentran dentro del círculo unitario x² + y² ≤ 1,
sabiendo que la distribución de los puntos es uniforme en dicho dominio. El criterio para ordenar es de menor a
mayor norma (distancia al origen). Tener en cuenta que los números pueden tener “infinitos” decimales.
El algoritmo debe ejecutar en tiempo lineal a la cantidad de elementos del arreglo a ordenar.
Justificar la complejidad del algoritmo propuesto.
*/

type valor struct {
	x float64
	y float64
}

func norma(v valor) float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
}

func Ordenar(valores []valor) []valor {
	// PRemisa: se sabe que x² + y² ≤ 1
	n := len(valores)
	b := n // n-baldes

	// Iterar cada valor y asignar un balde -> O(n)
	// Criterio de asignación: 0-1 * b
	baldes := make([][]valor, b)
	for i := range valores {
		_valor := valores[i]
		indice := int(norma(_valor) * float64(b)) // Valor entre [0, b]
		if indice >= b {
			indice = b - 1
		}
		baldes[indice] = append(baldes[indice], _valor)
	}

	// Ordenar cada balde con mergeSort.
	// EN PROMEDIO es 1 elemento por balde, alguno tendrá un poco más un poco menos.
	for i := range baldes {
		baldes[i] = mergeSortNorma(baldes[i])
	}

	// Iterar todos los elementos O(n)
	ordenado := make([]valor, 0, n)
	for _, balde := range baldes {
		for _, _valor := range balde {
			ordenado = append(ordenado, _valor)
		}
	}
	return ordenado
}
