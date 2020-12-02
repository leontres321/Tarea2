package pb

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"time"
)

var IPNAME string = "10.6.40.181"

var IPDATA [3]string = [3]string{"10.6.40.182", "10.6.40.183", "10.6.40.184"}

type DataNode struct {
	List_Chunk []Chunk
}

func (s *DataNode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var resp Respuesta
	resp.Gud = true
	//generar conexiones a otros DataNodes y NameNode
	var connlog1 *grpc.ClientConn

	var ftps [3]FTPClient
	var connftp1 *grpc.ClientConn
	var connftp2 *grpc.ClientConn
	var connftp3 *grpc.ClientConn

	//NameNode Conexion
	connlog1, err := grpc.Dial("dist41:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	logs := NewLOGClient(connlog1)

	//DataNode 1  dist42 conexion
	connftp1, err = grpc.Dial("dist42:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[0] = NewFTPClient(connftp1)

	//DataNode 2  dist43 conexion
	connftp2, err = grpc.Dial("dist43:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[1] = NewFTPClient(connftp2)
	//DataNode 3  dist44 conexion
	connftp3, err = grpc.Dial("dist44:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[2] = NewFTPClient(connftp3)

	if c == nil {
		log.Println("Error en el paquete")
		resp.Gud = false
		return &resp, nil
	}

	tiempito := time.Now()

	if c.Cliente == true {

		// guarda los chunks en memoria
		s.List_Chunk = append(s.List_Chunk, *c)
		if c.Last == true {
			//crea y envia la propuesta al NameNode
			p2 := CrearPropuesta(int(c.Parts), c.Name)
			p, err := logs.EnviarPropuesta(context.Background(), &p2)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			var ChunkEnviar Chunk
			for i := 0; i < len(p.Lista); i++ {
				//saca de la lista en memoria un chunk con el nombre de libro para enviar
				for i, v := range s.List_Chunk {
					if v.Name == c.Name {
						ChunkEnviar = v
						s.List_Chunk = append(s.List_Chunk[:i], s.List_Chunk[i+1:]...)
						break
					}
				}
				//envia el chunk obtenido al nodo segun propuesta
				ChunkEnviar.Cliente = false
				nodo, _ := strconv.Atoi(p.Lista[i : i+1])
				_, _ = ftps[nodo].Enviar(context.Background(), &ChunkEnviar)
			}
			tiempito2 := time.Since(tiempito)
			fmt.Println("Tiempo inicial: ", tiempito)
			fmt.Println("Tiempo transcurrido: ", tiempito2)

		}
	} else {
		// Guardarlo en disco
		NombreParte := "partes/" + c.Name + "_" + strconv.Itoa(int(c.ThisPart))
		//NombreParte := "partes/" + c.Name
		log.Println("guardando " + NombreParte + "...")
		//crear archivo
		f, err := os.Create(NombreParte)
		if err != nil {
			log.Println(err)
			os.Exit(3)
		}
		f.Close()

		file, err := os.OpenFile(NombreParte, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			log.Println(err)
			os.Exit(4)
		}
		_, err = file.Write(c.Chunk)
		if err != nil {
			log.Println(err)
			os.Exit(8)
		}
		file.Sync()
		file.Close()

		log.Println(NombreParte + " Guardada")
	}

	return &resp, nil
}

func (s *DataNode) Descargar(ctx context.Context, n *Nombre) (*Chunk, error) {
	newFileChunk, err := os.Open("partes/" + n.Text)
	if err != nil {
		fmt.Println("archivo"+n.Text+"no exite  err:", err)
		return nil, nil
	}
	defer newFileChunk.Close()
	chunkInfo, err := newFileChunk.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}
	var chunkSize int64 = chunkInfo.Size()
	chunkBufferBytes := make([]byte, chunkSize)
	newFileChunk.Read(chunkBufferBytes)

	return &Chunk{Chunk: chunkBufferBytes}, nil
}

func (s *DataNode) mustEmbedUnimplementedFTPServer() {}

func CrearPropuesta(NChunks int, Libro string) Propuesta {
	var propu Propuesta
	propu.Lista = ""
	propu.NombreLibro = Libro
	for i := 0; i < NChunks; i++ {
		n := i % 3
		if n == 0 {
			propu.Lista += "0"
		} else if n == 1 {
			propu.Lista += "1"
		} else {
			propu.Lista += "2"
		}
	}
	return propu
}
