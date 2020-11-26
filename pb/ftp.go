package pb

import (
	"fmt"
)

type DataNode struct {
}

func (s *Datanode) Enviar(ctx context.Context, c *Chunk) (*Respuesta, error) {
	var respuesta Respuesta
	respuesta.id = 1
	respuesta.gud = true
	return &respuesta, nil
}
