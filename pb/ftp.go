package pb

import (
	"fmt"
	"golang.org/x/net/context"
)

var IPNAME string = "10.6.40.181"

var IPDATA [3]string = [3]string{"10.6.40.182", "10.6.40.183", "10.6.40.184"}

type DataNode struct {
	List_Chunk []Chunk
	Cliente    FTPClient
	Log        LOGClient
}

func (s *DataNode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	if c == nil {
		fmt.Println("Error en el paquete")
		var resp Respuesta
		resp.Gud = false
		return &resp, nil
	}

	if c.Cliente == true {
		s.List_Chunk = append(s.List_Chunk, *c)
		if c.Last == true {
			EnviarPropuesta(s)
			//enviar a otras maquinas y escribir las partes que se queda esta
		}
	} else {
		// Guardarlo en disco
	}

	return nil, nil
}

func (s *DataNode) Descargar(ctx context.Context, n *Nombre) (*Chunk, error) {
	return nil, nil
}

func (s *DataNode) mustEmbedUnimplementedFTPServer() {}

func EnviarPropuesta(s *DataNode) bool {
	return false
}
