## Consigna

Carlos es nuevo en la empresa en la que trabajan Alan y Bárbara. Alan va a ser el mentor de Carlos, quien debe implementar un nuevo TDA Gatito. Alan, revisando el trabajo que hizo Carlos, nota que este agregó una primitiva Redimensionar, pública en la interfaz Gatito, para que la use Bárbara. Alan lo increpa a Carlos, preguntando para qué es dicha primitiva, y este le contesta “Tal como dice la documentación, es para que Bárbara me diga cómo redimensionar el arreglo de pelos que tiene el gatito”. Alan, que conoce bien el temperamento de Bárbara, decide evitar que echen a Carlos en su segunda semana de trabajo. En este ejercicio, te toca hacer de Alan.

Escribir una explicación de por qué esto que está haciendo Carlos está mal. Considerá que Carlos es muy testarudo (incluso, a pesar de su propio bien), así que tu argumentación deberá ser muy clara y contundente.

## Resolucion
El error acá está en exponer una función auxiliar que sería redimensión del arreglo el cual es parte de la estructura interna de TDA. Es una función que debería utilizarlo al implementar las primitivas no cuando se utilicen las primitivas.

## Corrección - Argumentos más contundentes para Carlos:

### 1. **Violación del principio de encapsulamiento**
- El TDA debe **ocultar** su implementación interna
- Bárbara NO debe saber que existe un "arreglo de pelos" 
- Si cambias la implementación (lista enlazada, hash table, etc.), ¿Bárbara tendría que cambiar su código?

### 2. **Responsabilidad incorrecta**
- **Responsabilidad del TDA**: Manejar automáticamente el crecimiento del arreglo
- **Responsabilidad de Bárbara**: Solo usar las operaciones conceptuales del gatito
- ¿Por qué Bárbara debe decidir el tamaño técnico del arreglo?

### 3. **Problemas prácticos graves**
- **Bugs**: ¿Qué pasa si Bárbara redimensiona más pequeño y pierdes datos?
- **Performance**: Bárbara no sabe cuándo es eficiente redimensionar
- **Complejidad**: Bárbara debe entender detalles técnicos innecesarios

### 4. **Mal diseño de interfaz**
```go
// MAL - Expone implementación
type Gatito interface {
    AgregarPelo(color string)
    Redimensionar(nuevoTamaño int) // ← ¡ERROR!
}

// BIEN - Solo operaciones conceptuales  
type Gatito interface {
    AgregarPelo(color string)
    ContarPelos() int
    TieneColor(color string) bool
}
```

**Resultado**: Bárbara puede usar el TDA sin saber nada de arreglos internos.