package tdacompfunc_test

import (
	COMP "guias/tdas_basicos/04_tda_comp_func"
	"testing"

	"github.com/stretchr/testify/require"
)

func f(x float64) float64 {
	return x + 1
}

func g(x float64) float64 {
	return 2 * x
}

func TestComposicion(t *testing.T) {
	x := float64(3)
	miComposicion := COMP.CrearComposicion()
	miComposicion.AgregarFuncion(f)
	miComposicion.AgregarFuncion(g)
	require.Equal(t, float64(2*x+1), miComposicion.Aplicar(x), "Los valores no coinciden")
}
