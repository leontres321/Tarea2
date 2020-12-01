package server

import (
	"Tarea2/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func Run(int_eleccion int8, algoritmo int8) {
	var eleccion string

	if int_eleccion == 3 {
		eleccion = "DataNode"
	} else {
		eleccion = "NameNode"
	}

	log.Printf("Corriendo servidor como: %s", eleccion)

	if int_eleccion == 3 {
		DataNode(algoritmo)
	} else {
		NameNode(algoritmo)
	}
}

func DataNode(algoritmo int8) {
	//Crear server DataNode puerto 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//incializar servidor con las conexiones antes creadas
	L := pb.DataNode{}
	Server1 := grpc.NewServer()
	pb.RegisterFTPServer(Server1, &L)
	if err := Server1.Serve(lis); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func NameNode(algoritmo int8) {
	//Crear server NameNode puerto 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//incializar servidor con las conexiones antes creadas
	L := pb.NameNode{}
	Server1 := grpc.NewServer()
	pb.RegisterLOGServer(Server1, &L)
	if err := Server1.Serve(lis); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
