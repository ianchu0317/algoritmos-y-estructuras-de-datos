## Consigna

Se quiere implementar un TDA Diccionario con las siguientes primitivas: Obtener(x) devuelve el valor de x en el diccionario; Insertar(x, y) inserta en el diccionario la clave x con el valor y (entero); Borrar(x) borra la entrada de x; Add(x, n) le suma n al contenido de x; AddAll(m) le suma m a todos los valores.

Proponer una implementación donde todas las operaciones sean O(1). Justificar el orden de las operaciones.

## Resolucion
Como es implementacion no uso de funciones, entonces podemos hacer lo que queramos. Basicamente puedo crear un nuevo struct que utilice un diccionario normal para guardar los datos, luego un campo donde guardo el int addAll que inicialmente va a ser 0. Pero para que se actualicen por ejemplo si guardo(x, 5) y luego hago addAll(5) y luego guardo(y, 5) entonces quiero que x=10, y=5. Para eso entonces necesito otro campo mas donde va a ser un diccionario donde guardo los "Insert". Entonces cada vez que creo uno nuevo quiero guardar primero en diccionario de datos su valor real, luego en el diccionario de insert -addAll . Cuando haga obtener a mi diccionario actual quiero devolver: su dato actual + dato insertado (que si no se modifico es -addAll) y addAll. Entonces si no se hizo addAll entonces se anulan y es 0. Luego si addAll cambia de valor, todavia sigo teniendo el anterior.