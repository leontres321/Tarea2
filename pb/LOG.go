package pb

import (
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type NameNode struct {
	Libros      []string
	Ubicaciones []string
}

func (s *NameNode) EnviarPropuesta(ctx context.Context, p *Propuesta) (*Propuesta, error) {
	var NuevaPropuesta *Propuesta
	//aqui se pude modificar la propuesta, verificar errores, etc

	NuevaPropuesta, _ = RevisarPropuesta(p)

	//guaradar el nombre del libro y la ubicacion de las partes
	s.Libros = append(s.Libros, NuevaPropuesta.NombreLibro)
	s.Ubicaciones = append(s.Ubicaciones, NuevaPropuesta.Lista)

	LogDistribucion(s)
	return NuevaPropuesta, nil
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

func (s *NameNode) PedirLibros(ctx context.Context, e *Empty) (*ListaLibros, error) {
	var listaLibros ListaLibros
	listaLibros.NombreLibro = s.Libros
	listaLibros.CantLibros = uint64(len(s.Libros))
	return &listaLibros, nil
}

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
		aux, _ := strconv.Atoi(s.Ubicaciones[posUltimoLibro][i : i+1])
		if _, err := file.WriteString("parte_" + strconv.Itoa(i+1) + " " + IPDATA[aux] + "\n"); err != nil {
			log.Println(err)
		}
	}
}

func RevisarPropuesta(p *Propuesta) (*Propuesta, error) {
	v := rand.Float64()

	//Rechazar propuesta
	if v < 0.1 {
		largo := len(p.Lista)
		random := rand.Intn(3)

		var nuevaPropuesta Propuesta
		nuevaPropuesta.Lista = ""
		nuevaPropuesta.NombreLibro = p.NombreLibro

		for i := 0; i < largo; i++ {
			if random == 0 {
				nuevaPropuesta.Lista += "0"
			} else if random == 1 {
				nuevaPropuesta.Lista += "1"
			} else {
				nuevaPropuesta.Lista += "2"
			}
		}

		return &nuevaPropuesta, nil
	}
	//todo gud
	return p, nil
}
