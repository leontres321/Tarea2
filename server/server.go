package server

import (
	"fmt"
)

func Run(int_eleccion int8) {
	var eleccion string

	if int_eleccion == 3 {
		eleccion = "DataNode"
	} else {
		eleccion = "NameNode"
	}

	fmt.Printf("Corriendo servidor como: %s", eleccion)

	if int_eleccion == 3 {
		DataNode()
	} else {
		NameNode()
	}
}

func DataNode() {
}

func NameNode() {
}
