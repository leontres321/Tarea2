package server

import (
	"Tarea2/pb"
	"fmt"
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

	fmt.Printf("Corriendo servidor como: %s", eleccion)

	if int_eleccion == 3 {
		DataNode(algoritmo)
	} else {
		NameNode(algoritmo)
	}
}

func DataNode(algoritmo int8) {
	//generar conexiones a otros DataNodes y NameNode
	var connlog1 *grpc.ClientConn

	var ftps [3]pb.FTPClient
	var connftp1 *grpc.ClientConn
	var connftp2 *grpc.ClientConn
	var connftp3 *grpc.ClientConn

	//NameNode Conexion
	connlog1, err := grpc.Dial("dist41:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logs := pb.NewLOGClient(connlog1)
	//DataNode 1  dist42 conexion
	connftp1, err = grpc.Dial("dist42:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[0] = pb.NewFTPClient(connftp1)
	//DataNode 2  dist43 conexion
	connftp2, err = grpc.Dial("dist43:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[1] = pb.NewFTPClient(connftp2)
	//DataNode 3  dist44 conexion
	connftp3, err = grpc.Dial("dist44:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[2] = pb.NewFTPClient(connftp3)

	//Crear server DataNode puerto 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//incializar servidor con las conexiones antes creadas
	L := pb.DataNode{Log: logs, Clientes: ftps}
	Server1 := grpc.NewServer()
	pb.RegisterFTPServer(Server1, &L)
	if err := Server1.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func NameNode(algoritmo int8) {
	//generar conexiones a otros DataNodes
	var ftps [3]pb.FTPClient
	var connftp1 *grpc.ClientConn
	var connftp2 *grpc.ClientConn
	var connftp3 *grpc.ClientConn
	//DataNode 1  dist42 conexion
	connftp1, err := grpc.Dial("dist42:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[0] = pb.NewFTPClient(connftp1)
	//DataNode 2  dist43 conexion
	connftp2, err = grpc.Dial("dist43:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[1] = pb.NewFTPClient(connftp2)
	//DataNode 3  dist44 conexion
	connftp3, err = grpc.Dial("dist44:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ftps[2] = pb.NewFTPClient(connftp3)

	//Crear server NameNode puerto 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//incializar servidor con las conexiones antes creadas
	L := pb.DataNode{Clientes: ftps}
	Server1 := grpc.NewServer()
	pb.RegisterFTPServer(Server1, &L)
	if err := Server1.Serve(lis); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
