package pb

import (
	"fmt"
	"golang.org/x/net/context"
	"strconv"
	"os"
)

var IPNAME string = "10.6.40.181"

var IPDATA [3]string = [3]string{"10.6.40.182", "10.6.40.183", "10.6.40.184"}

type DataNode struct {
	List_Chunk []Chunk
	Clientes    []FTPClient
	Log        LOGClient
}

func (s *DataNode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var resp Respuesta
	resp.Gud=true

	if c == nil {
		fmt.Println("Error en el paquete")
		resp.Gud = false
		return &resp, nil
	}

	if c.Cliente == true {
		// guarda los chunks en memoria
		s.List_Chunk = append(s.List_Chunk, *c)
		if c.Last == true {			
			//crea y envia la propuesta al NameNode
			p:=CrearPropuesta(c.Parts,c.Name)
			p:=s.Log.EnviarPropuesta(p)
			for i := 0; i < len(p.Lista); i++ {
				//saca de la lista en memoria un chunk con el nombre de libro para enviar
				for i, v := range s.List_Chunk {
					if v.Name==c.Name{
						ChunkEnviar:=v
						s.List_Chunk = append(s.List_Chunk[:i],s.List_Chunk[i+1:]...)
						break
					}
				}
				//envia el chunk obtenido al nodo segun propuesta
				ChunkEnviar.Cliente=false
				nodo, _ := strconv.Atoi(p.Lista[i:i+1])
				s.Clientes[nodo].Enviar(ChunkEnviar)
			}
			
		}
	} else {
		//TODO: revisar si create funciona asi, ademas de write file porque sino
		//quedara un problema
		/*filename = s.List_Chunk[:len(s.List_Chunk)-1].Name
		_, err := os.Create("./partes/" + filename)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		ioutil.WriteFile(filename, s.List_Chunk[:len(s.List_Chunk)-1].Chunk, os.ModeAppend)
		*/
		// Guardarlo en disco
		fmt.Println("guardando " + NombreParte + "...")
		//crear archivo
		NombreParte := c.Name + "_" + strconv.Itoa(c.ThisPart)
		_, err := os.Create(NombreParte)	
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}	
		file, err := os.OpenFile(nombre, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
		_, err = file.Write(c.Chunk)		
		if err != nil {
			fmt.Println(err)
			os.Exit(8)
		}
		file.Sync()
		file.Close()

		fmt.Println(NombreParte + " Guardada")
	}

	return resp, nil
}

func (s *DataNode) Descargar(ctx context.Context, n *Nombre) (*Chunk, error) {
	//var ParteEnviar Chunk
	//mensaje.Chunk = 
	return nil, nil
}

func (s *DataNode) mustEmbedUnimplementedFTPServer() {}

func CrearPropuesta(NChunks int,Libro string) Propuesta{
	var propu Propuesta
	propu.Lista = ""
	propu.NombreLibro = Libro
	for i := 0; i < NChunks; i++ {
		n:= i%3
		if n==0
			propu.Lista += "0"
		else if n==1
			propu.Lista += "1"
		else
			propu.Lista += "2"
	}
	return propu
}
