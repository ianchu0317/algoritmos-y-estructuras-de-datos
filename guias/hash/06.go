package hash

/*

Asumiendo que se tiene disponible una implementación completa del TDA Hash,
se desea implementar una función que decida si dos Hash dados representan o no el mismo Diccionario.
Considere para la solución que es de interés la mejor eficiencia temporal posible.
Indique, para su solución, eficiencia en tiempo y espacio.
Nota: Dos tablas de hash representan el mismo diccionario si tienen la misma cantidad de elementos;
todas las claves del primero están en el segundo; todas las del segundo, en el primero;
y los datos asociados a cada una de esas claves son iguales (se pueden comparar los valores con “==”).
*/

type Diccionario[K comparable, V comparable] struct{}

func sonIguales[K comparable, V comparable](dict1, dict2 Diccionario[K, V]) bool {
	// Chequear si tienen mismo tamaño
	if dict1.Cantidad() != dict2.Cantidad() {
		return false
	}
	// Chequear elementos de dict1 en dict2 (Clave y dato iguales)
	for iter := dict1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, dato := iter.VerActual()
		if !dict2.Pertenece(clave) || dict2.Obtener(clave) != dato {
			return false
		}
	}
	// Chequear elemetnos de dict2 in dict1
	for iter := dict2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, dato := iter.VerActual()
		if !dict1.Pertenece(clave) || dict1.Obtener(clave) != dato {
			return false
		}
	}
	return true
}
