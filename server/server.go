package server

import (
	"Tarea2/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

func Run(int_eleccion int8, algoritmo int8) {
	var eleccion string

	if int_eleccion == 3 {
		eleccion = "DataNode"
	} else {
		eleccion = "NameNode"
	}

	fmt.Printf("Corriendo servidor como: %s", eleccion)

	if int_eleccion == 3 {
		DataNode(algoritmo)
	} else {
		NameNode(algoritmo)
	}
}

func DataNode(algoritmo int8) {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	L := pb.DataNode{}
	Server1 := grpc.NewServer()
	pb.RegisterFTPServer(Server1, &L)

	if err := Server1.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//for para agarrar todos los chunk

	//Luego de tener todos los chunk debe enviarle una propuesta a alguien
	//Si algoritmo == 1 entonces debe enviarle la propuesta al NameNode
	//Si algoritmo == 2 entonces debe enviarle la propuesta al resto de DataNodes (1)

	//luego de tener una propuesta aceptada debe enviar los chunks a sus respectivos dueños
}

func NameNode(algoritmo int8) {
	//Si algoritmo != 1 entonces no es necesario que reciba propuestas, solo escrituras al log
}
