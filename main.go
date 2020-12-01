package main

import (
	"Tarea2/client"
	"Tarea2/server"
	"fmt"
	"os"
)

func main() {
	var respuesta int8
	var tipoAlgoritmo int8
	fmt.Println("\nInicio de sistema, ingrese el número de lo que desea hacer:")
	fmt.Println("1) Cargar un libro")
	fmt.Println("2) Descargar un libro")
	fmt.Println("3) Encender almacenamiento de libros (DataNode)")
	fmt.Println("4) Almacenar LOG (NameNode)")
	fmt.Println("5) Salir")
	fmt.Println("---------------")
	fmt.Printf("Eleccion: ")

	fmt.Scanf("%d", &respuesta)

	if respuesta == 5 {
		os.Exit(0)
	}

	fmt.Println("\nTipo de algoritmo que se utilizara:")
	fmt.Println("1) Centralizado")
	fmt.Println("2) Distribuido")
	fmt.Println("---------------")
	fmt.Printf("Eleccion: ")

	fmt.Scanf("%d", &tipoAlgoritmo)

	switch respuesta {
	case 1:
		client.Run(respuesta, tipoAlgoritmo)
	case 2:
		client.Run(respuesta, tipoAlgoritmo)
	case 3:
		server.Run(respuesta, tipoAlgoritmo)
	case 4:
		server.Run(respuesta, tipoAlgoritmo)
	}
}
