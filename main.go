package main

import (
	"Tarea2/client"
	"fmt"
	"os"
)

func main() {
	var respuesta int8
	fmt.Println("\nInicio de sistema, ingrese el número de lo que desea hacer:")
	fmt.Println("1) Cargar un libro")
	fmt.Println("2) Descargar un libro")
	fmt.Println("3) Encender almacenamiento de libros (DataNode)")
	fmt.Println("4) Almacenar LOG (NameNode)")
	fmt.Println("---------------")
	fmt.Println("5) Salir")

	fmt.Scanf("%d", &respuesta)

	switch respuesta {
	case 1:
		client.Run(respuesta)
	case 2:
		client.Run(respuesta)
	case 3:

	case 4:

	case 5:
		os.Exit(0)
	}
}
