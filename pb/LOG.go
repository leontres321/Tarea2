package pb

import (
	"golang.org/x/net/context"
)

type NameNode struct {
	Clientes []FTPClient
	Libros   []string
}

func (s *NameNode) EnviarPropuesta(ctx context.Context, p *Propuesta) (*Propuesta, error) {
	return nil, nil
}

func (s *NameNode) SolicitarUbicacion(ctx context.Context, n *Nombre) (*Propuesta, error) {
	return nil, nil
}

func (s *NameNode) mustEmbedUnimplementedLOGServer() {}
