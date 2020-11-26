package pb

import (
	"golang.org/x/net/context"
)

type DataNode struct {
}

func (s *DataNode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var respuesta Respuesta
	respuesta.Id = 1
	respuesta.Gud = true
	return &respuesta, nil
}
