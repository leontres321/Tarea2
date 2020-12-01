package pb

import (
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"strings"
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

	posUltimoLibro := len(s.Libros) - 1

	//ERROR: s.ubicaciones bla bla no entrega bien el largo de ubicaciones
	if _, err := file.WriteString(strings.TrimSpace(s.Libros[posUltimoLibro]+" "+strconv.Itoa(len(s.Ubicaciones[posUltimoLibro]))) + "\n"); err != nil {
		log.Println(err)
	}

	for i := 0; i < len(s.Ubicaciones[posUltimoLibro]); i++ {
		//log.Printf("parte_" + strconv.Itoa(i+1) + " " + string(s.Ubicaciones[posUltimoLibro][i]) + "\n")
		if _, err := file.WriteString("parte_" + strconv.Itoa(i+1) + " " + string(s.Ubicaciones[posUltimoLibro][i]) + "\n"); err != nil {
			log.Println(err)
		}
	}
}

func (s *NameNode) PedirLibros(ctx context.Context, e *Empty) (*ListaLibros, error) {
	var listaLibros ListaLibros
	listaLibros.NombreLibro = s.Libros
	listaLibros.CantLibros = uint64(len(s.Libros))
	return &listaLibros, nil
}
