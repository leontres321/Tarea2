package pb

import (
	"fmt"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type DataNodeD struct {
	List_Chunk     []Chunk
	VariableEstado int // 0:Liberada 1:Buscada 2:Tomada
	MarcaTiempo    int
}

func (s *DataNodeD) EnviarD(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var resp Respuesta

	//generar conexiones a otros DataNodes y NameNode
	var connlog1 *grpc.ClientConn

	var ftps [3]FTPDistribuidoClient
	var connftp1 *grpc.ClientConn
	var connftp2 *grpc.ClientConn
	var connftp3 *grpc.ClientConn

	//NameNode Conexion
	connlog1, err := grpc.Dial("dist41:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	logs := NewLOGDistribuidoClient(connlog1)

	//DataNode 1  dist42 conexion
	connftp1, err = grpc.Dial("dist42:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[0] = NewFTPDistribuidoClient(connftp1)

	//DataNode 2  dist43 conexion
	connftp2, err = grpc.Dial("dist43:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[1] = NewFTPDistribuidoClient(connftp2)

	//DataNode 3  dist44 conexion
	connftp3, err = grpc.Dial("dist44:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ftps[2] = NewFTPDistribuidoClient(connftp3)

	if c == nil {
		fmt.Println("Error en el paquete")
		resp.Gud = false
		return &resp, nil
	}

	var p2 Propuesta

	if c.Cliente == true {
		s.List_Chunk = append(s.List_Chunk, *c)
		if c.Last == true {
			addrs, err := net.InterfaceAddrs()
			if err != nil {
				fmt.Println("Oops: " + err.Error() + "\n")
			}
			MyIp := strings.Split(addrs[1].String(), "/")[0]
			aceptada := false
			for aceptada {
				s.VariableEstado = 1
				//aqui poner funcion para ir modificando propuesta
				p2 = CrearPropuesta(int(c.Parts), c.Name)
				//---
				for i, IpAEnviar := range IPDATA {
					var AvisoEnviar Aviso
					aceptada = true

					if IpAEnviar == MyIp {
						continue
					}
					t := time.Now()
					s.MarcaTiempo = t.Hour()*3600 + t.Minute()*60 + t.Second()
					AvisoEnviar.Tiempo = int32(s.MarcaTiempo)
					r, err := ftps[i].AvisoEscrituraD(context.Background(), &AvisoEnviar)
					if err != nil {
						log.Println(err)
						os.Exit(1)
					}
					if r.Gud == false {
						aceptada = false
						break
					}
				}
			}
			//aqui propuesta aceptada
			s.VariableEstado = 2
			//zona critica
			p, err := logs.EnviarPropuestaD(context.Background(), &p2)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			//----
			s.VariableEstado = 0
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
				r, _ := ftps[nodo].EnviarD(context.Background(), &ChunkEnviar)
				log.Println(r.Gud)
			}
		}
	} else {
		// Guardarlo en disco
		NombreParte := "partes/" + c.Name + "_" + strconv.Itoa(int(c.ThisPart))
		//NombreParte := "partes/" + c.Name
		log.Println("guardando " + NombreParte + "...")
		//crear archivo
		_, err := os.Create(NombreParte)
		if err != nil {
			log.Println(err)
			os.Exit(3)
		}
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

func (s *DataNodeD) DescargarD(ctx context.Context, n *Nombre) (*Chunk, error) {
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

func (s *DataNodeD) AvisoEscrituraD(ctx context.Context, n *Aviso) (*Respuesta, error) {
	var resp Respuesta
	resp.Gud = true
	v := rand.Float64()
	if v > 0.9 {
		resp.Gud = false
	}
	//si esta tomada no responder
	for s.VariableEstado == 2 {
		//esperanding
	}
	if s.VariableEstado == 1 {
		if s.MarcaTiempo < int(n.Tiempo) {
			for s.VariableEstado == 1 {
				//esperanding
			}
		}
	}
	return &resp, nil
}

func (s *DataNodeD) mustEmbedUnimplementedFTPDistribuidoServer() {}
