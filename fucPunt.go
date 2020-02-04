package main

import (
	"fmt"
)

//puntero en funcion
func crearPuntero() *float64 {
	var myFloat = 89.6
	return &myFloat
}

//funcion con el parametro del puntero
func impriirPuntero(miBoleanoPuntero *bool) {
	fmt.Println(*miBoleanoPuntero)
	fmt.Println(miBoleanoPuntero)
}

// funcción doble usando funciones
func doble(numero *int) {
	//el valor del puntero multiplicarlo por dos
	*numero *= 2
}

func main() {
	//puntero en función
	var myFloatPuntero *float64 = crearPuntero()
	fmt.Println(*myFloatPuntero)
	//funcion con el parametro del puntero
	var miBoleano bool = true
	fmt.Println(&miBoleano)
	impriirPuntero(&miBoleano)
	//doble funcion
	cantidad := 4
	doble(&cantidad)
	//se ejecuta la funcion sin afectar la variable
	fmt.Println(cantidad)
}
