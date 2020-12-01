package pb

import (
	"golang.org/x/net/context"
)

/*
type NameNode struct {
	Clientes []FTPClient
	Libros   []string
}
*/
func (s *NameNode) EnviarPropuestaD(ctx context.Context, p *Propuesta) (*Respuesta, error) {
	//guaradar el nombre del libro y la ubicacion de las partes
	s.Libros = append(s.Libros, p.NombreLibro)
	s.Ubicaciones = append(s.Ubicaciones, p.Lista)

	var respuesta Respuesta

	respuesta.Gud = true

	LogDistribucion(s)
	return &respuesta, nil
}

func (s *NameNode) SolicitarUbicacionD(ctx context.Context, n *Nombre) (*Propuesta, error) {
	return nil, nil
}

func (s *NameNode) mustEmbedUnimplementedLOGDistribuidoServer() {}

func (s *NameNode) PedirLibrosD(ctx context.Context, e *Empty) (*ListaLibros, error) {
	return nil, nil
}
