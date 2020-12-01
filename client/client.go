package client

import (
	"Tarea2/pb"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	//"io/ioutil"
	"math"
	"os"
	"strconv"
)

func Run(int_eleccion int8, algoritmo int8) {
	var eleccion string

	if int_eleccion == 1 {
		eleccion = "carga"
	} else {
		eleccion = "descarga"
	}

	fmt.Printf("Corriendo cliente como: %s", eleccion)

	if int_eleccion == 1 {
		Carga(algoritmo)
	} else {
		Descarga(algoritmo)
	}
}

func Carga(algoritmo int8) {
	const FILECHUNK = 250000 //1 * (1 << 18)

	var nombre string
	var dir_libros string = "./libros/"

	var conn *grpc.ClientConn

	conn, err := grpc.Dial("dist43:9000", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := pb.NewFTPClient(conn)

	fmt.Printf("\nIngrese el nombre del libro para subir: ")
	fmt.Scanf("%s", &nombre)

	file, err := os.Open(dir_libros + nombre)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	totalParts := uint64(math.Ceil(float64(fileSize) / float64(FILECHUNK)))

	//Para seguir el tutorial al pie de la letra mi nino
	fmt.Printf("\nDividiendo el archivo en %d partes de tamano %dkB\n", totalParts, FILECHUNK/1000)

	for i := uint64(0); i < totalParts; i++ {
		partSize := int(math.Min(FILECHUNK, float64(fileSize-int64(i*FILECHUNK))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		fileName := nombre

		/*_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		*/

		var mensaje pb.Chunk

		mensaje.Name = fileName
		mensaje.Parts = totalParts
		mensaje.ThisPart = i + 1

		mensaje.Chunk = partBuffer
		mensaje.Last = false
		if i == totalParts-1 {
			mensaje.Last = true
		}
		mensaje.Cliente = true

		//Aca se puede enviar el chunk
		resp, err := c.Enviar(context.Background(), &mensaje)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(resp.Gud)
	}

}

func Descarga(algoritmo int8) {
	var nombre string

	//Preguntar a NameNode si existe y preguntar cantidad de partes
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dist41:9000", grpc.WithInsecure())
	c := pb.NewLOGClient(conn)

	var libros *pb.ListaLibros
	var meh pb.Empty
	libros, err = c.PedirLibros(context.Background(), &meh)

	fmt.Println("Lista de libros en el sistema:")

	if libros.CantLibros == 0 {
		fmt.Println("No hay libros actualmente, no se puede descargar")
		os.Exit(0)
	}

	for i := uint64(0); i < libros.CantLibros; i++ {
		fmt.Println(libros.NombreLibro[i])
	}

	fmt.Println("--------------")
	fmt.Printf("\nIngrese el nombre del libro a bajar: ")
	fmt.Scanf("%s", &nombre)

	var n pb.Nombre
	n.Text = strings.TrimSpace(nombre)
	p, err := c.SolicitarUbicacion(context.Background(), &n)

	if p == nil {
		fmt.Println("No exite el archivo")
		os.Exit(0)
	}

	//Asumiendo que existe
	_, err = os.Create(nombre)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	file, err := os.OpenFile(nombre, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	var ftps [3]pb.FTPClient
	var connftp1 *grpc.ClientConn
	var connftp2 *grpc.ClientConn
	var connftp3 *grpc.ClientConn

	//DataNode 1  dist42 conexion
	connftp1, err = grpc.Dial("dist42:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[0] = pb.NewFTPClient(connftp1)

	//DataNode 2  dist43 conexion
	connftp2, err = grpc.Dial("dist43:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[1] = pb.NewFTPClient(connftp2)
	//DataNode 3  dist44 conexion
	connftp3, err = grpc.Dial("dist44:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[2] = pb.NewFTPClient(connftp3)

	for i := 0; i < len(p.Lista); i++ {
		currentChunkFileName := nombre + "_" + strconv.FormatUint(uint64(i+1), 10)
		nodo, _ := strconv.Atoi(p.Lista[i : i+1])
		chunkRecibido, err := ftps[nodo].Descargar(context.Background(), &pb.Nombre{Text: currentChunkFileName})

		/*newFileChunk, err := os.Open(currentChunkFileName)

				if err != nil {
					fmt.Println(err)
					os.Exit(5)
				}

				defer newFileChunk.Close()

				chunkInfo, err := newFileChunk.Stat()

				if err != nil {
					fmt.Println(err)
					os.Exit(6)
				}
		        // bytes cantidad
				var chunkSize int64 = chunkInfo.Size()
				chunkBufferBytes := make([]byte, chunkSize)


				writePosition = writePosition + chunkSize

				/*reader := bufio.NewReader(newFileChunk)
				_, err = reader.Read(chunkBufferBytes)

				if err != nil {
					fmt.Println(err)
					os.Exit(7)
				}*/

		_, err = file.Write(chunkRecibido.Chunk)

		if err != nil {
			fmt.Println(err)
			os.Exit(8)
		}

		file.Sync()

		//chunkBufferBytes = nil
	}
	fmt.Println("Libro descargado")

	file.Close()
}
