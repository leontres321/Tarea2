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

	L := pb.FTPServer{}
	Server1 := grpc.NewServer()
	pb.RegisterFTPServer(Server1, &L)

	if err := Server1.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//for para agarrar todos los chunk y luego ver en que tipo de algoritmo esta

}

func NameNode(algoritmo int8) {
}
