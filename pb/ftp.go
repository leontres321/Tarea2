package pb

import (
	"golang.org/x/net/context"
)

type DataNode struct {
	hola int8
}

func (s *DataNode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var respuesta Respuesta
	respuesta.Id = 1
	respuesta.Gud = true
	return &respuesta, nil
}

func (s *DataNode) mustEmbedUnimplementedFTPServer() {}
