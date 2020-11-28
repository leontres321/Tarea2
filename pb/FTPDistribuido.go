package pb

import (
	"fmt"
	"golang.org/x/net/context"
)

type DataNodeD struct {
	List_Chunk []Chunk
	Cliente    []FTPClient
	Log        LOGClient
}

func (s *DataNodeD) EnviarD(ctx context.Context, c *Chunk) (*Respuesta, error) {
	if c == nil {
		fmt.Println("Error en el paquete")
		var resp Respuesta
		resp.Gud = false
		return &resp, nil
	}

	if c.Cliente == true {
		s.List_Chunk = append(s.List_Chunk, *c)
		if c.Last == true {
			fmt.Println("hola")
			//EnviarPropuesta(s)
			//enviar a otras maquinas y escribir las partes que se queda esta
		}
	} else {
		// Guardarlo en disco
	}

	return nil, nil
}

func (s *DataNodeD) DescargarD(ctx context.Context, n *Nombre) (*Chunk, error) {
	return nil, nil
}

func (s *DataNodeD) AvisoEscrituraD(ctx context.Context, n *Nombre) (*Respuesta, error) {
	return nil, nil
}

func (s *DataNodeD) mustEmbedUnimplementedFTPDistribuidoServer() {}
