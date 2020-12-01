package pb

import (
	"golang.org/x/net/context"
)

type NameNode struct {
	Clientes    []FTPClient
	Libros      []string
	Ubicaciones []string
}

func (s *NameNode) EnviarPropuesta(ctx context.Context, p *Propuesta) (*Propuesta, error) {
	var NuevaPropuesta Propuesta
	//aqui se pude modificar la propuesta, verificar errores, etc
	//...
	//...
	NuevaPropuesta = p
	//guaradar el nombre del libro y la ubicacion de las partes
	Libros = append(Libros, Propuesta.NombreLibro)
	Ubicaciones = append(Ubicaciones, Propuesta.lista)
	return NuevaPropuesta, nil
}

func (s *NameNode) SolicitarUbicacion(ctx context.Context, n *Nombre) (*Propuesta, error) {
	var UbicacionesLibro Propuesta
	UbicacionesLibro.NombreLibro = n.Text
	for i, v := range s.Libros {
		if n.Text == v {
			UbicacionesLibro.Lista = s.Ubicaciones[i]
			break
		}
	}
	return UbicacionesLibro, nil
}

func (s *NameNode) mustEmbedUnimplementedLOGServer() {}
