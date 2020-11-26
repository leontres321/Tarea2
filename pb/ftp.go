package pb

import (
	"golang.org/x/net/context"
)

type DataNode struct {
}

func (s *Datanode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var respuesta Respuesta
	respuesta.id = 1
	respuesta.gud = true
	return &respuesta, nil
}
