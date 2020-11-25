package client

import (
	"fmt"
)

func Run(int_eleccion int8) {
	var eleccion string

	if int_eleccion == 1 {
		eleccion = "carga"
	} else {
		eleccion = "descarga"
	}

	fmt.Printf("Corriendo cliente como: %s", eleccion)

	if int_eleccion == 1 {
		Carga()
	} else {
		Descarga()
	}
}

func Carga() {
	var nombre string
	fmt.Printf("\nIngrese el nombre del libro para subir: ")
	fmt.Scanf("%s", &nombre)
}

func Descarga() {
}
