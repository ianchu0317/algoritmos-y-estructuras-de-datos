package main

import (
	"errors"
	"fmt"
)

type Punto struct {
	X int
	Y int
}

// Test struct persona
type Direccion struct {
	calle  string
	numero int
}

type Persona struct {
	direccion    Direccion
	nombre       string
	padre, madre *Persona
}

func main() {
	// Testing con structs
	var p Punto = Punto{1, 2}
	fmt.Println(p)
	p.X = 3
	fmt.Println(p)

	//Testing struct anidado
	jon := Persona{nombre: "Jon Voight"}
	angie := Persona{
		Direccion{"Beverly Hills St.", 234},
		"Angelina Jolie",
		&jon,
		nil,
	}

	errors.New("JIJIJA")
	fmt.Println(angie)
}
