package pb

import (
	"golang.org/x/net/context"
	"log"
	"os"
)

type NameNode struct {
	Libros      []string
	Ubicaciones []string
}

func (s *NameNode) EnviarPropuesta(ctx context.Context, p *Propuesta) (*Propuesta, error) {
	var NuevaPropuesta Propuesta
	//aqui se pude modificar la propuesta, verificar errores, etc
	//...
	//...
	NuevaPropuesta = *p
	//guaradar el nombre del libro y la ubicacion de las partes
	s.Libros = append(s.Libros, NuevaPropuesta.NombreLibro)
	s.Ubicaciones = append(s.Ubicaciones, NuevaPropuesta.Lista)

	LogDistribucion(s)

	return &NuevaPropuesta, nil
}

func (s *NameNode) SolicitarUbicacion(ctx context.Context, n *Nombre) (*Propuesta, error) {
	var UbicacionesLibro Propuesta
	UbicacionesLibro.NombreLibro = n.Text
	for i, v := range s.Libros {
		if n.Text == v {
			UbicacionesLibro.Lista = s.Ubicaciones[i]
			return &UbicacionesLibro, nil
		}
	}
	return nil, nil
}

func (s *NameNode) mustEmbedUnimplementedLOGServer() {}

///Le agrega texto al final y si el archivo no existe lo crea
func LogDistribucion(s *NameNode) {
	file, err := os.OpenFile("Log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	if _, err := file.WriteString(s.Libros[0]); err != nil {
		log.Println(err)
	}
}
