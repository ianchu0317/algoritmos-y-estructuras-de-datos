package tda_01_test

import (
	TDA01 "guias/tdas_basicos/tda_01"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepresentacion(t *testing.T) {
	frac1 := TDA01.CrearFraccion(1, 2)
	frac2 := TDA01.CrearFraccion(13, 32)
	frac3 := TDA01.CrearFraccion(11, 11)
	require.Equal(t, "1/2", frac1.Representacion(), "Las fracciones inicializadas tienen que ser igual a la representacion")
	require.Equal(t, "13/32", frac2.Representacion(), "Las fracciones de dos digitos tienen que ser igual a la representacion")
	require.NotEqual(t, "12/12", frac3.Representacion(), "Las fracciones tienen que ser distintos")
}

func TestMultiplicar(t *testing.T) {
	frac1 := TDA01.CrearFraccion(1, 2)
	frac2 := TDA01.CrearFraccion(2, 3)

	require.Equal(t, TDA01.CrearFraccion(2, 6), frac1.Multiplicar(frac2), "Las fracciones multiplican lo mismo")
	require.Equal(t, TDA01.CrearFraccion(2, 6), frac2.Multiplicar(frac1), "Las fracciones multiplican lo mismo")
}

func TestSuma(t *testing.T) {
	frac1 := TDA01.CrearFraccion(1, 2)
	frac2 := TDA01.CrearFraccion(2, 3)

	require.Equal(t, TDA01.CrearFraccion(7, 6), frac1.Sumar(frac2), "Las fracciones suman lo mismo")
	require.Equal(t, TDA01.CrearFraccion(7, 6), frac2.Sumar(frac1), "Las fracciones suman lo mismo")
}

func TestEntero(t *testing.T) {
	frac1 := TDA01.CrearFraccion(3, 2)
	frac2 := TDA01.CrearFraccion(2, 3)
	require.Equal(t, 1, frac1.ParteEntera(), "Las fracciones tienen que devolver parte entera")
	require.Equal(t, 0, frac2.ParteEntera(), "Las fracciones tienen que devolver 0")
}
