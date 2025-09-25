package tdacompfunc

// Interfaz de TDA
type ComposicionFunciones interface {
	AgregarFuncion(func(float64) float64)
	Aplicar(float64) float64
}

// DECLARACION E IMPLEMENTACION TDA COMPOSICION
// Tipo composicion
type Composicion struct {
	funciones [](func(float64) float64)
}

// AgregarFuncion agrega una función func(float64) float64 a la composicion
func (c *Composicion) AgregarFuncion(f func(float64) float64) {
	c.funciones = append(c.funciones, f)
}

// Aplicar devuelve la aplicacion de las funciones dadas
func (c Composicion) Aplicar(x float64) float64 {
	for i := len(c.funciones) - 1; i >= 0; i-- {
		funcionActual := c.funciones[i]
		x = funcionActual(x)
	}
	return x
}

func CrearComposicion() ComposicionFunciones {
	nuevaComposicion := Composicion{funciones: make([]func(float64) float64, 0)}
	return &nuevaComposicion
}
