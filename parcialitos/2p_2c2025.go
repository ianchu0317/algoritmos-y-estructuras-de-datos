package parcialitos

/*
1. A lo largo de su trayectoria, la empresa ha tenido varias rotaciones de equipos. Próximamente habrá una nueva, y se busca que los
nuevos equipos incluyan personas que hayan compartido equipos antes, para facilitar la transición.
Se dispone de una base de datos que registra, para N personas, en qué M equipos ha participado cada una. La base usa un hash donde
la persona es la clave y el valor, una lista de equipos:
{
'Ana': ['Frontend-Platform', 'Growth-Squad'],
'Beto': ['Backend-Services', 'Frontend-Platform', 'Mobile-Core'],
'Carla': ['Mobile-Core'],
}
Realizar una función personasEnComun que reciba el hash, y el nombre de una persona , y devuelva una lista con todas las personas
que hayan trabajado en al menos uno de sus equipos listados. Indicar y justificar la complejidad del algoritmo implementado,
expresada con las variables N y M del problema. Por ejemplo, si se consulta por ‘Beto’, la respuesta incluye a ‘Ana’ y ‘Carla’. Si se
pregunta por ‘Beto’, la respuesta incluye a ‘Beto’.
*/
func personasEnComun(hash Diccionario[string, Lista[string]], objetivo string) Lista[string] {
	equiposObjetivo := CrearHash[string, string]()
	enComun := CrearHash[string, string]()
	listaEnComun := CrearListaEnlazada[string]()

	// Iterar cada equipo del objetivo (puede ser O(M))
	hash.Obtener(objetivo).Iterar(func(equipo string) bool {
		equiposObjetivo.Guardar(equipo, "")
		return true
	})

	// Iterar cada persona y equipo -> Pero caso O(m*n)
	hash.Iterar(func(persona string, equipos Lista[string]) bool {
		equipos.Iterar(func(equipo string) bool {
			if equiposObjetivo.Pertenece(equipo) {
				enComun.Guardar(persona, "")
				return false
			}
			return true
		})
		return true
	})

	// Evitable si se usa lista adentro (no se me ocurrio en el momento)
	enComun.Iterar(func(persona, _ string) bool {
		listaEnComun.InsertarUltimo(persona)
		return true
	})
	return listaEnComun
}

/*
2. En nuestro juego de rol táctico, Final Fantasy Algorithms, el orden de ataque se decide por el atributo iniciativa: los de mayor
iniciativa atacan primero. En esta batalla, un hechizo global está a punto de activarse, por lo que solo quedan T turnos en los que se
pueda realizar un ataque. Se desea saber, de los N personajes, cuáles lograrán atacar antes de que se active el hechizo Final Fantástico
Algorítmico. Se tiene una lista con los personajes que participarán en el combate como una estructura de formato (nombre string,
iniciativa int):
[ ('Ma-Go Lang', 95), ('Bárbara', 75), ('Cléri-Go Lang', 60), ('Arquera de bugs', 90) ]
Se pide realizar una función determinarOrdenDeAtaque que reciba la lista de combatientes, y la cantidad T turnos de turnos restantes.
La función debe devolver una lista con los nombres de los personajes que logran actuar en esa ventana de tiempo, ordenados por turno
en el que actúan. Indicar y justificar la complejidad del algoritmo implementado, expresada con las variables N y T del problema.
*/

type Personaje struct {
	nombre     string
	iniciativa int
}

// Complejidad O(n) + O(n) + O(T*log(n)) -> O(n + Tlog(n)) OJO si T <= N
func determinarOrdenDeAtaque(combatientes Lista[Personaje], T int) Lista[Personaje] {
	// Pasar los combatientes a un arreglo O(n)
	arrCombatientes := make([]Personaje, 0)
	combatientes.Iterar(func(combatiente Personaje) bool {
		arrCombatientes = append(arrCombatientes, combatiente)
		return true
	})
	// Crear Max-Heap desde arreglo O(n) (Heapify)
	heap := CrearHeapArr[Personaje](arrCombatientes, func(a, b Personaje) int {
		return a.iniciativa - b.iniciativa
	})
	// Crear lista combatientes con T-elementos del heap
	listaOrden := CrearListaEnlazada[Personaje]()
	for i := 0; i < T; i++ {
		listaOrden.InsertarUltimo(heap.Desencolar())
	}
	return listaOrden
}

/*
3. Implementar en Go una primitiva para Árbol Binario func (ab *Arbol[int]) ArbolEsPlantable() bool que determine si un
árbol es plantable, o no. Para que lo sea, todo nodo debe cumplir: el dato del nodo debe ser mayor al dato de sus hijos (si los tiene),
y además, el dato del nodo no puede superar la altura de dicho nodo. Implementar la primitiva en O(n), y justificar su complejidad.
A fines del ejercicio considerar la estructura del árbol como la definida al dorso.

*/
